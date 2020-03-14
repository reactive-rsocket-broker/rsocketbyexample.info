+++
title = "30 seconds of RSocket code"
+++


借鉴 30-seconds-of-code 系列，将RSocket使用到的代码场景整理一下，方便大家参考。RSocket的API相对来说比较简单，而且Spring Framework和Spring Boot都有对其封装，所以API已经非常精简。


# Composite Metadata

### 构建Composite Metadata
Composite Metadata是RSocket Payload的metadata数据格式，方便我们为Payload提供各种元数据信息。构建Composite Data非常简单，代码如下：

```java
public static CompositeByteBuf buildConnectionSetupMetadata(final String clientId) {
        CompositeByteBuf metadataByteBuf = ByteBufAllocator.DEFAULT.compositeBuffer();

        // Adding the clientId to the composite metadata
        CompositeMetadataFlyweight.encodeAndAddMetadata(
                metadataByteBuf,
                ByteBufAllocator.DEFAULT,
                "messaging/x.clientId",
                ByteBufAllocator.DEFAULT.buffer().writeBytes(clientId.getBytes()));

        return metadataByteBuf;
    }

...
CompositeByteBuf metadataByteBuf = buildConnectionSetupMetadata(clientId);
RSocket rSocket = RSocketFactory.connect()
                .setupPayload(DefaultPayload.create(Unpooled.EMPTY_BUFFER, metadataByteBuf))
                .transport(TcpClientTransport.create(7000))
                .start()
                .block();
```

当然解析也非常简单，如下：

```java
private static Map<String, Object> parseMetadata(Payload payload) {
        Map<String, Object> metadataMap = new HashMap<>();

        CompositeMetadata compositeMetadata = new CompositeMetadata(payload.metadata(), true);
        compositeMetadata.forEach(entry -> {
            byte[] bytes = new byte[entry.getContent().readableBytes()];
            entry.getContent().readBytes(bytes);
            metadataMap.put(entry.getMimeType(), new String(bytes, StandardCharsets.UTF_8));
        });
        return metadataMap;
    }
```

# Request/Response

Request/Response是典型的RPC调用，发送请求然后等待返回。在RSocket中，这个等待是异步的，而不是同步的，不会浪费你宝贵的线程资源。

### Request请求

```
   // Sending the request
   rSocket.requestResponse(DefaultPayload.create(name))
           .map(Payload::getDataUtf8)
           .subscribe(msg -> {
               // Handling the response
               LOG.info("Response: {}", msg);
           });
```

### Response响应

```
public Mono<Payload> requestResponse(Payload payload) {
        String name = payload.getDataUtf8();
        if (name == null || name.isEmpty()) {
            name = "You";
        }
        return Mono.just(DefaultPayload.create(String.format("Hello, %s!", name)));
}
```

# Fire-and-Forget

Fire-and-Forget的使用场景也非常多，如会员注册过程中，发送短信验证码；注册后发送一封欢迎邮件等等，如果你使用堵塞的方式，用户要等待非常时间，体验非常不好，现在只需要Fire-and-Forget后，马上就可以返回。

```
   // Sending the request
   rSocket.fireAndForget(DefaultPayload.create(name))
           .map(Payload::getDataUtf8)
           .subscribe(msg -> {
               // Handling the response
               LOG.info("Response: {}", msg);
           });
```

# Request/Stream

### Subscribe a stream
Request/Stream可以像普通订阅一样

```java
 rSocket.requestStream(DefaultPayload.create(Unpooled.EMPTY_BUFFER))
                .doOnComplete(() -> {
                    LOG.info("Done");
                })
                .subscribe(payload -> {
                    byte[] bytes = new byte[payload.data().readableBytes()];
                    payload.data().readBytes(bytes);
                    LOG.info("Received: {}", new BigInteger(bytes).intValue());
                });
```

# RSocket中的设计模式

### RSocket Responder创建模式

RSocket Responder是指接收并处理通讯的对方发过来的请求，这里我们会使用一个工厂模式负责创建对应Responder Handler，结构图如下：

![RSocket Responder](/images/misc/rsocket_responder.png)

SimpleResponderFactory会创建一个"SocketAcceptor responder()"方法，主要是为RSocket便捷操作，同时包含一些验证操作，而核心createResponder()负责创建具体的Responder Handler，如SimpleResponderImpl类。
SimpleResponderImpl类不像Servlet那样，是Singleton的。 RSocket中Responder Handler通常是每一个连接对应一个Responder对象，类似于Session的机制，所以Responder Handler中的实力变量是针对连接的，
而其中的requester就是通讯的对方，这样对等通讯完全没有问题。

详细的代码可以参考： https://github.com/linux-china/rsocket-simple-demo

### RSocket Session设计
Session是指调用方的一个长连接生命存续期间的数据，如HTTP的Session是基于Cookie的机制实现的，所以我们可以非常简单地获取HttpSession对象，在多个请求之间共享一些数据。 那么在RSocket中如何设计Session机制呢？
Reactor框架中有一个Reactor Context的概念，你可以在 https://projectreactor.io/docs/core/release/reference/#context 这里找到。

* 由于Reactor Context默认是不可变的，当你对context进行put操作会生成新的Context对象，所以我们要进行一些调整，创建一个默认可变的Context，如下代码：

```java
public class MutableContext implements Context {

	HashMap<Object, Object> holder = new HashMap<>();

  ...
}
```

* 接下来我们一个RSocketInterceptor，对RSocket进行拦截处理。 在拦截代码中，我们会调用subscriberContext() 添加context支持，代码如下：

```
public class RSocketSessionInterceptor implements RSocketInterceptor {
    @Override
    public RSocket apply(RSocket source) {
        return new AbstractRSocket() {
            private MutableContext mutableContext = new MutableContext();

            @Override
            public Mono<Payload> requestResponse(Payload payload) {
                return source.requestResponse(payload).subscriberContext(mutableContext::putAll);
            }
        };
```

* 添加responder plugin，调用RSocketSessionInterceptor的context处理逻辑，代码如下：

```
RSocketFactory.receive()
                .addResponderPlugin(new RSocketSessionInterceptor())
                .acceptor(responderFactory.responder())

```

* 最后调用deferWithContext获取context，并访问session数据，代码如下：

```java
 public Mono<Payload> requestResponse(Payload payload) {
    return Mono.deferWithContext((context -> {
            //use context to get and set session data
    });
```

详细的代码在 https://github.com/linux-china/rsocket-simple-demo/blob/master/src/main/java/org/mvnsearch/rsocket/demo/RSocketSessionInterceptor.java


# 其他

### 异常日志处理
RSocket Java SDK中默认的异常处理是调用异常的printStackTrace()方法，如果你要调整异常的记录方式，可以调用errorConsumer进行调整，代码如如下：

```java
RSocketFactory.receive()
                .errorConsumer(error -> {
                    // logging
                })
```

### RSocket连接层拦截
如果想做一些连接层的拦截，也就是字节流发送到网络之前，你可以使用DuplexConnectionInterceptor。在这一层你可以进行一些网络扩展，如流控等，代码如下：

```java
public class TokenBucketInterceptor implements DuplexConnectionInterceptor {
    @Override
    public DuplexConnection apply(Type type, DuplexConnection source) {
        if (type.equals(Type.CLIENT)) {
            return new DuplexConnectionProxy(source) {
                @Override
                public Mono<Void> send(Publisher<ByteBuf> frames) {
                    //token bucket control
                    return super.send(frames);
                }
            };
        }
        return source;
    }
}
```

# References

* RSocket Java SDK: https://github.com/rsocket/rsocket-java
* Spring RSocket: https://docs.spring.io/spring/docs/5.2.3.RELEASE/spring-framework-reference/web-reactive.html#rsocket
