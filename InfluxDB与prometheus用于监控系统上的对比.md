# InfluxDB与Prometheus用于监控系统上的对比


在把玩两套监控系统, 对于 InfluxDB 与 Prometheus 都有一定程度的了解后, 尝试对这两个产品做一个横向对比.


### 总览
首先要明白, Prometheus 提供的是一整套监控体系, 包括数据的采集, 数据存储,报警, 甚至是绘图(只不过很烂,官方也推荐使用 grafana).
而 InfluxDB 只是一个时序数据库, 使用它做监控系统的话, 还需要物色数据采集器,如 telegraf, collectd 等. 甚至连报警模块, 也需要使用同为 Influxdata 公司出的 Kapacitor.从这个角度来说, Prometheus 会有运维上的优势, 因为它用起来确实很省事. 

### 数据的采集
Prometheus 和 InfluxDB 在数据的采集上两者就选择了不同的极端, 前者只能 pull, 后者只能 push, 关于 pull 和 push 的对比,这里就不多述了.
Prometheus 把 数据的采集器叫做 exporter, xxx-exporter 运行之后会在机器上占用一个端口, 等待 Prometheus server 拉取数据.

InfluxDB 的数据采集器我们使用了 telegraf, 官方宣传插件化驱动, 其实也就那么回事, 编译的时候把很多东西的采集功能包进去压在一个二进制里面了. 功能其实是蛮强大了. 也是 Go 编写, 部署算不上困难, 但用起来总会觉得多了一点什么东西.

存在感.



 