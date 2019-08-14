package zaplog

import (
	"os"
	"strings"

	"github.com/ThreeKing2018/k3log/conf"
	"github.com/davecgh/go-spew/spew"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Log struct {
	logger  *zap.SugaredLogger
	atom    zap.AtomicLevel
	options *conf.Options
}

func parseLevel(level conf.Level) zapcore.Level {
	switch level {
	case conf.DebugLevel:
		return zapcore.DebugLevel
	case conf.InfoLevel:
		return zapcore.InfoLevel
	case conf.WarnLevel:
		return zapcore.WarnLevel
	case conf.ErrorLevel:
		return zapcore.ErrorLevel
	case conf.PanicLevel:
		return zapcore.PanicLevel
	case conf.FatalLevel:
		return zapcore.FatalLevel
	}

	return zapcore.DebugLevel
}

var encoderConfig = zapcore.EncoderConfig{
	// Keys can be anything except the empty string.
	TimeKey:        "time",
	LevelKey:       "level",
	NameKey:        "flag",
	CallerKey:      "file",
	MessageKey:     "msg",
	StacktraceKey:  "stack",
	LineEnding:     zapcore.DefaultLineEnding,
	EncodeLevel:    zapcore.CapitalLevelEncoder,
	EncodeTime:     zapcore.ISO8601TimeEncoder,
	EncodeDuration: zapcore.SecondsDurationEncoder,
	EncodeCaller:   zapcore.ShortCallerEncoder,
}

func New(opts ...conf.Option) *Log {
	o := &conf.Options{
		Filename:   conf.Filename,
		LogLevel:   conf.LogLevel,
		MaxSize:    conf.MaxSize,
		MaxAge:     conf.MaxAge,
		Stacktrace: conf.Stacktrace,
		IsStdOut:   conf.IsStdOut,
		//ProjectName: conf.ProjectName,
		LogType: conf.LogNormalType,
	}

	for _, opt := range opts {
		opt(o)
	}

	var writers = []zapcore.WriteSyncer{}
	osfileout := zapcore.AddSync(&lumberjack.Logger{
		Filename:   o.Filename,
		MaxSize:    o.MaxSize, // megabytes
		MaxBackups: 3,
		MaxAge:     o.MaxAge, // days
		LocalTime:  true,
		Compress:   false,
	})
	if o.IsStdOut {
		writers = append(writers, os.Stdout)
	}
	writers = append(writers, osfileout)
	w := zapcore.NewMultiWriteSyncer(writers...)

	atom := zap.NewAtomicLevel()
	atom.SetLevel(parseLevel(o.LogLevel)) //改变日志级别

	var enc zapcore.Encoder
	if o.LogType == conf.LogNormalType {
		enc = zapcore.NewConsoleEncoder(encoderConfig)
	} else {
		enc = zapcore.NewJSONEncoder(encoderConfig)
	}
	core := zapcore.NewCore(
		//这里控制json 或者不是json 类型
		enc,
		w,
		atom,
	)
	logger := zap.New(
		core,
		zap.AddStacktrace(parseLevel(o.Stacktrace)),
		zap.AddCaller(),
		zap.AddCallerSkip(2))

	if o.ProjectName != "" {
		logger = logger.With(zap.String(conf.ProjectKey, o.ProjectName))
	}
	loggerSugar := logger.Sugar()
	return &Log{logger: loggerSugar, atom: atom, options: o}

}

//拼接完整的数组
func CoupArray(kv []interface{}) []interface{} {
	if len(kv)%2 != 0 {
		kv = append(kv, kv[len(kv)-1])
		kv[len(kv)-2] = "default"
	}
	return kv
}

func (l *Log) Sync() {
	l.logger.Sync()
}

func (l *Log) SetLogLevel(level conf.Level) {
	l.atom.SetLevel(parseLevel(level))
}
func (l *Log) Debug(keysAndValues ...interface{}) {
	l.logger.Debugw("", CoupArray(keysAndValues)...)
}
func (l *Log) Info(keysAndValues ...interface{}) {
	l.logger.Infow("", CoupArray(keysAndValues)...)
}
func (l *Log) Warn(keysAndValues ...interface{}) {
	l.logger.Warnw("", CoupArray(keysAndValues)...)
}
func (l *Log) Error(keysAndValues ...interface{}) {
	l.logger.Errorw("", CoupArray(keysAndValues)...)
}
func (l *Log) Panic(keysAndValues ...interface{}) {
	l.logger.Panicw("", CoupArray(keysAndValues)...)
}
func (l *Log) Fatal(keysAndValues ...interface{}) {
	l.logger.Fatalw("", CoupArray(keysAndValues)...)
}
func (l *Log) Dump(keysAndValues ...interface{}) {
	arr := CoupArray(keysAndValues)
	for k, v := range arr {
		if k%2 == 0 {
			arr[k] = v
		} else {
			arr[k] = strings.Replace(spew.Sdump(v), "\n", "", -1)
		}
	}
	l.logger.Debugw("Dump", arr...)
}
