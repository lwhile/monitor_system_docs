# 使用InfluxDB做监控平台

InfluxDB是一款使用Golang编写的开源时序数据库,其最大特点有:

- 技术实现上充分利用了Go语言的特性，无需任何外部依赖即可独立部署
- 提供了一个类似于SQL的查询语言并且一系列内置函数方便用户进行数据查询
- InfluxDB支持基于HTTP的数据插入与查询。同时也接受直接基于TCP或UDP协议的连接

同时0.8.4版本之后官方集成了collectd插件,可以很方便那它做监控系统使用.


## 安装:

首先在官网下载执行文件,官网提供了三种安装方式,一种是使用docker,一种是针对特定操作系统的安装文件,一种是编译后的二进制文件.我在这里选择最后一种.

在Linux下:

> wget https://dl.influxdata.com/influxdb/releases/influxdb-1.2.2_linux_i386.tar.gz

> tar xvfz influxdb-1.2.2_linux_i386.tar.gz

操作过后会生成三个文件夹:

- etc: 存放配置文件
- lib: 和执行相关的文件
- share: 存放共享数据

接下来我们需要的两个文件分配是etc/influxdb/influxdb.conf(注意不是/etc)和usr/bin下的可执行文件(也注意不是/usr)


这时候我们先把influxdb放一边,先安装collectd,做数据的采集器.

> sudo apt-get install collectd

安装后配置/etc/collectd/collectd.conf

启用LoadPlugin network(把注释去掉即可),
并配置influxdb的所在ip和接受端口

    <Plugin network>
            Server "127.0.0.1" "25826"
    </Plugin>

重新启动collectd:

> sudo /etc/init.d/collectd restart

此时再回到InfluxDB,将collectd插件启动.


    [[collectd]]
        enabled = true
        bind-address = ":25826" # the bind address
        database = "collectd" # Name of the database that will be written to
        retention-policy = ""
        batch-size = 5000 # will flush if this many points get buffered
        batch-pending = 10 # number of batches that may be pending in memory
        batch-timeout = "10s"
        read-buffer = 0 # UDP read buffer size, 0 means to use OS default
        typesdb = "/usr/share/collectd/types.db"
        security-level = "none" # "none", "sign", or "encrypt"
        auth-file = "/etc/collectd/auth_file"


**typesdb文件可在[这里](https://github.com/collectd/collectd/blob/master/src/types.db)下载**

配置完后启动influxdb,我们需要指定使用的配置文件

> sudo ./usr/bin/influxd -config=etc/influxdb/influxdb.conf


打开http://host:8083可以进入influxdb的web管理端,创建一个名叫collectd的数据库


## 查询:

influxdb提供两种查询途径,一种是基于http的API请求,一种是使用类似SQL的查询语法.

通过http的方式获取CPU的使用情况:

> curl -G 'http://localhost:8086/query?pretty=true' --data-urlencode "db=collectd" --data-urlencode "q=SELECT \"value\" FROM \"cpu_load_short\" 

返回:

    {
        "results": [
            {
                "statement_id": 0,
                "series": [
                    {
                        "name": "cpu_load_short",
                        "columns": [
                            "time",
                            "value"
                        ],
                        "values": [
                            [
                                "2015-01-29T21:55:43.702900257Z",
                                2
                            ],
                            [
                                "2015-01-29T21:55:43.702900257Z",
                                0.55
                            ],
                            [
                                "2015-06-11T20:46:02Z",
                                0.64
                            ]
                        ]
                    }
                ]
            }
        ]
    }

更多的API使用可以参考官方文档:
[API Reference](https://docs.influxdata.com/influxdb/v1.2/tools/api/)

## 二次开发

接下来介绍如何基于colletcd + InfluxBD 做二次开发.

collectd已经有很多插件可以使用,这些插件可以在官方网站或者社区里面找到.
[官方提供插件列表](https://collectd.org/wiki/index.php/Table_of_Plugins)

但是当一些需求没有现成的插件使用时,我们还是得自己编写.
接下来我们要关注如何把自己的数据告诉collectd,以及collectd是如何将数据告诉InfluxDB的.

collectd提供了一个叫做Exec的插件,用来执行特定脚本以及二进制文件,然后读取他们输出到STDOUT的数据.读取到的数据会被collectd的daemon进程当做监控数据处理.(这和Zabbix很像)


以Go语言为例:

    func collectd() {
        unixTs := time.Now().Unix()
        f := bufio.NewWriter(os.Stdout)
        hostLabel := os.Hostname()
        defer f.Flush()
        b := "PUTVAL " + hostlabel + "/" + "bar" + "/" + "gauge-name " + strconv.Itoa(unixTs) + ":" + "value\n"
        f.Write([]byte(b))
    }


代码中将一定格式的数据输出到了stdout上,collectd会自动从stdout上读取数据.
数据的格式要求为:

> 'PUTVAL'  {hostname  /application   /type-stat_name}   {epoch_time_stamp}:{value}

如:

> PUTVAL localhost/test/gauge-used-memory 1489799746:4834788