+++
title = "DDD + Reactive"
+++

领域驱动设计（DDD）是一个开放的软件设计方法论，从 Eric Evans 出版《领域驱动设计》之后，DDD 一直是业内推崇的设计方法学，其划分服务的Bounded Context理念已经被微服务设计所接受，而且在微服务架构设计中也是非常推崇DDD的Bounded Context的理念进行应用拆分。
但是在各个Bounded Context之间如何通讯，DDD给出的推荐方案是基于消息/事件的通讯设计，但是具体使用什么样的技术、什么样的协议来解决这个问题，DDD并没有给出明确的答案，可以说只是一个指导思想。

![DDD Context Map](/images/misc/ddd_context_map.png)

我们都知道Bounded Context之间的通讯是复杂的，只是使用RPC, Pub/Sub模型很难解决通讯中存在的问题，而RSocket的丰富通讯模型，完全可以满足Context之间个各种通讯需求，让我们看一下如何解决的。

* 基于异步消息通讯: 方便解耦，性能高
* 4个通讯: RPC, Pub/Sub, Channel, Fire-and-Forget各种场景全满足
* Metadata Pus: 元信息推送，如果Context存在元信息依赖场景，完全没有问题
* 对等通讯: Context之间是对等的，相互调用没有问题

DDD + Reactive + RSocket的结合，就很好地解决了Context(微服务)之间如何相互通讯的问题，将指导思想能够通过具体的技术进行落地，这个也是DDD倡导的从问题域到实现域一整套解决思路。
