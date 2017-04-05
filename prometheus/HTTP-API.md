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

\<value\>的数据类型与resultType字段有关,[Expression query result formats]()


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

