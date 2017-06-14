# api认证实现思路

在已有的username, password方案的基础上,为每一个账号生成一个 AccessKey(AK), 以及一个对应的 SecretAccess(SK)

当 API 的调用者发出 HTTP 请求时,  需要对请求的内容进行处理, 处理的流程如下(参考beego的认证流程):

1. 按照key值的顺序, 将HEADER中的所有 key + value 组合成一个字符串 x .

2. 拼接待签名字符串 s : http.method + "\n" + x + "\n" + http.requestURL + "\n"

3. 使用sha-256算法对sk与s进行编码, 生成签名signature

