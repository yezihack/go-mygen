![](https://img.shields.io/badge/go--mygen-tools-orange?style=flat-square&logo=appveyor)
![](https://img.shields.io/badge/download-4M-green?style=flat-square&logo=appveyor)
![](https://img.shields.io/badge/release-linux%2Cmac%2Cwin-blue?style=flat-square&logo=appveyor)
<br/>
![](https://img.shields.io/badge/go--mygen--en-3.1.0beta-green)
![](https://img.shields.io/badge/go--mygen--cn-3.0.0beta-green)
![](https://img.shields.io/github/stars/yezihack/go-mygen)
![](https://img.shields.io/github/issues/yezihack/go-mygen)
![](https://img.shields.io/github/forks/yezihack/go-mygen)
![](https://img.shields.io/github/license/yezihack/go-mygen)

[英文](README.md)|[中文](README-CN.md)

# go-mygen
> 代码生成器,避免ORM带来的性能损耗.

## debug 修复
1. 查询结果为空时，底层不再处理，交由上层处理。
```go
func getRows() (err error) {
    err = query.Scan(
        &row.Id,          
    )
    if err != nil {
        return
    }
}
func main() {
    err := getRows()
    if err == sql.ErrNoRows {
        fmt.Println("result is nil")
    }   
}
```

## 版本
1. 3.2.0 英文
1. 兼容linux,win,mac
1. 3.0.0beta为中文版本
1. 3.1.0beta为英文版本

## v3.2.0
1. 生成curd支持实体与配置存放不同目录

## 安装
> \>= go1.13.0
```
go install github.com/yezihack/go-mygen
```
## 下载
> [release](https://github.com/yezihack/go-mygen/releases/tag/3.0.0beta)

## 功能介绍
- 1.生成数据库表的markdown文档
- 2.生成golang表对应的结构实体
- 3.自定义结构体解析实体,如json,gorm,xml等
- 4.生成golang操作mysql的增删改查语句
- 5.可以自定义生成目录
- 6.清屏,退出等功能

## 使用方法
1. 查看帮助 `go-mygen help`
1. 查看版本 `go-mygen version`
1. 连接操作 `go-mygen -h localhost -P 3306 -u root -p 123456 -d default `


## 参数说明
```
-h value       数据库地址 (default: "localhost")
-P value       端口号 (default: 3306)
-u value       数据库用户名称 (default: "root")
-p value       数据库密码 (default: "root")
-c value       编码格式 (default: "utf8mb4")
-d value       数据库名称
```

## 依赖包
1. `go get -u github.com/go-bindata/go-bindata/...`

## 其它
> 发现错误,欢迎issues
```
xx>  表示输入状态
xx:  信息提示
xx>> 错误输出
```

1. go modules代理.
```
export GOPROXY=https://goproxy.io
export GOPROXY=https://gocenter.io
```

