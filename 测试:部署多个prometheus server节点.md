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

两个prometheus server得到的数据如下:

![](https://raw.githubusercontent.com/lwhhhh/monitorDoc/master/asset/images/dockerA1.png)

![](https://raw.githubusercontent.com/lwhhhh/monitorDoc/master/asset/images/dockerB1.png)


        