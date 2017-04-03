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