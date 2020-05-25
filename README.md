# go 语言实现的简单静态文件服务

在写前端代码时， 预览编译出来的文件时候需要启动一个静态服务， 就用 go 实现了个简单的，
自己写的也容易修改

> 只支持 Linux， 我只有 Linux 环境， 其他系统没有测试过

## 快速使用

直接在 [github下载](https://github.com/broqiang/fservice/releases/download/v1.0.0/fservice)
编译好的可执行文件就可以直接使用。

## 编译，快速使用

> 需要本地有 go 环境

```bash
go build

./fservice
```

这样就可以启动一个当前目录为根目录的静态文件服务

## 环境变量配置

```bash
sudo mv fservice /usr/local/bin
```

这样就可以在任意地方启动文件服务了

例如：

```bash
cd /www/website
fservice
```

## 带参数启动服务

```bash
# -d 是指定静态文件所在的目录， 做为项目的根目录
# 支持绝对路径和相对路径
#
./fservice -d 指定目录 -p 指定端口
```

参数说明：

+ -d 是指定静态文件所在的目录， 做为项目的根目录

    支持绝对路径和相对路径, `/` 开头的，被定义为绝对路径， 没有 '/' 的做为相对路径

+ -p 端口号， 如 8080 8888 , 如果不指定，默认为 8888
