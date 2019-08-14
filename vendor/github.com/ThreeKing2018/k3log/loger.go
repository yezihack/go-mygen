package k3log

import (
	"github.com/ThreeKing2018/k3log/conf"
)

//定义接口
type Loger interface {
	//key value
	Debug(...interface{})   //调试的
	Info(...interface{})    //提示的
	Warn(...interface{})    //警告的
	Error(...interface{})   //错误的
	Panic(...interface{})   //恐慌的
	Fatal(...interface{})   //致命的
	Dump(...interface{})    //详细结构类型,调试利器
	Sync()                  //同步
	SetLogLevel(conf.Level) //可以随机设置日志级别的
}
