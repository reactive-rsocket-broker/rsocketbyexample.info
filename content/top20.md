+++
title = "RSocket对Top 20编程语言的支持"
+++

RedMonk发布了2020年6月开发语言排名情况:  [The RedMonk Programming Language Rankings: June 2020](https://redmonk.com/sogrady/2020/07/27/language-rankings-6-20/)
相信不少同学都比较关注这个排行榜，想了解和学习一些热门语言，补充自己的知识。  RSocket是异步化高效通讯协议，当然少不了对多语言的支持，相信不少同学都比较关注这个排行榜，想了解和学习一些热门语言，补充自己的知识。  RSocket是异步化高效通讯协议，当然少不了对多语言的支持，
所以我们将前20名开发语言列一下，同时讨论一下RSocket对其支持情况。 相关的图标说明如：

* ⭐： RSocket的基本特性，是指RSocket的4+1通讯模型和Peer to Peer特性
* ⭐⭐： RSocket中等特性，主要Metadata和Encoding支持，能对接TCP和WebSocket、包括路由和基本的负载均的基本支持
* ⭐⭐⭐： RSocket高级特性则是指back pressure、Fragment、lease，resume、cancel等

### 1 JavaScript  ⭐⭐⭐
RSocket-JS很早就支持JavaScript啦，同时支持Node.js后端和浏览器前端，所以从浏览器端调用RSocket Service完全没有问题。 当然还有RxJS的对接，这个可能要有一个适配。我们打算在RxJS 7.0发布后，进行对应的适配。  RSocket JS另外一个问题就是实现有些复杂，完全基于JS开发，目前多数JS框架都是基于TS开发的，是否要用TS重新RSocket-JS，目前还没有确定。

### 2 Python ⭐⭐
目前RSocket Python的特性支持一般，主要是最基本的特性。 目前我们内部已经使用调整的一个版本，已经有半年以上，相对来说已经比较稳定，如果要更多的特性，这个还需要一定的开发。

### 3 Java ⭐⭐⭐
RSocket Java目前是支持最好的，特性和性能都非常好，Spring的工程师也在做起核心的维护，当然Reactor的开发团队也在做对应的支持。  考虑到RSocket Java完备的特性，如果你开发如RSocket Gateway，RSocket Broker这类产品，强烈建议使用Java开发，很节省你非常多的时间，而且有Spring Rsocket支持，相对也比较稳定。

###  4 PHP ⭐⭐
RSocket PHP在2020年七月完成了Alpha版本的开发，主要是基于ReactPHP，这个是PHP下非常知名基于EventLoop的异步框架。 目前的特性开发都已完成，基本没有问题。

### 5 C++ ⭐
RSocket CPP好久没有更新啦，目前RSocket CPP主要的问题是metadata的支持欠缺，目前我们在看Folly框架，很快会加入metadata的特性，其实最主要是ByteBuffer的使用。

###  6 C# ⭐⭐
C#的支持是通过RSocket.NET完成的，RSocket的基本特性都支持。

###  7 Ruby ⭐
RSocket Ruby，基于Sinatra思想设计，支持RSocket基本特性，目前考虑在重构代码，添加路由等特性，也让代码更容易维护。 目前的问题是RxRuby无人维护，另外EventMachine好久都不更新啦。 目前在考虑一些候选方案，如 concurrent-ruby 实现。

###  8 CSS
这个不会要求RSocket做些什么吧？  如果大家觉得Svelte组件模型对CSS支持不错，Svelte是和RSocket的RxJS适配是完全没有问题的。

###  9 TypeScript ⭐⭐
目前的TS支持主要是通过RSocket Deno来实现的，如果你有关注Deno + TS，那就不用担心啦，已经支持。 后续投入更多的精力在Deno的支持上。

### 10 C 🏗
之前打算基于 [Nano Msg](https://nng.nanomsg.org/) 开发对应的实现，目前打算基于Rust FFI来完成C的对接，这个在开发中，可行性还比较高。

### 11 Swift  🏗
RSocket Swift正在开发中，主要基于Swift NIO和Combine Framework，很快和大家见面。

### 12 Objective-C 🏗

目前的方案是 Importing Swift into Objective-C，也就是 Objective-C 调用Swift SDK相关的代码，这个会在Swift SDK开发完成后进行。

###  13 R  🔌
目前没有对应的规划。 如果你的R运行在GraalVM之上，那么R调用RSocket服务完全没有问题。

### 14 Scala  ⭐⭐⭐
调用RSocket Java即可

### 15 Go ⭐⭐
RSocket很早就增加对Golang的支持， 请访问 rsocket-go

### 15 Shell ⭐⭐⭐
RSocket包含对应的命令行程序，对Sell支持完全没有问题。 目前主要有两个CLI程序，分别是rsoket-cli和rsc，rsc支持Spring RSocket。

### 17 PowerShell ⭐⭐⭐
rsocket-cli 和 rsc都是基于Java开发的，另外rsc支持GraalVM Native，在PowerShell下没有问题。 当然如果有同学想基于RSocket .Net开发对应的CLI，也没有问题。

### 18 Perl 🔌
目前的方案是打算基于Rust FFI做，也就是perl to call Rust by FFI， 目前还没有启动。

###  19 Kotlin ⭐⭐⭐
如果是Kotlin JVM，那么直接RSocket Java就可以，而且Kotlin Coroutines和Reactive无缝对接。 如果是Kotlin Android，就需要使用rsocket-kotlin，目前是基本特性。  由于在规划Kotlin 和 RSocket更多的整合，如Kotlin/JS, Kotlin Native等支持，所以RSocket Kotlin在规划中，所以开发部太活跃。 但是大家都知道的，Kotlin丰富的语言特性，实现一些功能都比较容易，如果RSocket Kotlin也就欠缺一下metadata的支持。

### 20 Rust ⭐⭐
RSocket Rust开发已经有很长一段时间啦，主要基于Tokio开发，另外添加了对Rust WebAssembly支持。 考虑到WebAssembly、Deno和Rust本身的受欢迎程度，后续我们会投入较多精力，加强对Rust的支持。


### 总结
另外RSocket还对其他一些语言有支持，如RedMonk提及的Dart，也有对应的RSoccket实现，另外比较受欢迎的Julia，其本身就包含对Asynchronous Programming的支持，对应的实现也不太有问题。
