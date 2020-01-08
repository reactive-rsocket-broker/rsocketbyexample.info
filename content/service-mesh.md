+++
title = "Service Mesh"
+++

谈到Service Mesh，大家可能马上会想到Istio + Envoy的SideCar的Service Mesh架构方案，这也是目前非常流行的做法。我们只需要在应用侧部署对应的通讯代理，如Envoy，接下来有Proxy负责完成和应用侧的Proxy的通讯，再完成Proxy到应用的通讯。

![Istio Envoy](/images/traffic/istio_envoy.png)

这样做没有什么问题，但是这里还是有一些值得思考的问题：

* Proxy代理的性能损耗: Proxy是独立的应用，需要消耗特定的资源，如CPU和内存，你可能很难想象，一个Node.js应用只需要64M内存，而应用身边的Envoy则需要1G的内存；另外整个通讯是代理的工作方式，网络调用也会有一定的延迟，主要存在于物理网络和经过Proxy的各种协议解析和路由。
* 架构的复杂性: 你需要Control Plane、Data Plane，不同应用的规则推送，Proxy之间的通讯安全性等等
* 运维成本的提升: 没有自动化的运维工具，你基本无法来部署Sidecar，当然目前Service Mesh的典型方案都是基于Kubernetes，这会减少不少工作量

那么有没有一种更好的方案来实现Service Mesh的特性？下面我们给出基于RSocket的Service Mesh架构方案。 RSocket Service Mesh方案是通过一个中心化的Broker完成的，典型的架构图如下:

![RSocket Service Mesh](/images/traffic/service_mesh.png)

RSocket Service Mesh的架构会带来什么样的变化？

* 中心化管理: 中心化会让管理更加方面，如我们介绍的Logging, Metrics，自定义Filter等，中心部署后就全网生效啦。 不少同学会担心中心化的性能瓶颈，我们前面介绍过，Broker是完全异步化的结构设计，性能高且稳定，这种软路由的设计能够保证中心化性能和稳定性不受影响
* 无SideCar: 性能损耗没有啦，运维简单啦。
* 无网络和基础设施依赖: 之前都是基于Kubernetes的Service Mesh方案，而Broker的中心化设计对网络和基础设施没有任何要求，在不同云服务厂商、边缘等都可以接入。
* 安全模型简单且统一: 请参考RSocket的安全设计，目前主要是TLS + JWT保障
* 其他: 如服务注册、RSocket协议对比HTTP的性能10倍提升、无端口监听保证等等

关于RSocket Service Mesh的更多资料可以参考 The New Service Mesh with RSocket: https://www.netifi.com/solutions-servicemesh

