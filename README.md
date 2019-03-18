# go-mygen
> 这个是一个命令行工具,
golang操作mysql的便捷工具, 快速生成golang结构语句,快捷生成增删改查等等

## 安装
```
go install github.com/yezihack/go-mygen
```

## 参数说明
```
-h value       数据库地址 (default: "localhost")
-P value       端口号 (default: 3306)
-u value       数据库用户名称 (default: "root")
-p value       数据库密码 (default: "123456")
-c value       编码格式 (default: "utf8mb4")
-d value       *数据库名称
--help         显示命令帮助
--version, -v  显示版本号
```

## 使用说明
```
gomygen -h localhost -P 3306 -u root -p 123456 -d default
```

## Makefile 使用
```
git clone github.com/yezihack/go-mygen # 下载项目
cd go-mygen # 进入项目
make dev  #进行自动编辑生成可执行文件
gomygen/output/go-mygen # 回车运行,并查看帮助
gomygen/output/go-mygen -h localhost -P 3306 -u root -p 123456 -d default #连接数据库

```

## todo
- 设置输出路径 (后期补充)