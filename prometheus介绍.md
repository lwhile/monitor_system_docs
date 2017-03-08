# prometheus

项目地址: [Github](https://github.com/prometheus/prometheus)

官网: [prometheus.io](https://prometheus.io)


## 简介:

Prometheus由SoundCloud开发并完全开源,使用Go语言编写.项目开始于2012年,2016年成为继Kubernetes之后,第二个加入Cloud Native Computing基金会的项目

## 特点:

- 多维度数据模型
- 灵活的查询语言
- 不依赖于分布式存储;单个服务节点自治
- 基于http,使用pull的方式采集时序数据
- 支持采用独立网关的模式实现push模型
- 可通过服务发现或者静态配置的方式发现target
- 支持多种数据展现方式
- 容器化支持很好
- 数据量占用空间小,据说是influxdb的1/10

## 不足(特定场景下)

- 没有存储完整事件数据
- 数据类型只支持float64
- 数据存放到本机,没有集群方案
- 数据先存入内存再定期存入磁盘,查询性能好,但这样也带来了数据丢失的风险

## 结构图:

![](https://prometheus.io/assets/architecture.svg)


## prometheus是如何工作的:

prometheus有两个最重要的组件,支撑其最基本的运行,分别是Server, Job/Exporter

### job/exporter:

与zabbix等监控系统不同,prometheus没有agent的概念,而是切分为job和exporter

job用来监控自定义的服务,比如某些web业务,这类服务通常由某些个人或团体开发,不具有通用性.

exporter用来监控某些通用的服务,比如操作系统,数据库,中间件的性能指标.[目前社区有多个exporter供开发者使用开箱即用](https://prometheus.io/docs/instrumenting/exporters/)

不管还是job还是exporter,开发者都能自由编写他们.相关教程请参考: [如何编写prometheus监控组件](https://github.com/lwhhhh/monitorDoc/blob/master/如何编写prometheus监控组件.md)

### server:

job和exporter采集target的数据,就由server来收集并储存.prometheus采用的是pull模型,需要server主动去job或者exporter索取.关于pull和push模型的区别参考:[push与pull的区别](https://github.com/lwhhhh/monitorDoc/blob/master/push%E4%B8%8Epull%E7%9A%84%E5%8C%BA%E5%88%AB.md)

server拿到数据后会将数据保存到磁盘上,关于prometheus server是如何存储数据的内容请参考: [prometheus是如何做数据存储的](https://github.com/lwhhhh/monitorDoc/blob/master/prometheus是如何做数据存储的.md)


server不仅仅请求和保存数据,还有另外一个重要的功能是提供一种叫做PromQL的语法做查询,类似于SQL.接下来提到的数据展示功能都依赖与该功能.

### query:

开发者可以使用多个可视化工具展示prometheus收集到的监控数据,比如prometheus自带的web ui,以及grafana.同时prometheus也提供了API供开发者调用这些数据.

### altermanager:

待续