+++
title = "Alibaba RSocket Broker - 基于RSocket新一代通讯中枢"
+++

我们在不同设计模式中都谈到了RSocket Broker发挥的核心作用，这里我们想阐述一下为何有Broker这个概念。 Broker是介于服务调用方和服务提供方之间，协调两者之间的服务通讯。
加入一个中间Broker后，最大的好处就是调用方和服务提供方之间的解耦。

* 调用方: 不用感知服务提供者的信息: IP，端口号，当前健康度状态，服务方服务器数量等等问题，请求方只要将请求发给Broker就可以啦
* 服务提供者: 服务启动后想Broker注册一下就就可以啦，不用通知服务注册中心、然后再有服务注册中心通知调用方。另外服务提供者只和Broker打交道，其他安全、流控等等全不用担心，这些Broker帮我搞定啦
* Broker中间人: Broker作为调用方和服务方的中间人，担当着服务注册、请求路由、安全、流量管控等角色，所有的信息都经过中间人，那么Observability和服务治理等都非常简单啦。

RSocket Broker典型的架构如下：

![RSocket Broker Diagram](/images/integration/alibaba_rsocket_broker.png)

RSocket Broker是一个中心化的结构，这个和网关类产品的结构设计差不多，称之为软路由设计。不少同学可能会担心Broker成为网路瓶颈，这里解释一下Broker的设计:

* 全异步架构设计: 非传统的ThreadPool的结构设计，None-Blocking，不会出现阻塞导致的雪崩场景。Broker在解析RSocket协议时，采用零拷贝(Zero-Copy)的设计，只需读取协议头部分字节就完成请求转发
* 服务注册: 由于所有的应用都会连接到Broker进行注册，但是每一个应用只会和broker建立一个长连接，应用的元信息都保存在内存中。关于长连接的数量，不同语言实现的Broker性能不太一样，如基于Java开发的Broker，单机连接总数在30-50万之间，而C++ Broker可以做到单机100万左右。如果有更多连接，目前RSocket Broker通过集群方案进行解决。
* 开放式Ops API: Broker不参与到具体的计算中，如限流策略、安全防护等，这些都是Broker将数据给其他计算应用，然后将计算的结果返回给Broker，所有Broker会提供非常多的Ops API，这个和路由设备管理都是类似的。

基于以上的架构设计，大家基本不用担心RSocket Broker集群会成为系统单点或者瓶颈。当然不同厂商的RSocket Broker实现还有各自的特点，可以参考具体的产品。 这里也将当前的RSocket Broker产品列一下，方便大家参考：


* Alibaba RSocket Broker: Alibaba开源产品 https://github.com/alibaba/alibaba-rsocket-broker
* Netifi Broker: 商业产品  https://www.netifi.com/
* Spring Cloud RSocket: Spring Cloud团队推出的RSocket Broker，目前孵化中 https://github.com/spring-cloud-incubator/spring-cloud-rsocket



