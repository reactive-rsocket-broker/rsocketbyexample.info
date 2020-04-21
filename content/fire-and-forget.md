+++
title = "Fire and Forget - 无回执数据发送"
+++

在网络通讯中，存在着不需要接收方回执确认的调用模型，如数据采集的场景: 打点采集、日志传输、metrics上报等，由于这些数据通常是一些非关键性数据，发送出去后接收方不需要返回确认收到结果，即便数据丢失啦，也是可以接受的，这个类似于UDP的通讯模型。
在RSocket的协议设计中，这种场景是通过Fire-and-Forget实现的，该模型的好处在于性能极致，同时不会给应用增加过多的负担。Fire-and-Forget的整体结构图如下:

![Fire-and-Forget Diagram](/images/communication/fire_and_forget.png)


在RSocket通讯协议中，Fire-and-Forget是内置的模型，你可以在RSocket接口中看到，接口如下:

```java

public interface RSocket extends Availability, Closeable {

  /**
   * Fire and Forget interaction model of {@code RSocket}.
   *
   * @param payload Request payload.
   * @return {@code Publisher} that completes when the passed {@code payload} is successfully
   *     handled, otherwise errors.
   */
  Mono<Void> fireAndForget(Payload payload);
}

```

最后我们建议你可以参考一下Spring Blog上的 [Getting Started With RSocket: Spring Boot Fire-And-Forget](https://spring.io/blog/2020/03/16/getting-started-with-rsocket-spring-boot-fire-and-forget) 文章。


