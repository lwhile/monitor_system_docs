# 测试部署多个prometheus server节点

## 测试内容

使用docker运行两个prometheus容器,监控node_exporter输出的系统指标

## 关注的结果

- 两个prometheus容器获取到的数据是否一致
- 关闭其中一个docker容器查看另外一个容器是否能继续正常运行
- 查看prometheus server存储的数据源是否有多个


## 结构图

![](https://raw.githubusercontent.com/lwhhhh/monitorDoc/master/asset/images/multiple_prometheus_server.png)

## 测试过程

分别使用以下命令启动两个prometheus容器

> docker run -p 9090:9090 -v ~/Go/src/github.com/prometheus/prometheus/:/etc/prometheus/ prom/prometheus

> docker run -p 9091:9090 -v ~/Go/src/github.com/prometheus/prometheus/:/etc/prometheus/ prom/prometheus
 
 **容器A开放9090端口到宿主机的9090端口,容器B开放9090端口到宿主机的9091端口,两个容器都要挂载prometheus配置文件所在文件夹到容器内**

 配置文件prometheus.yaml为:

    global:
        scrape_interval:     5s
        evaluation_interval: 5s

    scrape_configs:
    - job_name: linux
        static_configs:
        - targets: ['192.168.1.185:9100']
            labels:
            instance: linux1


接下来运行测试程序,让CPU和内存的使用率明显增加:

    func main() {
        // 生成50万个goroutine
        num := 100 * 100 * 50
        end := make(chan struct{})
        for i := 0; i < num; i++ {
            go func(i int) {
                if i == num-1 {
                    end <- struct{}{}
                }
                fmt.Println("goroutine#", i)
            }(i)
        }
        select {
        case <-end:
            fmt.Println("finish.")
        }
    }


运行grafana,两个prometheus server得到的数据如下:

![](https://raw.githubusercontent.com/lwhhhh/monitorDoc/master/asset/images/dockerA1B1.png)


数据的波动曲线和预期一致.

接下来关闭其中一个docker容器,模拟其中一个prometheus server挂掉

得到的数据为:

![](https://raw.githubusercontent.com/lwhhhh/monitorDoc/master/asset/images/dockerA2B2.png)

可以看到右边的grafana已经停止了刷新,而左边的grafana依旧能正常获取node_exporter输出的数据


重新启动docker容器,grafana开始重新绘制

![](https://raw.githubusercontent.com/lwhhhh/monitorDoc/master/asset/images/dockerA3B3.png)


prometheus server将采集到的数据存放到磁盘上,默认存放路径是运行目录下的data文件夹,而且一个prometheus server单独拥有一个保存数据的存储源.
关于prometheus如何存储采集到的数据可以查看:

[prometheus是如何存储数据的](https://github.com/lwhhhh/monitorDoc/blob/master/prometheus是如何存储数据的.md)
