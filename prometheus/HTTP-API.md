# HTTP API

目前API的稳定接口为/api/v1,接下来新接口的加入会保持对/api/v1的兼容性

API返回数据为JSON格式,每一个成功的HTTP请求返回的状态码为2xx

    {
    "status": "success" | "error",
    "data": <data>,

    // Only set if status is "error". The data field may still hold
    // additional data.
    "errorType": "<string>",
    "error": "<string>"
    }

## 即时查询

即时查询返回某一时间戳下的数据

> GET /api/v1/query

参数:

- query=\<string\>:string内容为prometheus的查询表达式
- time=\<unix_timestamp\>:指定查询的范围,可选.默认为服务器当前时间

在返回的数据中,data字段的格式如下:

    {
    "resultType": "matrix" | "vector" | "scalar" | "string",
    "result": <value>
    }

\<value\>的数据类型与resultType字段有关,[查看查询结果的四种resultType](#1)


举例

    $ curl 'http://localhost:9090/api/v1/query?query=up&time=2015-07-01T20:10:51.781Z'
    {
    "status" : "success",
    "data" : {
        "resultType" : "vector",
        "result" : [
            {
                "metric" : {
                "__name__" : "up",
                "job" : "prometheus",
                "instance" : "localhost:9090"
                },
                "value": [ 1435781451.781, "1" ]
            },
            {
                "metric" : {
                "__name__" : "up",
                "job" : "node",
                "instance" : "localhost:9100"
                },
                "value" : [ 1435781451.781, "0" ]
            }
        ]
    }
    }

## 范围查询

范围查询返回某一时间段内的数据

> GET /api/v1/query_range

参数

- query=\<string\>:prometheus的查询表达式
- start=\<unix_timestamp\>:开始时间的时间戳
- end=\<unix_timestamp\>:结束时间的时间戳
- step=<duration>:数据精度

返回的数据格式为:

    {
    "resultType": "matrix",
    "result": <value>
    }

# todo!

举例:

    $ curl 'http://localhost:9090/api/v1/query_range?query=up&start=2015-07-01T20:10:30.781Z&end=2015-07-01T20:11:00.781Z&step=15s'
    {
    "status" : "success",
    "data" : {
        "resultType" : "matrix",
        "result" : [
            {
                "metric" : {
                "__name__" : "up",
                "job" : "prometheus",
                "instance" : "localhost:9090"
                },
                "values" : [
                [ 1435781430.781, "1" ],
                [ 1435781445.781, "1" ],
                [ 1435781460.781, "1" ]
                ]
            },
            {
                "metric" : {
                "__name__" : "up",
                "job" : "node",
                "instance" : "localhost:9091"
                },
                "values" : [
                [ 1435781430.781, "0" ],
                [ 1435781445.781, "0" ],
                [ 1435781460.781, "1" ]
                ]
            }
        ]
    }
    }


## 查询元信息:


>  GET /api/v1/series

参数:

- match[]=\<series_selector\>:返回符合label集合的序列
- start=\<unix_timestamp\>.
- end=\<unix_timestamp\>

举例:
查询所有的符合up或者process_start_time_seconds{job="prometheus"}的集合

**注意:多个match[]参数的之间的关系为或**

    $ curl -g 'http://localhost:9090/api/v1/series?match[]=up&match[]=process_start_time_seconds{job="prometheus"}'
    {
    "status" : "success",
    "data" : [
        {
            "__name__" : "up",
            "job" : "prometheus",
            "instance" : "localhost:9090"
        },
        {
            "__name__" : "up",
            "job" : "node",
            "instance" : "localhost:9091"
        },
        {
            "__name__" : "process_start_time_seconds",
            "job" : "prometheus",
            "instance" : "localhost:9090"
        }
    ]
    }

## 查询标签值

> GET /api/v1/label/<label_name>/values

    $ curl http://localhost:9090/api/v1/label/job/values
    {
    "status" : "success",
    "data" : [
        "node",
        "prometheus"
    ]
    }

## 删除数据

> DELETE /api/v1/series

参数:

- match[]=/<series_selector/>

举例,删除符合up或者process_start_time_seconds{job="prometheus"}的数据:

    $ curl -XDELETE -g 'http://localhost:9090/api/v1/series?match[]=up&match[]=process_start_time_seconds{job="prometheus"}'
    {
    "status" : "success",
    "data" : {
        "numDeleted" : 3
    }

<h2 id="1">查询结果的四种resultType</h2>

- matrix

查询某个时间段内的数据返回的vector称为matrix类型,格式如下:

    [
        {
            "metric": { "<label_name>": "<label_value>", ... },
            "values": [ [ <unix_time>, "<sample_value>" ], ... ]
        },
        ...
    ]


- vector

查询即使数据返回的结果称为vector,格式如下:

    [
        {
            "metric": { "<label_name>": "<label_value>", ... },
            "value": [ <unix_time>, "<sample_value>" ]
        },
        ...
    ]

*matrix和vector之间的区别在于value字段,一个是单数,一个是复数*


- Scalars

- String