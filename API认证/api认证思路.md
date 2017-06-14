# api认证实现思路

在已有的username, password方案的基础上,为每一个账号生成一个 AccessKey(AK), 以及一个对应的 SecretAccess(SK)

当 API 的调用者发出 HTTP 请求时,  需要对请求的内容进行处理, 处理的流程如下(参考beego的认证流程):

1. 按照key值的顺序, 将HEADER中的某几个 Header 键值对组合成一个字符串 x .

2. 拼接待签名字符串 s : http.method + "\n" + x + "\n" + http.requestURL + "\n"

3. 使用sha-256算法对sk与s进行编码, 生成签名signature


接下来生成认证字符串:

sunrun-auth-v1/{AK}/{timestamp}/{expire}/{生成签名signature}

添加到 HTTP 请求 Header 的Authorization字段里

服务端接受到请求后, 从数据库读取AK和SK后, 采用相同的处理流程得出signature, 最后根据请求中的signature以及{expire}判断是否合法

[前端Javascript Demo](/API认证/demo.js) 


[后端Go Demo](/API认证/demo.go)