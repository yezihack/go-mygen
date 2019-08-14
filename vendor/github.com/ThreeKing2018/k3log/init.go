package k3log

import (
	"github.com/ThreeKing2018/k3log/conf"
	"github.com/ThreeKing2018/k3log/plugins/zaplog"
)

//默认
var l Loger = zaplog.New()

//设置
func SetLogger(opts ...conf.Option) {
	l = zaplog.New(opts...)
}

//快捷使用,开发使用
func NewDevelopment(projectName, filePath string) {
	SetLogger(conf.WithProjectName(projectName),
		conf.WithFilename(filePath),
		conf.WithLogType(conf.LogJsontype),
		conf.WithIsStdOut(true))
}

//快捷使用,生产使用
func NewProduction(projectName, filePath string) {
	SetLogger(conf.WithProjectName(projectName),
		conf.WithLogType(conf.LogJsontype),
		conf.WithFilename(filePath),
		conf.WithLogLevel(conf.ErrorLevel),
		conf.WithIsStdOut(false))
}

//目前只有zap生效
func SetLogLevel(level conf.Level) {
	l.SetLogLevel(level)
}

//日志同步写入
//目前只有zap生效
func Sync() {
	l.Sync()
}

//日志等级 调试时使用
func Debug(keysAndValues ...interface{}) {
	l.Debug(keysAndValues...)
}

//日志等级 提示时使用
func Info(keysAndValues ...interface{}) {
	l.Info(keysAndValues...)
}

//日志等级 警告时使用
func Warn(keysAndValues ...interface{}) {
	l.Warn(keysAndValues...)
}

//日志等级 错误时使用
func Error(keysAndValues ...interface{}) {
	l.Error(keysAndValues...)
}

//日志等级 恐慌时使用
func Panic(keysAndValues ...interface{}) {
	l.Panic(keysAndValues...)
}

//日志等级 致命时使用
func Fatal(keysAndValues ...interface{}) {
	l.Fatal(keysAndValues...)
}

//日志等级 详细结构类型,调试利器
func Dump(keysAndValues ...interface{}) {
	l.Dump(keysAndValues...)
}
