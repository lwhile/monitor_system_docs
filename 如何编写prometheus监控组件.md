# 如何编写client library

prometheus采集数据有两种方式,一种是通过job,一种是通过exporter.使用job需要开发者自己开发client library.使用exporter除了自己开发,目前社区也有不少exporter可以做到开箱即用.

下面先介绍如何使用Go开发适合web请求的client library.

prometheus将监控数据分为四种类型:

## Counter:
Counter类型只允许单调递增(正如其名字一样,这是一个只能进行加操作的计数器).其绝对禁止进行减操作,但可以被重置(比如服务被重新启动)
Counter拥有下列两个函数:

    inc():对计数器加1
    inc(v float64):对计数器加v,v>=0

Counter必须从0开始计数.

下面的例子演示如何统计web请求的次数:

    package main

    import (
        "fmt"
        "net/http"

        "github.com/prometheus/client_golang/prometheus"
        "github.com/prometheus/client_golang/prometheus/promhttp"
    )

    // 针对全局的counter
    var (
        counter prometheus.Counter
    )

    func sayHello(resp http.ResponseWriter, req *http.Request) {
        // 访问该handler时对计数器加1
        counter.Inc()
        fmt.Fprint(resp, "Hello.")
    }

    func main() {
        // 初始化counter
        counter = prometheus.NewCounter(prometheus.CounterOpts{
            Name: "api_count",
            Help: "web api counter",
        })
        // 注册该counter,这一步是必须的
        prometheus.MustRegister(counter)
        http.HandleFunc("/", sayHello)
        http.Handle("/metrics", promhttp.Handler())
        err := http.ListenAndServe(":8081", nil)
        if err != nil {
            panic(err)
        }
    }

> 使用ab工具发出请求: ab -n 100 127.0.0.1:8081/

![](https://raw.githubusercontent.com/lwhhhh/monitorDoc/master/asset/images/webAPI.png)


## Gauge:

Gauge在Counter的基础上,取消了单调递增的限制,允许增加或减少,以及直接设置.

    inc(): 进行加1操作
    inc(v double): 进行加v操作
    dec(): 进行减1操作
    dec(v double): 进行减v操作
    set(v double): 将gauge的值设置为v

Gauge的基本使用与Couter类似,用法可参考Counter.

## Summary:
待续

## Histogram: