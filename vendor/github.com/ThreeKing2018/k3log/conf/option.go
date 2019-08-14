package conf

type Option func(*Options)

type Options struct {
	Filename    string //日志保存路径
	LogLevel    Level  //日志记录级别
	MaxSize     int    //日志分割的尺寸 MB
	MaxAge      int    //分割日志保存的时间 day
	Stacktrace  Level  //记录堆栈的级别
	IsStdOut    bool   //是否标准输出console输出
	ProjectName string //项目名称
	LogType     string //日志类型,普通 或 json
}

func WithLogType(logtype string) Option {
	return func(o *Options) {
		o.LogType = logtype
	}
}

func WithFilename(logpath string) Option {
	return func(o *Options) {
		o.Filename = logpath
	}
}

func WithLogLevel(loglevel Level) Option {
	return func(o *Options) {
		o.LogLevel = loglevel
	}
}

func WithMaxSize(maxsize int) Option {
	return func(o *Options) {
		o.MaxSize = maxsize
	}
}

func WithMaxAge(maxage int) Option {
	return func(o *Options) {
		o.MaxAge = maxage
	}
}

func WithStacktrace(stacktrace Level) Option {
	return func(o *Options) {
		o.Stacktrace = stacktrace
	}
}

func WithIsStdOut(isstdout bool) Option {
	return func(o *Options) {
		o.IsStdOut = isstdout
	}
}

func WithProjectName(projectname string) Option {
	return func(o *Options) {
		o.ProjectName = projectname
	}
}
