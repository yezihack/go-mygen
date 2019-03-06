# go 生成数据库表结构和markdown表结构命令工具
> 这个是一个命令行工具,
golang操作mysql的便捷工具, 快速生成表结束,快捷生成增删改查

## 安装
```
go install github.com/yezihack/gm2m
```

## 查看帮助
- 命令使用: help, h 查看更多命令如何使用
```
gm2m h
```

## 生成配置文件
- 命令使用: config, c 生成配置文件,用于设置数据库连接 和 其它参数
```
gm2m c
```

## 生成表结构体
- 命令使用: structure, s  生成golang数据结构数据
```
gm2m s
```

## 生成CRUD增删改查

```
gm2m curd
```

## 生成表markdown文档

- 命令使用: markdown, m   生成表结构的markdown文档
```
gm2m m
```



## 环境变量 GM2M_CONFIG 设置
> 设置mysql数据库的连接信息
- username 登陆用户名
- password 登陆密码
- localhost 连接地址
- 3306 端口
- dbname 数据库名称
- charset 编码

```
export GM2M_CONFIG="username:password@tcp(localhost:3306)/dbname?charset=utf8mb4"
# export GM2M_CONFIG="root:123456@tcp(localhost:3308)/kindled?charset=utf8mb4"
```

## default.ini的参数解析
```
;主机地址
host:
;端口
port:3308
;数据库名称
db_name:kindled
;连接名称
username:root
;连接密码
password:123456
;数据库编码
charset:utf8
[extra]
; 为空则默认所有表. 逗号隔开,如 表1,表2,表3
table_list=,data_comment,book
; structure 格式名称,默认为:json, 额外添加如: xml, gorm
struct_format=json,xml,
```
