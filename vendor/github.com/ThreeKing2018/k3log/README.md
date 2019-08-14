# k3log
> **一款开箱即用且高效,快捷,安全的golang日志,基于uber zap
> <br/>实现日志切割,日志过期时间,动态改变日志的打印级别**

## 导航
1. [介绍](#介绍)
1. [日志级别](#日志级别)
1. [设置参数](#设置参数)
1. [使用方法](#使用方法)
1. [动态改变日志的打印级别](#动态改变日志的打印级别)
1. [效率](#效率)
1. [Dump的使用](#dump的使用)

### 介绍

取名Three King Log

- 由uber zap 日志扩展而来
- 实现分隔,异步,动态级别打印,json/txt
- 以key-value形式打印日志,适合项目里使用
- 加入Dump打印数据详细类型结构,融入[go-spew](https://github.com/davecgh/go-spew)调度利器

### 文档参考
- https://godoc.org/github.com/ThreeKing2018/k3log

### 安装
```
go get -u github.com/ThreeKing2018/k3log
```

### 日志级别
- `Debug` 调度使用, 程序继续运行
- `Info` 提示使用, 程序继续运行
- `Warn` 警告使用, 程序继续运行
- `Error` 错误使用, 程序继续运行
- `Panic` 恐慌的,退出函数,不会退出应用唾弃,会执行defer
- `Fatal` 致命的,退出应用程序,不会执行defer,因为底层多一个os.Exit
- `Dump` 打印数据类型,方便调度,级别为:Debug

[TOP](#k3log)

### 设置参数
- `WithFilename`    日志保存路径
- `WithLogLevel`    日志记录级别
- `WithMaxSize`     日志分割的尺寸 MB
- `WithMaxAge`      分割日志保存的时间 day
- `WithStacktrace`  记录堆栈的级别
- `WithIsStdOut`    是否标准输出console输出
- `WithProjectName` 项目名称
- `WithLogType`     日志类型,普通 或 json

[TOP](#k3log)

### 使用方法

- 简单使用

```golang
Debug("debug日志", 1)
Info("info日志", 2)
Warn("warn日志", 3)
Error("error日志", 4)
Panic("panic", 5)
Fatal("fatal", 6)
Dump("dump", 7)
```

[TOP](#k3log)

- 开发使用

```
NewDevelopment("dev", "log.txt")
defer Sync()
Info("Info", "dev")
```

- 开发使用2

```golang
SetLogger(conf.WithIsStdOut(true),
		conf.WithLogType(conf.LogJsontype))
	Debug("self test", 100)
```

- 开发使用3

```golang
SetLogger(conf.WithLogType(conf.LogJsontype), //打印json格式
    conf.WithProjectName("Zelog日志"),          //设置项目名称
    conf.WithFilename("log.txt"),             //设置输出文件名,或输出的路径
Debug("debug日志", 1)
Info("info日志", 2)
Warn("warn日志", 3)
Error("error日志", 4)
Panic("panic", 5)
Fatal("fatal", 6)
```

[TOP](#k3log)

- 生产使用
```golang
NewProducttion("pro", "log.txt")
defer Sync()
Error("pro", "ok")

```

- 生产使用2

```golang
SetLogger(conf.WithLogType(conf.LogJsontype), //打印json格式
		conf.WithProjectName("Zelog日志"),          //设置项目名称
		conf.WithFilename("log.txt"),             //设置输出文件名,或输出的路径
		conf.WithLogLevel(conf.InfoLevel),        //设置日志级别,默认debug
		conf.WithMaxAge(30),                      //日志保存天数,默认30天
		conf.WithMaxSize(512),                    //多少M进行分隔日志,默认100M
		conf.WithIsStdOut(false)) //是否同时输出控制台
defer Sync()
Debug("debug日志", 1)
Info("info日志", 2)
Warn("warn日志", 3)
Error("error日志", 4)
Panic("panic", 5)
Fatal("fatal", 6)
```
[TOP](#k3log)

### 动态改变日志的打印级别
```golang
Info("aa", 11)
SetLogLevel(conf.InfoLevel)
Info("info", 100)
Warn("warn", 200)
SetLogLevel(conf.ErrorLevel)
Info("info-100", 300) //这个无法输出,因为上面设置日志级别为:error
Error("err", 400)
```

### 效率
```json
runtime.GOMAXPROCS(runtime.NumCPU())
BenchmarkInfo-4   	  100000	     10776 ns/op
BenchmarkInfo-4   	  200000	     12442 ns/op
```

[TOP](#k3log)


### Dump的使用
- 级别为:debug

```
type s struct {
    Name string
    Age int
}
SetLogger(conf.WithIsStdOut(true))
Dump("name", "dump", "s", s{Name:"k3", Age: 2})

//{"name": "(string) (len=4) \"dump\"", "s": "(k3log.s) { Name: (string) (len=2) \"k3\", Age: (int) 2}"}
```
