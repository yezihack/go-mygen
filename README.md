[![](https://img.shields.io/badge/go--mygen-tools-orange?style=flat-square&logo=appveyor)](https://github.com/yezihack/go-mygen/releases)
[![](https://img.shields.io/badge/download-4M-green?style=flat-square&logo=appveyor)](https://github.com/yezihack/go-mygen/releases)
[![](https://img.shields.io/badge/release-linux%2Cmac%2Cwin-blue?style=flat-square&logo=appveyor)](https://github.com/yezihack/go-mygen/releases)
<br/>
[![](https://img.shields.io/badge/go--mygen--en-3.1.0beta-green)](https://github.com/yezihack/go-mygen/tree/3.1.0beta)
[![](https://img.shields.io/badge/go--mygen--cn-3.0.0beta-green)](https://github.com/yezihack/go-mygen/tree/3.0.0beta)
[![](https://img.shields.io/github/stars/yezihack/go-mygen)](https://github.com/yezihack/go-mygen/stargazers)
[![](https://img.shields.io/github/issues/yezihack/go-mygen)](https://github.com/yezihack/go-mygen/issues)
[![](https://img.shields.io/github/forks/yezihack/go-mygen)](https://github.com/yezihack/go-mygen/network/members)
[![](https://img.shields.io/github/license/yezihack/go-mygen)](https://github.com/yezihack/go-mygen/blob/3.1.0beta/LICENSE)

[English](README.md)|[中文](README-CN.md)

![](assets/img/golang.png)

# go-mygen
> Code generator tools with MYSQL,avoid ORM performance penalty


## Install
> \>= go1.13.0
```
go get -u -x github.com/yezihack/go-mygen
```

## debug fix 
1. When the query result is null, the bottom layer is no longer processed and the upper layer is processed.
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

## v3.3.0 version
1. Compatible with linux,win,mac
1. add stack error
1. different int32 or int64

### v3.2.0
1. Operation 3 sequence, support entity and SQL, configure to store different directories


## Using v3 releases
> [release](https://github.com/yezihack/go-mygen/releases/tag/3.0.0beta)

## fix 
get 

## QuickStart
1. See Help `go-mygen help`
1. See Version `go-mygen v`
1. Connection `go-mygen -h localhost -P 3306 -u root -p 123456 -d default `

## Features
- Generate the markdown document for the database table
- Generate the structure entities for the golang table
- Custom structure parsing entities, such as json,gorm, XML, etc
- Generate golang operation mysql add, delete, modify and select code statements
- Custom build directory
- Clear screen, exit and other functions

## Parameters info
```
-h value       Database address (default: "localhost")
-P value       port number (default: 3306)
-u value       Database user name (default: "root")
-p value       Database password (default: "root")
-c value       Coding format (default: "utf8mb4")
-d value       Database name
```

## Other
> If an error is found, please welcome issues
```
xx>  input state
xx:  message tips
xx>> error output
```

## package 
1. `go get -u github.com/go-bindata/go-bindata/...`
1. `github.com/urfave/cli/v2`