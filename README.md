# go 生成数据库表结构和markdown表结构命令工具
>

## 安装
```
go get -u github.com/yezihack/gm2m
```

## 生成表结构体
- 命令使用: structure, s  生成golang数据结构数据
```
gm2m s
```

## 生成表markdown文档

- 命令使用: markdown, m   生成表结构的markdown文档
```
gm2m m
```

## 生成CRUD简单查询


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

## 环境变量 GM2M_TABLES 设置
> 设置表的环境变量,用于指定生成哪些表的数据,如果不变量,或者为空,则读取所有的表

```
export GM2M_TABLES="table1, table2, table3, ..."
```