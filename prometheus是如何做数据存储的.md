# prometheus是如何做数据存储的


## 内存使用  

prometheus会将最经常使用的数据块保持在内存中,默认保存的数据块数量是1048576个.(如果使用type1方式编码,每个数据块大小是1024字节,那么该情况下缓存的数据量会是1GB)

可以配置
> storage.local.memory-chunks

修改缓存的数据块数量

另外要注意,内存实际占用空间不能简单做乘法得出,存储层也会占用一些空间.在极端情况下,如果所有的数据块都被同时使用,缓存的数据量可能会远远大于已配置的数值.在生产环境使用之前应当做好充分的测试.

测试时有两个参数可以提供帮助:
> prometheus_local_storage_memory_chunks 

> process_resident_memory_bytes

根据经验, 至少应该预留3倍缓存的空间.


还需要注意一点,PromQL的查询实现依赖于LevelDB提供的索引支持,如果大量使用PromQL做查询,那么可能需要调整索引缓存的大小.涉及到的参数有:

> -storage.local.index-cache-size.label-name-to-label-values: 用于正则匹配.

> -storage.local.index-cache-size.label-pair-to-fingerprints: Increase the size if a large number of time series share the same label pair or name.



> -storage.local.index-cache-size.fingerprint-to-metric  
> -storage.local.index-cache-size.fingerprint-to-timerange :
> Increase the size if you have a large number of archived time series, i.e. series that have not received samples in a while but are still not old enough to be purged completely.

 如果需要查询的时间序列超过10万条,那么单次查询所需的内存空间可能就要超过100Mib.所以如果内存空间足够的话,应当尝试分配更多的内存空间给予LevelDB

## 存储结构

prometheus会对样本数据进行分块处理,每块的大小为1024字节.

### 编码方式

prometheus目前提供了三种算法(主要是为了压缩数据)用于块的编码,可以通过-storage.local.chunk-encoding-version进行配置.参数的有效值为0,1,2.

chunk-encoding为0时,采用的是一种叫做delta encoding的算法.早期的prometheus存储层用的就是该实现.

chunk-encoding为1时,是一种改进型的double-delta encoding算法,目前的额prometheus默认使用该编码方式.

这两种编码方式对每个块使用固定的字节长度,这样有利于随机读取.

chunk-encoding为2时,使用的则是可变长的编码方式.这种编码比起上面两种方式,特点在于牺牲压缩速度换取了压缩率.

下面展示了压缩同样大小的数据对比(文档说样本很大,但没说具体多少):

|编码类型|压缩后样本大小|所用时间|
|-------|----------|------|
|1      | 3.3bytes | 2.9s|
|2      |1.3bytes  |4.9s|


## 测试:

官方给出在生产环境中,每个样本加上索引信息后的大小一般为3-4bytes,我们可以做下测试看看实际的样本有多大.

因为数据文件是经过处理后写入磁盘的,所以没办法查看单个样本的大小,只能采集一段时间的数据后计算.

测试的监控目标的有两个,一个是prometheus本身的信息,一个是node-exporter输出的硬件数据,我们的分别访问host:port/metrics获取采集到的数据内容.在这个例子中,每进行一次采集,prometheus server就会取回145756 bytes的数据.(即访问两个/metrics接口返回的数据相加)


第一次:
开始时间: 14:22
data文件夹为空(存放prometheus server取回的数据,下同)
抓取频率: 5s
结束时间: 14:42
样本大小: 1597440bytes
原始大小: 20 min * 60s / 5 * 145756bytes = 34981440bytes
压缩率: (34981440 - 1597440) / 34981440 ~= 95%


第二次:
开始时间:14:57
data不动
结束时间:15:07
样本大小:2600960bytes
原始大小: 10 min * 60s / 5 * 145756bytes = 17490720bytes
压缩率: (17490720bytes - (2600960bytes - 1597440bytes)) / 17490720bytes ~= 94

<!--第三次:
开始时间:15:23
结束时间:17:03
间隔:   100min
结束大小:7720960bytes
原始大小: 100min * 60s / 5 * 145756bytes = 174907200bytes
压缩率: (174907200 - (7720960 - 2600960)) / 174907200 ~= 97%-->


第三次:
开始时间:08:40
结束时间:11:15
开始大小:10640
结束大小:14784
间隔: 155min
原始大小: 155 * 60 / 5 * 145756 = 271106160
压缩率: (271106160 -  (14784 - 10640) * 1024)  / 271106160 ~= 98%

第四次:
开始时间: 12:13
结束时间: 12:23
开始大小:14784
间隔: 10min
结束大小: 16404
原始大小:10 min * 60s / 5 * 145756bytes = 17490720bytes
压缩率: (17490720 - (16404-14784)*1024)/17490720 = ~90%

第五次:
开始时间:12:28
结束时间:12:48
开始大小:16404
结束大小:19804
间隔:20min
原始大小:34981440
压缩率:(34981440 - (19804-16404)*1024)/34981440

<!--
第六次:
开始时间:12:53
结束时间:13:03
开始大小:19804
结束大小:20608
间隔: 10min
原始大小:17490720
压缩率:(17490720 - (20608-19804)*1024)/17490720-->


|      |用时   |抓取频率  |数据变化量(bytes)|原始大小(bytes)|压缩率|
|------|------|---------|---------------|----------------|-----|
|第一次 |10min |5s       |+1003520       |17490720        | 94% |
|第二次 |20min |5s       |+1597440       |34981440        | 95% |
|第三次 |155min|5s       |+4243456       |271106160       | 98% |
|第四次 |10min |1s       |+1658880       |17490720        | 90% |
|第五次 |20min |1s       |+3481600       |34981440        | 90% |

按照抓取频率5s,压缩率90%进行粗略估算.

假设检测的数据为系统的硬件指标,即node-exporter的输出(145756个字节),且集群中有10台机器,那么24个小时的数据量将不超过200m.假设监控数据保留1个月,那么大概需要6-7G左右的空间


