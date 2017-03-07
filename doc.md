# 开源监控方案对比

|  | 集群部署和配置 | 数据存储 | API | 前端展现 | 监控插件二次开发 | 安装时依赖项|推送模型|
|--|--|--|--|--|--|--|--|
| prometheus | |使用levelDB做索引,自带存储层 | √| 自带组件,支持第三方如grafana| √| 无|默认为poll,提供push模式|
| Open-Falcon| | 自带的Graph组件或使用opentsdb| 支持,文档详细 | | | redis,mysql| |
| zabbix | | MySQL| | | |  | |