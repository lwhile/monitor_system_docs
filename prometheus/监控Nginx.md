# 监控Nginx

prometheus监控Nginx需要对Nginx做一些配置,才能顺利拿到Nginx的数据

## 环境依赖:
- [Nginx源码](http://nginx.org/en/download.html)
- [lua-nginx-module](https://github.com/openresty/lua-nginx-module)
- [Lua-JIT](http://luajit.org/download.html)
- [nginx-lua-prometheus](https://github.com/knyar/nginx-lua-prometheus)


## 1.编译安装Lua-JIT环境

> make PREFIX=/usr/local/luajit

> make install PREFIX=/usr/local/luajit

设置环境变量:

> export LUAJIT_LIB=/usr/local/luajit/lib

> export LUAJIT_INC=/usr/local/luajit/include/luajit-2.1

## 2. 下载lua-nginx-module


## 3. 编译Nginx

> ./configure --prefix=/usr/local/nginx (这里填写你需要安装的模块) --add-module=/path/to/lua-nginx-module

> make

> make install


## 4.验证lua-nginx-module是否安装成功

修改Nginx配置:

    location /hello_lua { 
      default_type 'text/plain'; 
      content_by_lua 'ngx.say("hello, lua")'; 
    }

启动Nginx服务:

> /usr/local/nginx/sbin/nginx

访问localhost/hello_lua查看浏览器能否显示数据


**参见错误: error while loading shared libraries: libluajit-5.1.so.2: cannot open shared**

**解决方法: 重新安装libluajit-5.1这个库,ubuntu下使用apt即可安装,centos需要下载独立的rpm包**

## 5.配置nginx-lua-prometheus模块

将以下代码放进Nginx配置文件的Http配置块里:

    lua_shared_dict prometheus_metrics 10M;
    lua_package_path "/path/to/nginx-lua-prometheus/?.lua";
    init_by_lua '
    prometheus = require("prometheus").init("prometheus_metrics")
    metric_requests = prometheus:counter(
        "nginx_http_requests_total", "Number of HTTP requests", {"host", "status"})
    metric_latency = prometheus:histogram(
        "nginx_http_request_duration_seconds", "HTTP request latency", {"host"})
    metric_connections = prometheus:gauge(
        "nginx_http_connections", "Number of HTTP connections", {"state"})
    ';
    log_by_lua '
    local host = ngx.var.host:gsub("^www.", "")
    metric_requests:inc(1, {host, ngx.var.status})
    metric_latency:observe(ngx.now() - ngx.req.start_time(), {host})
    ';

**注意第二行修改你本机的nginx-lua-prometheus模块位置**

使用命令检查配置是否正确:

> /usr/local/nginx/sbin/nginx -t

配置Nginx的监听端口:

    server {
        listen 9145;
        location /metrics {
            content_by_lua '
            metric_connections:set(ngx.var.connections_reading, {"reading"})
            metric_connections:set(ngx.var.connections_waiting, {"waiting"})
            metric_connections:set(ngx.var.connections_writing, {"writing"})
            prometheus:collect()
            ';
        }
    }

**注意在生产环境中要设置只允许prometheus server能够访问这个端口**

使用命令检查配置是否正确:

> /usr/local/nginx/sbin/nginx -t

若准确无误,更新Nginx服务

> /usr/local/nginx -s reload


至此Nginx配置完成

