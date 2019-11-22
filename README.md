# go-mygen
> 代码生成器,避免ORM带来的性能损耗.

## 版本
1. 2.0.0beta版本,又一次重构
1. 兼容linux,win,mac(还在测试)

## 安装
```
go get -u github.com/yezihack/go-mygen
```
## 下载
> release

## 功能介绍
- 1.生成数据库表的markdown文档
- 2.生成golang表对应的结构实体
- 3.自定义结构体解析实体,如json,gorm,xml等
- 4.生成golang操作mysql的增删改查语句
- 5.可以自定义生成目录
- 6.清屏,退出等功能


## 参数说明
```
-h value       数据库地址 (default: "localhost")
-P value       端口号 (default: 3306)
-u value       数据库用户名称 (default: "root")
-p value       数据库密码 (default: "root")
-c value       编码格式 (default: "utf8mb4")
-d value       数据库名称
--help         显示命令帮助
--version, -v  显示版本号
```

## 使用说明
- 首先必须先连接指定数据库
- 进行操作提示页面
- 输入界面上不同的命令进行操作即可

## 显示带密码操作
```
go-mygen -h localhost -P 3306 -u root -p 123456 -d default
```

## 保护密码操作
```
go-mygen -h localhost -P 3306 -u root -d dbname
```

## 其它
```
xx>  表示输入状态
xx:  信息提示
xx>> 错误输出
```
