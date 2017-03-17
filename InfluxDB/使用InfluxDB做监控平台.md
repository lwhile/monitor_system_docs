# 使用InfluxDB做监控平台

InfluxDB是一款使用Golang编写的开源时序数据库,其最大特点有:

- 技术实现上充分利用了Go语言的特性，无需任何外部依赖即可独立部署
- 提供了一个类似于SQL的查询语言并且一系列内置函数方便用户进行数据查询
- InfluxDB支持基于HTTP的数据插入与查询。同时也接受直接基于TCP或UDP协议的连接

同时0.8.4版本之后官方集成了collectd插件,可以很方便那它做监控系统使用.下面介绍了如何搭建一个最简单的InfluxDB + collectd的试验环境.


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


打开htpt://host:8083可以进入influxdb的web管理端

   