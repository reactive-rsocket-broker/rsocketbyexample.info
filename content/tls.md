+++
title = "TLS - 安全通讯"
+++

通讯安全已经是目前应用间通讯的最基本要求，RSocket是应用层的协议，这个和HTTP一样，所以非常容易为RSocket添加TLS支持。
RSocket一个推荐的设计是Broker架构，也就是通讯的双方都连接到Broker上，然后请求由Broker进行转发，这种设计的好处是只需要RSocket Broker安装了TLS证书，服务通讯的请求方和提供方都不需要进行TLS证书安装，而之前端口监听是需要安装证书的，某些情况下可能需要mTLS支持，证书管理等一堆工作，在RSocket的设计场景中是不需要的。

一句话: 只要Broker有TLS证书就可以啦，什么mTLS、证书管理等等，这些工作都不再需要啦。

![Fire-and-Forget Diagram](/images/security/tls.png)



