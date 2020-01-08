+++
title = "ACL - 权限验证"
+++

前面我们介绍了JWT机制，这个是身份验证，确保了应用能够接入到RSocket Broker，但是在具体的请求过程中，还会涉及到更细粒度的权限验证，如客户信息是非常保密的，不是所有的应用都能调用UserService，从而获得客户的email、手机和身份证信息等。
在RSocket的ACL设计中，我们引入了以下几个概念，当然这些信息都保存在JWT Token中，不会增加额外的查询工作。

* AppID: 应用唯一ID，你可以使用应用名称等，确保内部唯一
* TenentID: 应用所在的公司、组织或者部门ID，这个在后续的多租户设计中用处非常大
* Service Account: 这个概念来自Kubernetes的设计，同一个service account下的应用可以相互访问
* Role: 角色列表，如admin, ops等，可用于基于角色的权限控制，在Spring Security中可使用hasAnyRole(role1,role2)验证
* authority List: 权限列表,可用于基于具体权限的控制，如能否更新用户信息、能否访问用户信用卡信息等，在Spring Security中可使用hasAnyAuthority(authority1,authority2)验证

![Fire-and-Forget Diagram](/images/security/acl.png)

这些元信息，可以更好地帮助进行细粒度的权限验证，配合RSocket的Filter机制，你可以进行各种ACL Filter扩展。

目前Spring Security已经支持RSocket，关于JWT + ACL的规范整合还在进行中，希望大家持续关注。
