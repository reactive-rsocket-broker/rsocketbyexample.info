+++
title = "Request/Response - RPC"
+++


RPC(远程过程调用)是常见的通讯模式，也就是通讯的一方发出请求，而远程通讯的对方做出响应，也就是常说的请求/响应模式。服务通讯的一方发起创建连接的请求，在连接创建完毕后发起请求，服务方在接收到请求后，调用内部的处理逻辑，然后返回请求对应的结果，这个过程就是一个标准的RPC调用过程。
考虑到网络通讯中连接的创建是相对比较耗时的，所以大多数RPC通讯都是基于长连接的，这样可以避免频繁创建连接带来的巨大开销。出发稳定性和扩容的考虑，服务提供方可能是一个集群，由多台服务器对外提供服务，可以服务端在流量比较大的情况下可以快速扩容，从而保证系统的稳定性。
调用方出于稳定性和负载均衡的考虑，会和服务提供方创建多个连接，这样RPC调用会在不同的连接上进行，这样可以做到负载均衡；同时如果一个服务提供方失败不能继续提供服务，调用方也可以进行快速切换，保证系统的稳定性。 RPC的整体结构图如下:

![RPC Diagram](/images/communication/rpc.png)


在RSocket通讯协议中，RPC的实现是通过Request/Response模式完成的，也就是我们在RSocket接口中看到的以下API:

```java

public interface RSocket extends Availability, Closeable {

  /**
   * Request-Response interaction model of {@code RSocket}.
   *
   * @param payload Request payload.
   * @return {@code Publisher} containing at most a single {@code Payload} representing the
   *     response.
   */
  Mono<Payload> requestResponse(Payload payload);
}

```
