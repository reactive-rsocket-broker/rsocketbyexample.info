+++
title = "Request/Stream - Pub/Sub"
+++


Pub/Sub(发布/订阅)是基于消息的发布订阅模式，也是消息中间件典型的通讯模式。订阅方发起请求，然后发送方会源源不断地给订阅方发送消息。如在消息中间件系统中，典型的场景就是基于Topic的消息订阅。
通讯模型如下：

![Request/Stream Diagram](/images/communication/stream.png)

在Pub/Sub模型中，有两种典型的消息消费方式:

* Push: 发送方在消息生成后马上推送给订阅方，从而完成消息的消费。
* Pull: 订阅方会基于位点(offset)的方式，不断从消息生成方以小批量的方式获取数据，然后进行逻辑处理。 消费方要负责轮询和位点保存等

Push的模型并没有对消息方进行保护，在极限的情况下会导致消费方压力过大，不能即时处理消息，而Pull模型则可以很好地更加自我的处理能力动态调整处理消息的速度。

在RSocket通讯协议中，Pub/Sub的实现是通过Request/Stream模式完成的，也就是我们在RSocket接口中看到的以下API:

```java

public interface RSocket extends Availability, Closeable {

  /**
   * Request-Stream interaction model of {@code RSocket}.
   *
   * @param payload Request payload.
   * @return {@code Publisher} containing the stream of {@code Payload}s representing the response.
   */
  Flux<Payload> requestStream(Payload payload);
}

```

另外在RSocket规范中，引入了被压(Back Pressure)的概念。 被压其实就是有限制的Push模型。
消息订阅方发起订阅同时告知接下来要请求的消息最大数量(Request N)，消息发送方在发送消息时会引入计数器，保证推送的消息不超过最大的消息数量，达到最大推送数量后，消息发送方不会再发送消息。
订阅方在本批N个消息处理完毕后，会再发起一个Request N的请求，随后消息发送方再次进行消息发送。Back Pressure流程如下:

![Request/Stream Diagram](/images/traffic/back_pressure.png)

总体来说 Back Pressure =  Request N + Push


### Push, Pull 和 Back Pressure 模型对比

|     | Push               | Pull               | Back Pressure            |
|:----|:-------------------|:--------------------|:-------------------|
| 复杂度 | 低     | 高  | 中 |
| 性能            |  高 | 中      | 高 |
| 订阅方保护 | :red_circle:       | :white_check_mark:  | :white_check_mark: |



