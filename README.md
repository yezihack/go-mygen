[![Build Status](https://travis-ci.org/yezihack/go-mygen.svg?branch=master)](https://travis-ci.org/yezihack/go-mygen)
[![](https://img.shields.io/github/stars/yezihack/go-mygen)](https://github.com/yezihack/go-mygen/stargazers)
[![](https://img.shields.io/github/issues/yezihack/go-mygen)](https://github.com/yezihack/go-mygen/issues)
[![](https://img.shields.io/github/forks/yezihack/go-mygen)](https://github.com/yezihack/go-mygen/network/members)
[![](https://img.shields.io/github/license/yezihack/go-mygen)](https://github.com/yezihack/go-mygen/blob/3.1.0beta/LICENSE)

[English](README.md)

![](assets/img/golang.png)

# go-mygen
> Code generator tools with MYSQL,avoid ORM performance penalty, 
> Go-mygen is a tool that automatically generates MySQL table entities,CRUD and Markdown documents.

> go-mygen是一款自动生成MySQL表实体,CRUD,Markdown文档工具

## Install
> \>= go1.13.0
```
go get -u -x github.com/yezihack/go-mygen
```

## Using v3.2 releases
> [release](https://github.com/yezihack/go-mygen/releases/tag/v3.2.0)

## Version information
[go-mygen-v3.3.8](CHANGELOG.md)

## QuickStart
1. See Help `go-mygen help`
1. See Version `go-mygen v`
1. Connection `go-mygen -h localhost -P 3306 -u root -p 123456 -d default `

## Features
- Generate the markdown document for the database table
- Generate the structure entities for the golang table
- Custom structure parsing entities, such as json,gorm, XML, etc
- Generate golang operation mysql add, delete, modify and select code statements
- Config, Entity and CURD directories are stored separate
- Custom build directory
- Clear screen, exit and other functions
- Use an asterisk instead of the password you entered

## Parameters info

> Go-mygen GLOBAL OPTIONS

```
GLOBAL OPTIONS:
   -h value  Database address (default: "127.0.0.1")
   -P value  port number (default: 3306)
   -u value  database username (default: "root")
   -p value  database password
   -c value  database format (default: "utf8mb4")
   -d value  database name
   --debug   debug (default: false) 
```

> Operation Command (input number)

```
NO:0 Set build directory
NO:1 Generate the table markdown document
NO:2 Generate table structure entities
NO:3 Generate CURD insert, delete, update and select
NO:4 Sets the struct mapping name
NO:5 Find or set the table name
NO:7, c, clear Clear the screen
NO:8, h, help Show help list
NO:9, q, quit Quit
```

> Go-mygen 全局参数说明
```
GLOBAL OPTIONS:
   -h value  Database address (default: "127.0.0.1") 数据库地址,必填
   -P value  port number (default: 3306) 数据库端口号,必填
   -u value  database username (default: "root") 数据库名称,必填
   -p value  database password 数据库登陆密码,必填
   -c value  database format (default: "utf8mb4") 数据库编码
   -d value  database name 数据库名称,必填
   --debug   debug (default: false) 开启调试
```
> 操作命令,只需要输入NO号,即可操作相关命令

```
NO:0 Set build directory 设置存储目录
NO:1 Generate the table markdown document 生成Markdonw文档
NO:2 Generate table structure entities 生成Golang代码的实体(即表的实体)
NO:3 Generate CURD insert, delete, update and select 生成Golang的CRUD代码
NO:4 Sets the struct mapping name 设置实体解析参数,多个,以逗号分隔. 如`json:"name"`
NO:5 Find or set the table name 选择需要操作的数据库表, 默认全部
NO:7, c, clear Clear the screen 清屏
NO:8, h, help Show help list 显示帮助
NO:9, q, quit Quit 安全退出
```

## Package 
> Thanks for these packages
1. `go get -u github.com/go-bindata/go-bindata/...`
1. `github.com/urfave/cli/v2`
1. `github.com/howeyc/gopass`


## Other
Welcome to issue, Thanks.