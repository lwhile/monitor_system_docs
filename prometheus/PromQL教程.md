# PromQL教程

## 四种数据类型

1. Instant vector : 每个指标返回一个值,且指标集合里面的时间戳都相同.类似于即时数据.

2. Range vector: 每个指标包含多个时序数据.

3. Scalar: 只有一个浮点值

4. String: 未开放使用.

## 二元操作

1. \+
2. \- 
3. \*
4. \/ 
5. \%
6. \^

## 聚合操作

## 函数

1. rate(v range-vector)
    
     根据每个点计算每秒的平均变化率.

2. irate(v range-vector)

     根据最后两个数据点计算变化率.irate()适合于变化较快的数据,rate()适合变化较慢的数据.

    

二元操作符使用范围:

1. scalar/scalar
2. vector/scalar
3. vector/vector 


## 常用例子集合:

查询CPU使用率:

    100 - (avg by (job) (irate(node_cpu{mode="idle"}[5m])) * 100)

查询网卡即时速率

    irate(node_network_receive_bytes[2m])