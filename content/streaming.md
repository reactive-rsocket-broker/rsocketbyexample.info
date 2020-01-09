+++
title = "Streaming - 流式计算"
+++

在很多的业务场景中，我们都需要消费流式消息，然后对消息进行处理、计算和逻辑处理等，从而得出最终的数据结果或者业务处理结果。 典型的处理场景就是，应用要使用各个消息系统对应的SDK，如Kafka Client, RabbitMQ Client等，然后进行进行消息订阅。这种方式存在一定的问题：

* 应用要连接到消息系统，需要特定语言的SDK支持，如果你是使用一个新的语言，而该语言对应的消息系统SDK没有或者不稳定，这个时候就非常麻烦，你自己开发，那么难度还是非常高的。举一个例子，你现在想使用Rust开发应用，想使用Apache RocketMQ，那么不好意思，没有Rust Client，你需要等要不你自己开发。
* 如果对接多个消息系统，如Kafka、RabbitMQ、RocketMQ等，要使用各种SDK，对接不同的系统等，非常麻烦，同时也导致你系统臃肿不堪。
* API不统一: 不同的消息系统，API各不相同，要进行抽象等。

借助于RSocket的Request/Stream支持，让我们看一下新的结构会如何：

![Streaming](/images/integration/streaming.png)

在上述架构中，我们可以使用一个Streaming Broker帮助屏蔽各个消息系统对接的细节，然后由Streaming Broker对外提供流式消息。应用只需要通过RSocket协议连接到Streaming Broker就可以，这样做的好处有：

* 不用再对接各种消息系统啦，各种语言SDK等问题，只需要使用RSocket SDK就可以，当然RSocket SDK非常小，而且各种语言都支持，不用担心Rust连接不上RocketMQ的问题啦。一个微服务应用只为了发一条Kafka消息，结果引入了20M的jar包，这是多不划算的事情。
* API统一啦。 其实应用只关心要处理的流式消息，至于是从Kafka、RabbitMQ等，全部不用关心，只要发起一个Request/Stream请求就可以啦。
* 系统升级和切换容易： 如果你想升级Kafka集群，担心SDK兼容问题，需要协调各个应用升级，现在不用担心啦，直接搞定Streaming Broker和集群本身就可以啦，消息消费端无感。

Request/Stream可以更好支持现在Streaming、Pub/Sub场景，如果你有上述的一些担心，我们建议你使用尝试一下RSocket的Streaming集成方案，当然如果你消息消费场景不复杂，同时消息系统统一，那么之前的方案也是不错的选择。

