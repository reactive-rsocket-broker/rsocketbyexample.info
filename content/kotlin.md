+++
title = "RSocket Kotlin SDK"
+++

RSocket Kotlin是基于kotlinx.coroutines的多平台的实现，我们知道Kotlin多平台，主要涉及如Kotlin JVM、Kotlin/JS、Kotlin Mobile和Kotlin/Native等，
而RSocket-Kotlin则可以让这些Kotlin应用可以通过RSocket无缝对接，当然访问其他语言开发的RSocket服务也没有问题。此外Kotlin Coroutines和Flow都是异步化的，
这个和RSocket这样异步化消息通讯协议是完全匹配的，可以说RSocket和Coroutines/Flow完全是一体的，没有任何违和感。

![RSocket Kotlin Stack](/images/language/rsocket-kotlin-stack.png)

目前RSocket Kotlin主要支持以下一些平台和对应的传输层：

* JVM: 支持Client/Server端的TCP/WebSocket通讯，主要是Java Backend
* Android: 支持ClientTCP/WebSocket通讯
* JS Node.js: 支持Client/Server的TCP/WebSocket
* JS Browser: 浏览器端的WebSocket接入
* Native: 支持Client/Server的TCP通讯，主要平台为：linux x64, macos, ios, watchos, tvos，目前还不支持 windows x64 yet

另外RSocket Kotlin支持RSocket的5种通讯模型，另外支持Stream之上的被压支持，更多的信息请访问： https://github.com/rsocket/rsocket-kotlin
