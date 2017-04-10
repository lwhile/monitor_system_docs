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
