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

## 2.下载lua-nginx-module

## 3.