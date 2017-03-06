# prometheus

项目地址: [prometheus](https://github.com/prometheus/prometheus)

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

## 结构图:

![](https://prometheus.io/assets/architecture.svg)

## 集群部署与配置:

集群内使用一个结点作为prometheus server, 在需要被监控的机器或者服务上, 需要安装prometheus的exporter或者