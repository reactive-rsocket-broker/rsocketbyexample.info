+++
title = "metadataPush - 元信息推送"
+++

在通讯的模型中，不仅仅是各种请求的模型(RPC, Pub/Sub, Channel等)，还存在着管理(Ops)的需求，如告知通讯对方我即将下线重新发布，也就是优雅下线；如果连接到服务集群，服务存在扩容和缩容的需求，集群的拓扑结构发生变更，需要应用重新做负载均衡调整。
当然还有一个更常见的场景，就是配置推送，如Database, Redis服务器迁移，希望应该能够连接到新的服务器上，如果大家对Spring Config Server有了解的话，就会对这个场景非常熟悉。

Metadata Push的整体结构图如下:

![Fire-and-Forget Diagram](/images/communication/metadata_push.png)

也就是在连接到目标Server或者Broker(中间服务器)，Server或者Broker会源源不断推送相关的配置变更信息。

在RSocket中，metadataPush是内置的，你可以在RSocket接口中看到，接口如下:

```java

public interface RSocket extends Availability, Closeable {

  /**
   * Metadata-Push interaction model of {@code RSocket}.
   *
   * @param payload Request payloads.
   * @return {@code Publisher} that completes when the passed {@code payload} is successfully
   *     handled, otherwise errors.
   */
  Mono<Void> metadataPush(Payload payload);
}

```

在处理配置推送的场景中，我们还需要有以下的一些考量:

* 优先级: 配置推送不同于普通的消息，通常都是和Ops相关的，具有最高的发送优先级。在RSocket SDK实现中，METADATA_PUSH消息具有最高的发送优先级，不会进入消息发送队列，从而影响发送时效。
* 消息格式: 配置推送信息通常有两个最重要的字段： 配置类型和数据格式。不同的类型会有不同的处理逻辑；对应的配置数据也是多样的，如普通文本、properties格式、init格式、json格式等等。 这里建议使用采用CloudEvents规范(https://cloudevents.io/)，对接入方也比较方便。

