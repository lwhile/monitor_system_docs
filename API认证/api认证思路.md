# api认证实现思路

在已有的username, password方案的基础上,为每一个账号生成一个 AccessKey(AK), 以及一个对应的 SecretAccess(SK)

当 API 的调用者发出 HTTP 请求时,  需要对请求的内容进行处理, 处理的流程如下(参考百度云实现):

![](https://doc.bce.baidu.com/bce-documentation/Reference/GenerateKeyProcess06.png?responseContentDisposition=attachment)
