+++
title = "30 seconds of RSocket code"
+++


借鉴 30-seconds-of-code 系列，将RSocket使用到的代码场景整理一下，方便大家参考。


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

