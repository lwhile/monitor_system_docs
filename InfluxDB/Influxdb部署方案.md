# Influxdb部署方案

监控平台部署两套Influxdb环境

### 机器性能监控方面:

数据存放在表 xxx-monitor-device 里

该表设置两个Retention Policy:

- autogen: 该RP数据的保存时间为1年, 存放采集频率为1min的数据

- save_1h: 该RP数据的保存时间为一个小时, 存放采集频率为10s的数据

采集工具的配置文件里, 针对cpu, disk, memory, net 这四个采集项, 都配置两个不同的inputs.plugin配置, 以及两个不同的outputs.plugin配置

    # inputs.plugin
    - 采集频率为1min: 全局配置采集频率为1min, 无需重新配置interval字段.
    - 采集频率为10s: retention_policy 配置为save_1h, interval配置为10s, name_suffix为"_10s"

    # outputs.plugin
    将其中 retention_policy 配置为"save_1", namepass 配置为 ["*_10s"]


### 其他监控数据

数据存放在表 xxx-monitor-plugin 里