# 使用prometheus监控docker容器

本文介绍如何使用cAdvisor监控docker容器的运行状态

*cAdvisor之前prometheus有个container-exporter可以做容器的监控,后来项目停止维护就抛弃了.*

首先我们需要运行cAdvisor,可以让cAdvisor运行在docker内,也可以在host运行.本文采用docker的方式运行.

> docker run \
  --volume=/:/rootfs:ro \
  --volume=/var/run:/var/run:rw \
  --volume=/sys:/sys:ro \
  --volume=/var/lib/docker/:/var/lib/docker:ro \
  --publish=8080:8080 \
  --detach=true \
  --name=cadvisor \
  google/cadvisor:latest

接下来配置prometheus.yml:

    global:
    scrape_interval:     5s
    evaluation_interval: 5s

    - job_name: docker
        static_configs:
        - targets: ['192.168.1.185:8080']
            labels:
            instance: db1

*192.168.1.185是宿主机的地址,因为我们的prometheus也将运行在docker中*

启动prometheus容器:

> docker run -it -d -p 9090:9090 -v ~/Go/src/github.com/prometheus/prometheus/:/etc/prometheus/ prom/prometheus 

这个时间对docker的监控就已经开始了.要注意的是,cAdvisor的监控对象是宿主机上的所有容器,包括运行cAdvisor本身那个容器!

这时候我们就可以使用grafana或者其他工具查看监控数据.

cAdvisor自带一个可视化工具,访问host:port/containers即可看到.