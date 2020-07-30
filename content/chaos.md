+++
title = "Chaos - 混沌测试"
+++

在应用向微服务迁移的工程中，测试整个系统的故障容忍能力，已经变得非常重要啦。当然混沌测试涉及面非常广，如网络及其设备、数据、底层基础架构、服务间调用等。
这里我们主要讲述一下RSocket对Chaos的一些良好支持，当然其中不少特性来自于Reactive的Operator方法。

* Latency(Delay): 这个不用说啦，Reactive天生就支持这一特性，一个delaySequence API就可以搞定
* Down(网络不可用): 我们只需要调用Mono.error() 或者 Flux.error() 就可以马上模拟网络不可用的异常场景
* RSocket Interceptor和Filter: RSocket Java SDK中提供了非常好的Interceptor插件机制，只需一个Interceptor就可以模拟网络、服务调用、超时等各种情况。 在RSocket Broker的设计中，通过Filter机制，也能模拟各种场景，如权限不足等等。你只需要创建ChaosInterceptor和ChaosFilter类进行扩展就可以啦。
* Reactive的错误和异常操作方法:  Reactive提供了非常多关于错误处理的方法，如错误消费、错误重试、错误转换等等，这些多可以方便你进行相关的错误处理。

[Toxiproxy](https://github.com/Shopify/toxiproxy)提供了不少模拟Chaos的测试，如latency、down、bandwidth、slow_close、timeout等等，有兴趣的同学可以参考一下。

这里不是想说明RSocket天生为Chaos设计的，而是RSocket让服务间调用的混沌测试更简单，在不借助外部工具的情况下，我们通过Reactive+Mock的方式就可以很快做的混沌测试的效果。


