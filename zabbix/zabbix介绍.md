# zabbix

zabbix是目前各大互联网公司使用最广泛的开源监控之一,其历史最早可追溯到1998年,在业内拥有各种成熟的解决方案.


zabbix属于CS架构,Server端基于C语言编写,相比其他语言具有一定的性能优势(在数据量不大的情况下!).Web管理端则使用了PHP.
而其client端有各种流行语言的库实现,方便使用其API


是数据的存储方面,zabbix使用了关系性数据库,包括SQLite,MySQL,PostgreSQL,Oracle,DB2



## 架构图

![](http://img1.51cto.com/attachment/201208/130335905.png)


接下来我们进一步了解zabbix,为技术选型提供更多帮助.

### 安装:

zabbix的安装比较繁琐,但也不算困难(主要是因为网上提供的资料足够多)

我们需要安装关系型一种关系型数据库,目前提供的选择有MySQL,SQLite, PostgreSQL,Oracle,DB2

接下来需要安装PHP的运行环境,Web服务器可是使用Apache或者Nginx都可以.

最后一步是安装zabbix服务.

完整的安装教程可以参考:

[zabbix安装指南](http://www.jianshu.com/p/4d98ff87db5f)


### 数据的采集

zabbix采集的数据的方式为proxy或agentpush给server.push和pull的区别可以参考[push与pull的区别](https://github.com/lwhhhh/monitorDoc/blob/master/push%E4%B8%8Epull%E7%9A%84%E5%8C%BA%E5%88%AB.md)