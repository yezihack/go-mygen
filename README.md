# golang 操作mysql的命令行
> 这个是一个命令行工具,
golang操作mysql的便捷工具, 快速生成表结束,快捷生成增删改查

## 安装
```
go install github.com/yezihack/gm2m
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
gm2m -h localhost -P 3306 -u root -p 123456 -d default
```

## 实例
```
序号:1 生成表markdown文档
序号:2 生成表结构数据
序号:3 生成CURD增删改查
序号:4 设置结构体的映射名称
序号:7, c, clear 清屏
序号:8, h 查看帮助
序号:9, e, exit 退出
```