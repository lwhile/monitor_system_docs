# api认证实现思路

在已有的username, password方案的基础上,为每一个账号生成一个 AccessKey(AK), 以及一个对应的 SecretAccess(SK).SK绝对不可泄露

当 API 的调用者发出 HTTP 请求时,  需要对请求的内容进行处理, 处理的流程如下(参考beego的认证流程):

1. 按照key值的顺序, 将HEADER中的某几个 Header 键值对组合成一个字符串 b 

```go
var b bytes.Buffer
keys := make([]string, len(params))
pa := make(map[string]string)
for k, v := range params {
    pa[k] = v[0]
    keys = append(keys, k)
}

sort.Strings(keys)

for _, key := range keys {
    if key == "signature" {
        continue
    }

    val := pa[key]
    if key != "" && val != "" {
        b.WriteString(key)
        b.WriteString(val)
    }
}
```

2. 拼接待签名字符串 s : http.method + "\n" + b + "\n" + http.requestURL + "\n"

```go
stringToSign := fmt.Sprintf("%v\n%v\n%v\n", method, b.String(), RequestURL)
```

3. 使用sha-256算法对sk与s进行编码, 生成签名signature
```go
sha256 := sha256.New
hash := hmac.New(sha256, []byte(appsecret))
hash.Write([]byte(stringToSign))
```

接下来生成认证字符串(这一步beego源码里面没有, 而是参考了[相关参考-鉴权认证-简介-百度云](https://cloud.baidu.com/doc/Reference/AuthenticationMechanism.html)):

sunrun-auth-v1/{AK}/{timestamp}/{expire}/{signature}

添加到 HTTP 请求 Header 的Authorization字段里

服务端接受到请求后, 从数据库读取AK和SK后, 采用相同的处理流程得出signature, 最后根据请求中的signature以及{expire}判断是否合法

[前端Javascript Demo](/apiKey认证/demo.js) 



