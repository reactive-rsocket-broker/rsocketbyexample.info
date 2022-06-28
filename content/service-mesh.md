+++
title = "Service Mesh"
+++

谈到Service Mesh，大家可能马上会想到Istio + Envoy的SideCar的Service Mesh架构方案，这也是目前非常流行的做法。我们只需要在应用侧部署对应的通讯代理，如Envoy，接下来有Proxy负责完成和应用侧的Proxy的通讯，再完成Proxy到应用的通讯。

![Istio Envoy](/images/traffic/istio_envoy.png)

这样做没有什么问题，但是这里还是有一些值得思考的问题：

* Proxy代理的性能损耗: Proxy是独立的应用，需要消耗特定的资源，如CPU和内存，你可能很难想象，一个Node.js应用只需要64M内存，而应用身边的Envoy则需要1G的内存；另外整个通讯是代理的工作方式，网络调用也会有一定的延迟，主要存在于物理网络和经过Proxy的各种协议解析和路由。
* 架构的复杂性: 你需要Control Plane、Data Plane，不同应用的规则推送，Proxy之间的通讯安全性等等
* 运维成本的提升: 没有自动化的运维工具，你基本无法来部署Sidecar，当然目前Service Mesh的典型方案都是基于Kubernetes，这会减少不少工作量

那么有没有一种更好的方案来实现Service Mesh的特性？下面我们给出基于RSocket的Service Mesh架构方案。 RSocket Service Mesh方案是通过一个中心化的Broker完成的，这里我们将Sidecar和RSocket架构进行一个对比，如下:

![RSocket Service Mesh](/images/traffic/sidecar-vs-rsocket.png)

这里列举一下两个架构方案典型特征的对比：

* 基础设施层：两者都是，但是部署结构不太一样，一个是sidecar proxy + control plane，一个中心化broker并集成了control plane功能
* 服务到服务间的通讯： 两者都是在协调服务间通讯，不同的是sidecar proxy要适配各种通讯协议，而RSocket则通过自身丰富的通讯模型，来实现其他协议提供的功能。
* 应用或者设备接入： 并不是所有的设备都可以安装Proxy的，主要有几种原因：设备和系统本身就不支持，如IoT设备；另外一个是不经济，浪费钱： 如你用Node.js写一个function，只需要128M内存+1核CPU，如果要应用侧安装Envoy，那么1G内存可能就没有啦，卧榻之侧，岂容他人鼾睡； 特殊的网络环境，如在Edge端，要从云端control plane来反向控制Edge端的Proxy，那么可能要进行各种网络打通，设备，VPN，安全等等，相对非常复杂。
* 运维成本增加： 管理中心化的20台服务器RSocket Broker集群，和管理10K个Proxy实例是不一样的。 就好比我们在实现微服务架构时，如果你还没有容器化，那最好不要微服务，众多的应用数量会让你非常头痛，无容器不微服务。
* FAST: RSocket是异步化的， 异步化虽然没有解决请求响应的时长问题，但是异步化能提升系统的处理能力，大家不需要等待，理论上也是更快啦。 在SideCar的模式下，如果没有介入异步化处理，必然导致响应时间更长，系统就更慢啦。 即便你完成了Sidecar Service Mesh架构，你还是需要做异步化，不然TCP多跳和Proxy层协议解析，只会让系统更慢。 上述的架构图中，Sidecar是TCP 6跳，虽然其中4跳是本机内部进行的，额外2次协议解析，而RSocket是TCP 4跳，额外1次或者0次协议解析。
* 安全：这方面只能说在同样的安全要求下，RSocket的安全实现更简单。Broker目前主要是TLS + JWT，而非mTLS，不需要证书管理。 同时借助于JWT的安全模型，可以轻松实现更细粒度的权限控制。同时RSocket服务是无端口监听的，不配置端口监听对应的TLS证书，同时也避免了一定的网络攻击。
* RELIABLE：消息通讯可靠性这个大家都是能明白的，重新、序列化、转发投递等都非常容易，大家参考一下Reactive宣言，就明白Reactive架构设计的出发点。
* 无网络要求： sidecar还是采用传统的端口监听模式对外提供服务，而RSocket对外提供服务完全是无端口监听的，只要能能连接上Broker，任何服务就可以通讯，没有任何网络要求，iptables配置等。
* 网络交互模型： 如果要你采集应用的metrics怎么做？如果该应用在Edge端该怎么办？ Sidecar Service Mesh模型对网络还是非常强的依赖，你要做网络打通和安全保障等，这样应用间才能相互访问。 而Broker的介入，你完全不用考虑网络的问题，你完全不用知道应用在哪里，要采集metrics等，直接发给broker，broker帮你和应用进行通讯，然后拿回你需要的数据。RSocket的网络模型中，通常不允许其他系统直接访问应用，存在很大的安全风险，同时需要网络配合。

RSocket Service Mesh的架构会带来什么样的变化？

* 中心化管理: 中心化会让管理更加方面，如我们介绍的Logging, Metrics，自定义Filter等，中心部署后就全网生效啦。 不少同学会担心中心化的性能瓶颈，我们前面介绍过，Broker是完全异步化的结构设计，性能高且稳定，这种软路由的设计能够保证中心化性能和稳定性不受影响
* 无SideCar: 性能损耗没有啦，运维简单啦。
* 无网络和基础设施依赖: 之前都是基于Kubernetes的Service Mesh方案，而Broker的中心化设计对网络和基础设施没有任何要求，在不同云服务厂商、边缘等都可以接入。
* 安全模型简单且统一: 请参考RSocket的安全设计，目前主要是TLS + JWT保障
* 其他: 如服务注册、RSocket协议对比HTTP的性能10倍提升、无端口监听保证等等

关于Istio, eBPF和RSocket Broker的综合对比，可以参考 [《Istio, eBPF and RSocket Broker: A deep dive into service mesh》](https://medium.com/geekculture/istio-ebpf-and-rsocket-broker-a-deep-dive-into-service-mesh-7ec4871d50bb)。
