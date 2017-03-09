# 整合prometheus与consul

在一个集群中往往需要部署多个监控server以提高服务的可用性,本文介绍如何使用整合consul与prometheus,尽量做到服务的高可用.


关于consul的更多介绍可以参考:[consul.io](https://www.consul.io/)

本文的实验用到的结构为:

![](https://raw.githubusercontent.com/lwhhhh/monitorDoc/master/asset/images/prometheus-consul.png)

使用4个docker运行做consul的agent,其中dockerA和dockerD做server,dockerB和dockerC做client.同时4个容器都运行prometheus server服务,向node-exporter获取监控数据

四个容器得到的数据再通过grafana做可视化.