+++
title = "RSocket JavaScript SDK"
+++

RSocket JavaScript SDK主要包括三方面的支持: 浏览器到RSocket服务后端调用、Node.js后端到RSocket服务调用，以及Deno TypeScript后端到RSocket服务调用。
通讯协议主要支持TCP Socket和WebSocket，浏览器到RSocket后端的调用，通讯协议只能为WebSocket。当前只支持JSON数据格式，其他的数据格式需要自行扩展。
对Deno的支持，是通过rsocket-deno SDK提供，该SDK由TypeScript开发，只适用Deno。
你需要根据不同的业务场景选择对应的开发包和通讯方式。

RSocket JS技术栈整体结构图如下：

![RSocket Tags Cloud](/images/language/rsocket-js-stack.png)

RSocket JS对应的SDK参考如下：

* RSocket JS: https://github.com/rsocket/rsocket-js
* RSocket Deno: https://deno.land/x/rsocket

*友情提示:*  如果也想尝试Web UI层的Reactive技术，这里我们推荐使用[Svelte](https://svelte.dev/)，是一款完全Reactive的前端框架，我们提供了对应多的Demo，你可以在这里 https://github.com/linux-china/svelte-rsocket-demo 查阅。
