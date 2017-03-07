# 测试部署多个prometheus server节点

## 测试内容

使用docker运行两个prometheus容器,监控node_exporter输出的系统指标

## 关注的结果

- 两个docker获取到的数据是否一致
- 关闭其中一个docker容器查看另外一个容器是否能继续正常运行
- 查看prometheus server存储的数据源是否有多个


## 结构图

![](https://raw.githubusercontent.com/lwhhhh/monitorDoc/master/asset/images/multiple%20prometheus%20server.png)