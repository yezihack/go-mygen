package conf

//日志级别
type Level int8

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	PanicLevel
	FatalLevel
)

func (level Level) String() string {
	switch level {
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warning"
	case ErrorLevel:
		return "error"
	case FatalLevel:
		return "fatal"
	case PanicLevel:
		return "panic"
	}
	return "unknown"
}

const LogJsontype = "json"
const LogNormalType = "normal"

//默认参数
const (
	Filename    string = "./log/default.log" //日志保存路径 //需要设置程序当前运行路径
	LogLevel    Level  = DebugLevel          //日志记录级别
	MaxSize     int    = 100                 //日志分割的尺寸 MB
	MaxAge      int    = 30                  //分割日志保存的时间 day
	Stacktrace  Level  = PanicLevel          //记录堆栈的级别
	IsStdOut    bool   = true                //是否标准输出console输出
	ProjectName string = "ZeLog"             //项目名称
	ProjectKey  string = "service"           //
	LogType            = LogNormalType
)
