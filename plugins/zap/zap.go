package zap

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/natefinch/lumberjack"
	"github.com/tian-yuan/iot-common/basic/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	l                              *Logger
	sp                             = string(filepath.Separator)
	errWS, warnWS, infoWS, debugWS zapcore.WriteSyncer
	debugConsoleWS                 = zapcore.Lock(os.Stdout)
	errorConsoleWS                 = zapcore.Lock(os.Stderr)
)

func init() {
	l = &Logger{
		Opts: &Options{},
	}
}

type Logger struct {
	*zap.Logger
	sync.RWMutex
	Opts      *Options `json:"opts"`
	zapConfig zap.Config
	inited    bool
}

func Init(callerSkip int, path ...string) {
	l.Lock()
	defer l.Unlock()

	if l.inited {
		return
	}

	l.loadCfg(path...)
	l.init(callerSkip)
	l.inited = true

}

func GetLogger() (ret *Logger) {
	return l
}

func (l *Logger) init(callerSkip int) {
	l.setSyncers()
	var err error

	l.Logger, err = l.zapConfig.Build(l.cores(), zap.AddCallerSkip(callerSkip))
	if err != nil {
		panic(err)
	}

	defer l.Logger.Sync()
}

func (l *Logger) loadCfg(path ...string) {
	c := config.C()

	config := c.GetConfig()
	v := config.Get(path...)
	var err error
	if v != nil {
		err = v.Scan(l.Opts)
	} else {
		err = fmt.Errorf("config not exist in center.")
	}
	if err != nil {
		panic(err)
	}

	if l.Opts.Development {
		l.zapConfig = zap.NewDevelopmentConfig()
	} else {
		l.zapConfig = zap.NewProductionConfig()
	}

	if l.Opts.OutputPaths == nil || len(l.Opts.OutputPaths) == 0 {
		l.zapConfig.OutputPaths = []string{"stdout"}
	}

	if l.Opts.ErrorOutputPaths == nil || len(l.Opts.ErrorOutputPaths) == 0 {
		l.zapConfig.OutputPaths = []string{"stderr"}
	}

	if l.Opts.LogFileDir == "" {
		l.Opts.LogFileDir, _ = filepath.Abs(filepath.Dir(filepath.Join(".")))
		l.Opts.LogFileDir += sp + "logs" + sp
	}

	if l.Opts.AppName == "" {
		l.Opts.AppName = "app"
	}

	if l.Opts.ErrorFileName == "" {
		l.Opts.ErrorFileName = "error.log"
	}

	if l.Opts.WarnFileName == "" {
		l.Opts.WarnFileName = "warn.log"
	}

	if l.Opts.InfoFileName == "" {
		l.Opts.InfoFileName = "info.log"
	}

	if l.Opts.DebugFileName == "" {
		l.Opts.DebugFileName = "debug.log"
	}

	if l.Opts.MaxSize == 0 {
		l.Opts.MaxSize = 50
	}
	if l.Opts.MaxBackups == 0 {
		l.Opts.MaxBackups = 3
	}
	if l.Opts.MaxAge == 0 {
		l.Opts.MaxAge = 30
	}
}

func (l *Logger) setSyncers() {
	f := func(fN string) zapcore.WriteSyncer {
		return zapcore.AddSync(&lumberjack.Logger{
			Filename:   l.Opts.LogFileDir + sp + l.Opts.AppName + "-" + fN,
			MaxSize:    l.Opts.MaxSize,
			MaxBackups: l.Opts.MaxBackups,
			MaxAge:     l.Opts.MaxAge,
			Compress:   true,
			LocalTime:  true,
		})
	}

	errWS = f(l.Opts.ErrorFileName)
	warnWS = f(l.Opts.WarnFileName)
	infoWS = f(l.Opts.InfoFileName)
	debugWS = f(l.Opts.DebugFileName)

	return
}

func (l *Logger) cores() zap.Option {
	l.zapConfig.EncoderConfig.TimeKey = "timestamp"
	l.zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	fileEncoder := zapcore.NewJSONEncoder(l.zapConfig.EncoderConfig)
	consoleEncoder := zapcore.NewConsoleEncoder(l.zapConfig.EncoderConfig)

	errPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl > zapcore.WarnLevel && zapcore.WarnLevel-l.zapConfig.Level.Level() > -1
	})
	warnPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.WarnLevel && zapcore.WarnLevel-l.zapConfig.Level.Level() > -1
	})
	infoPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.InfoLevel && zapcore.InfoLevel-l.zapConfig.Level.Level() > -1
	})
	debugPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.DebugLevel && zapcore.DebugLevel-l.zapConfig.Level.Level() > -1
	})

	cores := []zapcore.Core{
		zapcore.NewCore(fileEncoder, errWS, errPriority),
		zapcore.NewCore(fileEncoder, warnWS, warnPriority),
		zapcore.NewCore(fileEncoder, infoWS, infoPriority),
		zapcore.NewCore(fileEncoder, debugWS, debugPriority),
		zapcore.NewCore(consoleEncoder, errorConsoleWS, errPriority),
		zapcore.NewCore(consoleEncoder, debugConsoleWS, warnPriority),
		zapcore.NewCore(consoleEncoder, debugConsoleWS, infoPriority),
		zapcore.NewCore(consoleEncoder, debugConsoleWS, debugPriority),
	}

	return zap.WrapCore(func(c zapcore.Core) zapcore.Core {
		return zapcore.NewTee(cores...)
	})
}

func (l *Logger) Log(v ...interface{}) {
	l.Info(fmt.Sprint(v...))
}

func (l *Logger) Logf(format string, v ...interface{}) {
	l.Info(fmt.Sprintf(format, v...))
}
