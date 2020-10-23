+++
title = "RSocket Kotlin SDK"
+++

RSocket Kotlin是基于kotlinx.coroutines的多平台的实现，也就是说使用RSocket-Kotlin，各个Kotlin平台，如Kotlin JVM, Kotlin/JS, Kotlin/Native等都可以和RSocket对接。
对应RSocket JVM来说，同时还支持Android接入，你的Android App可以直接使用RSocket Kotlin然后访问RSocket服务。
目前RSocket Kotlin主要支持以下一些平台和对应的传输层：

* JVM: 支持Client/Server端的TCP/WebSocket通讯，主要是Java Backend
* Android: 支持ClientTCP/WebSocket通讯
* JS Node.js: 支持Client/Server的TCP/WebSocket
* JS Browser: 浏览器端的WebSocket接入
* Native: 支持Client/Server的TCP通讯，主要平台为：linux x64, macos, ios, watchos, tvos，目前还不支持 windows x64 yet

另外RSocket Kotlin主要RSocket的5种通讯模型，另外支持Stream之上的被压支持。
