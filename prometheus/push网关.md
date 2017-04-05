# push网关

    package pushgateway

    import (
        "fmt"
        "io/ioutil"
        "net/http"
        "strings"
    )

    var (
        countM map[string]int
    )

    func init() {
        countM = make(map[string]int)
    }

    // TestController :
    func TestController(resp http.ResponseWriter, req *http.Request) {
        ctrlName := "TestController"
        req.ParseForm()
        countM[ctrlName]++
        sendPushgateway()
    }

    func sendPushgateway() {
        client := &http.Client{}

        sendValue := "some_metrics_test{controller=\"test\"} 1243\n"
        req, err := http.NewRequest("POST", "http://127.0.0.1:9091/metrics/job/job_test/instance/instance_test", strings.NewReader(sendValue))
        req.Header.Set("Content-Type", "text/plain;charset=utf-8")
        if err != nil {
            fmt.Println(err)
        }

        resp, err := client.Do(req)
        if err != nil {
            fmt.Println(err)
        }
        fmt.Printf("resp:%v\n", resp)

        body, err := ioutil.ReadAll(resp.Request.Body)
        if err != nil {
            fmt.Println(err)
        }
        fmt.Println(body)
    }


prometheus提供一个push网关让一些监控指标以push的方式输出到网关上,再有prometheus server去网关抓取,满足一些特殊的场景,比如无法新开一个供pull使用的端口,或者为了安全考虑不允许本机被其他网络访问.

使用该push网关的方式非常简单,只需要将metric输出到网关上即可,网关的默认端口是9091.
有两种方式可以推数据,一种是引入官方的库,使用方式和pull的方式差不多,由库将指标透明化得输出到网关上.
一种则是在业务代码中维护metrics,自己控制push的逻辑.

鉴于使用官方提供的库需要先了解prometheus的基本数据类型以及使用方法,于是我做了一层最简单的封装,降低使用的门槛,将监控指标精细到业务上.当然这样的缺点也是显而易见的,封装后的代码通用性会降低.

目前可以提供的类型有:

  - Go语言的HTTP API访问统计



## 数据的输出格式

有两种输出格式可选:

- text

    \# HELP api_http_request_count The total number of HTTP requests.

    \# TYPE api_http_request_count counter

    http_request_count{method="post",code="200"} 1027 1395066363000

    http_request_count{method="post",code="400"}    3 1395066363000

    \# Escaping in label values:

    msdos_file_access_time_ms{path="C:\\DIR\\FILE.TXT",error="Cannot find file:\n\"FILE.TXT\""} 
    1.234e3
    
    \# Minimalistic line:
    
    metric_without_timestamp_and_labels 12.47
    
    \# A weird metric from before the epoch:
    
    something_weird{problem="division by zero"} +Inf -3982045
    
    \# Finally a summary, which has a pretty complex representation in the text format:
    
    \# HELP telemetry_requests_metrics_latency_microseconds A histogram of the response latency.
    
    \# TYPE telemetry_requests_metrics_latency_microseconds summary
    
    telemetry_requests_metrics_latency_microseconds{quantile="0.01"} 3102
    
    telemetry_requests_metrics_latency_microseconds{quantile="0.05"} 3272
    

- protocol-buffer

**2014年4月之后的prometheus版本都大于0.0.4,其pushgateway不支持JSON**


注意事项:

- 编码格式必须为utf8
- 每一行必须以"\n"做结尾
- 以任意数量的空格或\t隔开一条监控指标的内容
- 首尾的空格会被忽略
- 注释以#开头,并且必须以"HELP"或者"TYPE"做第二个关键字,todo:补充更加详细的文档!
- 值为float类型,Nan,+Inf,-Inf表示值不可用
- 时间戳可以自己添加,类型为int64(微妙).若不添加则默认为prometheus server的抓取时间.
[为什么以抓取时间而不是采集时间]()