+++
title = "RSocket Java SDK"
+++

在RSocket众多开发语言SDK中，目前支持的最好就是Java SDK，所以这里我们介绍一下[RSocket Java SDK](https://github.com/rsocket/rsocket-java)。 RSocket Java SDK是基于Reactor这款Java Reactive框架开发，所以Reactor的一些编辑特性在SDK中都有体现。
另外RSocket得到Spring社区的大力支持，如Spring Framework 5.2+内置RSocket支持，Spring Boot 2.2+的RSocket Starter将RSocket开发变得更加简单。 相信不久你可以看多更多的产品会支持RSocket协议。

如果你想马上尝试一下RSocket，可以参考 [Getting Started With RSocket: Spring Boot Server](https://spring.io/blog/2020/03/02/getting-started-with-rsocket-spring-boot-server)

当然最快捷的方式尝试RSocket Java就是使用JBang，样例项目地址为： https://github.com/linux-china/jbang-rsocket

```java
///usr/bin/env jbang "$0" "$@" ; exit $?
//JAVA 8
//DEPS org.slf4j:slf4j-simple:1.7.36
//DEPS io.rsocket:rsocket-core:1.1.2
//DEPS io.rsocket:rsocket-transport-netty:1.1.2

import io.rsocket.Payload;
import io.rsocket.RSocket;
import io.rsocket.SocketAcceptor;
import io.rsocket.core.RSocketServer;
import io.rsocket.transport.netty.server.TcpServerTransport;
import io.rsocket.util.DefaultPayload;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Hooks;
import reactor.core.publisher.Mono;

public class ServerExample {
    public static void main(String[] args) {
        Hooks.onErrorDropped(e -> {});
        RSocket handler = new RSocket() {
            @Override
            public Mono<Payload> requestResponse(Payload payload) {
                System.out.println("RequestResponse: " + payload.getDataUtf8());
                return Mono.just(payload);
            }

            @Override
            public Flux<Payload> requestStream(Payload payload) {
                System.out.println("RequestStream: " + payload.getDataUtf8());
                return Flux.just("First", "Second").map(DefaultPayload::create);
            }

            @Override
            public Mono<Void> fireAndForget(Payload payload) {
                System.out.println("FireAndForget: " + payload.getDataUtf8());
                return Mono.empty();
            }
        };
        RSocketServer.create(SocketAcceptor.with(handler))
                .bindNow(TcpServerTransport.create("localhost", 7000))
                .onClose()
                .doOnSubscribe(subscription -> System.out.println("RSocket Server listen on tcp://localhost:7000"))
                .block();
    }
}
```

这里给出一些常见样例：

* Spring Tips: @Controllers and RSocket: https://spring.io/blog/2021/12/01/spring-tips-controllers-and-rsocket
* Easy RPC with RSocket: https://spring.io/blog/2021/01/18/ymnnalft-easy-rpc-with-rsocket
* Spring Boot with Kotlin and RSocket: https://spring.io/guides/tutorials/spring-webflux-kotlin-rsocket/
* GraphQL over RSocket: https://docs.spring.io/spring-graphql/docs/1.0.0-RC1/reference/html/#server-rsocket
* RSocket Load Balancing – Client Side: https://www.vinsguru.com/rsocket-load-balancing-client-side/
* RSocket Java Examples: https://github.com/gregwhitaker?utf8=%E2%9C%93&tab=repositories&q=rsocket+example&type=&language=
* Spring cloud function with RSocket: https://github.com/linux-china/spring-cloud-function-demo
* Spring Boot RSocket Demo with RPC style: https://github.com/linux-china/spring-boot-rsocket-demo
* Spring Framework with RSocket:  https://docs.spring.io/spring-integration/docs/current/reference/html/rsocket.html
* Spring Boot RSocket: https://docs.spring.io/spring/docs/current/spring-framework-reference/web-reactive.html#rsocket
* Spring Security RSocket: https://docs.spring.io/spring-security/site/docs/current/reference/html/rsocket.html
* Getting Started With RSocket: Spring Security https://spring.io/blog/2020/06/17/getting-started-with-rsocket-spring-security
* Spring Integration RSocket: https://docs.spring.io/spring-integration/docs/current/reference/html/rsocket.html#rsocket
* Alibaba RSocket Broker: https://github.com/alibaba/alibaba-rsocket-broker
* vlingo wire: https://github.com/vlingo/vlingo-wire
* RSocket Routing Broker: https://github.com/rsocket-routing/rsocket-routing-broker
* Spring Retrosocket: Feign-like or Retrofit-like experience for declarative RSocket-based clients https://github.com/spring-projects-experimental/spring-retrosocket
