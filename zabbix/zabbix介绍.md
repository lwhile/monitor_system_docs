# zabbix

zabbix是目前各大互联网公司使用最广泛的开源监控之一,其历史最早可追溯到1998年,在业内拥有各种成熟的解决方案.


zabbix属于CS架构,Server端基于C语言编写,相比其他语言具有一定的性能优势(在数据量不大的情况下!).Web管理端则使用了PHP.
而其client端有各种流行语言的库实现,方便使用其API


是数据的存储方面,zabbix使用了关系性数据库,包括SQLite,MySQL,PostgreSQL,Oracle,DB2



## 架构图

![](http://img1.51cto.com/attachment/201208/130335905.png)


接下来我们进一步了解zabbix,为技术选型提供更多帮助.

## 安装:

zabbix的安装比较繁琐,但也不算困难(主要是因为网上提供的资料足够多)

我们需要一种关系型关系型数据库,目前提供的选择有MySQL,SQLite, PostgreSQL,Oracle,DB2

接下来需要安装PHP的运行环境,Web服务器可是使用Apache或者Nginx都可以.

最后一步是安装zabbix服务.

完整的安装教程可以参考:

[zabbix安装指南](http://www.jianshu.com/p/4d98ff87db5f)


## 数据的采集

在目标机器上采集数据(metrics)需要安装zabbix agent.

agent采集到数据后,会立即push给proxy或者server 

zabbix对分布式的数据采集非常好,支持两种分布式架构,一种是Proxy,一种是Node.Proxy作为zabbix server的代理去监控服务器,并发数据汇聚到Zabbix server.而Node本身就是一个完整的Zabbix server,
使用Node可以将多个Zabbix server组成一个具有基层关系的分布式架构.

两者的区别如下:

|              |proxy|Node|
|--------------|-----|----|
|轻量级         |√    |×   |
|GUI前端        |×    |√  |
|是否可以独立运行 |√    |×   |
|容易运维        |√   | ×   |
|本地Admin管理   |×   |√   |
|中心化配置      |√   |×    |
|产生通知       |×   |√    |


## 监控目标的发现

zabbix支持自动添加监控目标,只需要提供IP范围,就能在机器新加入后使用已经设置好的模式进行监控.

但是该机制不支持使用静态文件进行配置,只能在Web面板进行设置.

## 二次开发

二次开发的内容我们分为两个方面,一个是Zabbix API,一个是自定义监控目标.

- Zabbix API:

    Zabbix提供的API非常强大,提供了几乎Zabbix的所有功能,比如更新Item,添加Host监控等.

    Zabbix API以HTTP服务的形式对外开发,并要求以JSON-RPC的方式发起请求.

    举个例子,发出一个登录的请求:

    > curl -i -X POST -H 'Content-Type: application/json' -d '{"jsonrpc":"2.0","method":"user.login","params":{"user":"Admin","password":"zabbix"}, "auth":null,"id":0}' http://127.0.0.1/zabbix/api_jsonrpc.php

    能得到返回值:

    > {"jsonrpc":"2.0","result":"3144f3ier0394efasdg342as","id":0}

    官方也提供了很详细的API文档:

    [Zabbix API Documentation](https://www.zabbix.com/documentation/2.2/manual/api)

- 自定义监控目标

    Zabbix使用编写脚本输出的方式自定义监控目标.
    zabbix agent安装后会生成一份zabbix_agentd.conf文件,这个文件记录的监控项的配置.
    添加监控项只需要在文件中添加如下形式的命令: 
    > UserParameter = \<key> ,\<shell command>


    比如想监控http的连接数,我们需要在zabbix_agentd.conf中写入:
    > UserParameter=httpcount, netstat -an | grep 192.168.1.126:80 wc -l
    
    shell command也可以换成相同输出的python脚本位置.

    之后只需要在web监控端添加httpcount这个item即可.


## 获取监控数据

监控数据的获取可以web前台或者API获取,因为这方面有需求所以特地拿API这一方式出来单独说明.

以获取CPU负载的例子来说明,依旧是向zabbix server发送JSON-RPC请求:

    {
        "jsonrpc": "2.0",
        "method": "graph.get",
        "params": {
            "output": "extend",
            "sortfield": "name",
            "hostid": "10084"
        },
        "auth": "50ba559f6d083aa6454b8b3c4c203baa",
        "id": 1
    }

得到的响应是:

    {
        "jsonrpc": "2.0",
        "result": [
        
            {
                "graphid": "481",
                "name": "CPU load",
                "width": "900",
                "height": "200",
                "yaxismin": "0.0000",
                "yaxismax": "100.0000",
                "templateid": "0",
                "show_work_period": "1",
                "show_triggers": "1",
                "graphtype": "0",
                "show_legend": "1",
                "show_3d": "0",
                "percent_left": "0.0000",
                "percent_right": "0.0000",
                "ymin_type": "1",
                "ymax_type": "0",
                "ymin_itemid": "0",
                "ymax_itemid": "0",
                "flags": "0"
            }
        ],
        "id": 1
    }

前端再对JSON数据进行处理.zabbix的API只支持JSON-RPC的形式,所以前端调用的时候也需要了解JSON-RPC的使用方法.也可以自己部署一个代理服务将JSON-RPC转换为Http的形式


## zabbix的数据存储

或许是开发时间早, 与这几年出现的监控系统不同,zabbix的数据是存储在关系性数据库里面的.
这意味着数据量或者并发量大时,比起时序数据库,性能将会成为一个问题,需要进行优化.

使用关系型数据库的最大好处就是继承了数据库的各种高级功能如分区,读写分离,主从备份等.

## 告警系统

zabbix的报警功能很强大...同时也很复杂.
zabbix的报警单位叫做Trigger(触发器),分为两个状态(2.X后的版本),"OK","Problem".

Trigger使用Expression配置Trigger的判断逻辑,通过Expression可以很灵活得实现需求.

Expression的基本形式如下:

> {\<server>:\<key>.\<function>(\<parameter>)}\<operation>\<constant>


> {zabbix1:system.cpu.load[all,avg1].last(0)}>5 : zabbix1这台机器CPU负载值是否大于5


> 	{zabbix2:system.cpu.load[all,avg1].last(0)}>5 |           {zabbix2:system.cpu.load[all,avg1].min(10m)}>2 : zabbix2这台机器当前cpu负载是否大于5或者最近10分内的cpu是否负载大于2




