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
