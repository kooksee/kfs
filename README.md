# kfs
类似于ipfs的去中心化的文件存储系统,带有缓存优化,外网穿透

## kfs config
配置文件

## kfs daemon
后台运行程序

## kfs add file
添加一个文件到kfs中

## kfs ls -n 10
列出所有的文件名称以及hash

## kfs metadata file_hash
查看该文件的本地的所有的metadata

## kfs get file_hash
得到该文件
查询方式有dht和cache两种方式
数据接收为被动

## kfs pin file_hash
把该文件固定到本地中
该命令会在本文件下载完毕之后然后把本节点地址以及文件地址发布到去中心化的cache中
