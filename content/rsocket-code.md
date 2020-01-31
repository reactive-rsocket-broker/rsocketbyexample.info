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

# References

* RSocket Java SDK: https://github.com/rsocket/rsocket-java
* Spring RSocket: https://docs.spring.io/spring/docs/5.2.3.RELEASE/spring-framework-reference/web-reactive.html#rsocket
