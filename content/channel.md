+++
title = "Channel - 双向流式通讯"
+++

Channel(通道)是在连接之上建立的一个虚拟的双向通讯的一个管道，通过该管道，通讯双方都可以发送和接收特定含义的消息。
举一个聊天软件的例子，我们创建一个文本Channel，用于聊天过程中的文本信息发送和接收；创建一个图像Channel，用于发送和接收聊天过程中的图片。
Channel创建后，接下来就是消息在通道中的双向流动，就是通常所说的同时发送和接收消息，如聊天场景中，聊天的内容都可以通过Channel发送和接收，从而实现IM的特性。
Channel的设计结构图如下:

![Channel](/images/communication/channel.png)

不少同学可能会问，Socket就具有这个特性啊，如WebSocket，为何还要引入Channel这样抽象层？ 主要有以下的考虑：

* 通常来说，一个Channel中的消息含义基本是固定的，处理的逻辑也也差不多，可能需要创建不同的Channel，如IM的场景，我们可能需要创建多个Channel:

    * 1对1聊天的Text Channel: 保证点对点聊天的实时性，这也是为何你可以看到聊天软件中会给你这样的即时提示，"对方正在输入中".
    * 群聊的Text Channel: 群聊的逻辑相对复杂点，如消息广播和推送策略，实时性可能不那么高
    * 聊天的Image Channel: 聊天中的图片发送和接收，图片都比较大，加上网络等问题，可能需要进行特殊处理，如先给你一个低质量的缩略图，然后再是原图。

* Channel的创建和关闭的成本非常低，不需要创建物理连接，这个也是WebSocket、Redis Streams做不到的，它们需要创建新的连接，如创建连接后，传输的数据并不多，完全是一种资源浪费。

在RSocket通讯协议中，Channel是内置的模型，你可以在RSocket接口中看到，接口如下::

```java

public interface RSocket extends Availability, Closeable {

  /**
   * Request-Channel interaction model of {@code RSocket}.
   *
   * @param payloads Stream of request payloads.
   * @return Stream of response payloads.
   */
  Flux<Payload> requestChannel(Publisher<Payload> payloads);

}

```

考虑到Channel相互通讯中，第一个message可能包含特殊的含义，如包含元信息、路由信息等，所以RSocket又增加了一个ResponderRSocket接口，可以将第一个消息单独提取出来处理。

```java

public interface ResponderRSocket extends RSocket {
  /**
   * Implement this method to peak at the first payload of the incoming request stream without
   * having to subscribe to Publish&lt;Payload&gt; payloads
   *
   * @param payload First payload in the stream - this is the same payload as the first payload in
   *     Publisher&lt;Payload&gt; payloads
   * @param payloads Stream of request payloads.
   * @return Stream of response payloads.
   */
  default Flux<Payload> requestChannel(Payload payload, Publisher<Payload> payloads) {
    return requestChannel(payloads);
  }
}

```

最后我们建议你可以参考一下Spring Blog上的 [Getting Started With RSocket: Spring Boot Channels](https://spring.io/blog/2020/04/06/getting-started-with-rsocket-spring-boot-channels) 文章。
