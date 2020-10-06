OpenAL bindings for Go, which works well with windows platform
====================================

本项目基于[https://github.com/phf/go-openal](https://github.com/phf/go-openal) 进行修改

本仓库包含Openal-soft 1.2的静态链接库和头文件，用于解决Golang在Windows平台编译时出现的找不到``<AL/al.h>``的问题

## 安装
```
go get -u github.com/dextercai/go-openal/
```
