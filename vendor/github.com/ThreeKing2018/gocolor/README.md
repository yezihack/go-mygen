# golang 命令行输出带颜色的字符串
> 让你的命令行丰富多彩,输出不同颜色,自由组合

# 文档
- https://godoc.org/github.com/ThreeKing2018/gocolor

# 安装
```
go get -u github.com/ThreeKing2018/gocolor
```

# 引用包
```
import "github.com/ThreeKing2018/gocolor"
```

# 基本色
```
gocolor.Green("绿色")
gocolor.White("白色")
gocolor.Yellow("黄色")
gocolor.Red("红色")
gocolor.Blue("蓝色")
gocolor.Magenta("洋红色")
gocolor.Cyan("蓝绿色")
```

![](picture/1.jpg)

# 背景色

```
gocolor.GreenBG("绿色背景")
gocolor.WhiteBG("白色背景")
gocolor.YellowBG("黄色背景")
gocolor.RedBG("红色背景")
gocolor.BlueBG("蓝色背景")
gocolor.MagentaBG("洋红色背景")
gocolor.CyanBG("蓝绿色背景")
```

![](picture/2.jpg)

# 组合色

```
gocolor.BlueBG(SRed("红字蓝底色"))
gocolor.CyanBG(SYellow("黄字蓝绿底色"))
gocolor.MagentaBG(SRed("红字洋底色"))
gocolor.GreenBG(SMagentaBG("洋红字绿底"))
```

![](picture/3.jpg)