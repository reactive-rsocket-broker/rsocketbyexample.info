+++
title = "JWT - Token验证"
+++

JWT(JSON Web Token)是一种验证连接到服务端的用户身份机制，JWT中包含了客户端身份信息、额外属性和签名信息，更多的信息请参考 https://jwt.io/
在RSocket的安全验证机制中，请求方的身份验证可以通过JWT Token完成，通讯的安全通过TLS + JWT进行保证。 在具体的实时过程推荐以下做法:

* 默认开启Token验证: 如果不是开发环境，我们建议在产品环境中开启验证机制，这样可以保证只有可信的应用才能接入到RSocket系统中，而且方便后续的应用管理、服务治理和权限设置
* JWT的RSA验证: JWT Token验证通过RSA Public Key完成，而JWT token的生成由private key，这种分离策略可以保证Token的安全性
* 多租户问题: 建议在JWT Token中默认添加调用方ID，通常做法是不同应用或者服务分配不同的token
* 非关键性可以存放在JWT Token中: 如org ID，角色等，JWT Token是不可篡改的，可以保证数据的正确性，避免反复去查询数据库获取调用方信息，实现调用方信息自描述。

一些设计中会采用mTLS的设计，对比mTLS的方案，TLS + JWT会更简单些，而且很好地满足了安全的需求，避免了复杂的证书管理和维护等工作。

![Fire-and-Forget Diagram](/images/security/jwt.png)


# 参考

* RSocket Security Metadata: https://github.com/rsocket/rsocket/blob/master/Extensions/Security/Authentication.md
