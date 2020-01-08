+++
title = "Peer-to-Peer - 角色对等"
+++

在网络通讯中，大家对Client/Server模式一定不陌生，也就是Client发起请求，Server做出响应，如我们了解的HTTP协议、自定义RPC等都属于这个范畴。Client/Server模型是对TCP通讯模型的裁剪，让其符合我们的理解，模型也更简单些。
而在TCP模型中，通讯的双方都是可以相互发送请求，都可以向Socket写入数据，并没有Client/Server这样说法。
RSocket完全采用TCP的通讯模型，也就是通讯双方是对等，可以相互发送请求并作出对应的响应，也就是对等通讯(Peer 2 Peer)，所以在RSocket中会使用Requester/Responder概念。

Peer 2 Peer的结构图如下，考虑到理解的方便，我们依然采用Server这一说法:

![Fire-and-Forget Diagram](/images/communication/p2p.png)

在RSocket中，通讯双方可以基于RSocket接口进行消息发送，同时也可以各自设置对应的Handler(acceptor函数接口)，完成对请求的响应，所以双方都是Requester和Responder。

* Responder监听方的请求响应接口:

```java

RSocketFactory
    .receive()
    .acceptor(((setup, sendingSocket) -> handler));
```


* Requester请求方的请求响应接口:
```java

RSocketFactory
    .connect()
    .setupPayload(DefaultPayload.create("metadata here"))
    .acceptor(handler2);
```

对等通讯的好处在于没有Client/Server模式的束缚，接下来我们会在RSocket Broker章节进行介绍。
