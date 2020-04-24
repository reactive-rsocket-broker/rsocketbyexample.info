+++
title = "ACL - 访问控制列表"
+++

前面我们介绍了JWT机制，JWT Token是保证了是身份验证(Authentication)，确保了应用能够接入到RSocket Broker然后调用其他服务，但是在具体的服务调用过程中，还会涉及到更细粒度的访问权限控制(Authorization)，如客户信息是非常保密的，不是所有的应用都能调用UserService，从而获得客户的email、手机和身份证等隐私信息。
在RSocket的ACL设计中，我们引入了以下几个概念，当然这些信息都保存在JWT Token中，不会增加额外的查询工作。

* AppID: 应用唯一ID，你可以使用应用名称等，确保内部唯一
* OrgID: 应用所在的公司、组织或者部门ID，这个在后续的多租户设计中用处非常大，在通常的安全设计中，只有同一组织下的服务才能相互调用，不同组织下的服务是隔离。
* Service Account: 这个概念来自Kubernetes的设计，同一个service account下的应用可以相互访问
* Role: 角色列表，如admin, ops等，可用于基于角色的权限控制，在Spring Security中可使用hasAnyRole(role1,role2)验证
* Authority List: 权限列表,可用于基于具体权限的控制，如能否更新用户信息、能否访问用户信用卡信息等，在Spring Security中可使用hasAnyAuthority(authority1,authority2)验证
* 验证方式: RSocket Broker采用的是基于RSA算法的JWT验证，也就是私钥控制JWT Token生成，而公钥负责JWT Token验证，可以分发给应用使用，不需要安全验证中心和远程API验证，且无秘钥泄露风险。

![ACL Diagram](/images/security/acl.png)

以上这些元信息，可以更好地帮助进行细粒度的权限控制，配合RSocket的Filter机制，你可以进行各种ACL Filter扩展。 目前Spring Security已经支持RSocket，你只需要少量的配置就可以完成RSocket和Spring Security集成，实现访问权限控制。

**友情提示**: 关于JWT Token中是否要包含permission claims, 有不同的观点，这里我们不进行讨论。 JWT token失效的问题，这个建议使用一个黑名单列表，这个在RSocket设计中非常容易做到。
