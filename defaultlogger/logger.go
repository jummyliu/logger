package defaultlogger

import (
	"fmt"
	"io"
	"log"

	"github.com/jummyliu/logger"
)

// DefaultLogger default logger
type DefaultLogger struct {
	Logger   *log.Logger
	MaxLevel logger.Level
}

const defaultLFlag = log.Ldate | log.Ltime | log.Lmicroseconds

const allLFlag = log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile | log.Lshortfile

// New init default logger
func New(out io.Writer, maxLevel logger.Level) logger.Logger {
	return &DefaultLogger{
		Logger:   log.New(out, "", defaultLFlag),
		MaxLevel: maxLevel,
	}
}

// Log log by level
func (l *DefaultLogger) Log(lev logger.Level, format string, params ...interface{}) {
	// 超过配置的日志级别，不输出
	if lev > l.MaxLevel {
		return
	}
	if lev == logger.LevelDebug {
		l.Logger.SetFlags(l.Logger.Flags() | log.Lshortfile)
	} else {
		l.Logger.SetFlags(l.Logger.Flags() & (allLFlag ^ log.Lshortfile))
	}
	l.Logger.Printf(fmt.Sprintf("| [%s]\t| ", logger.LogNameMap[lev])+format, params...)
}

// LogEmerg log by emerg
func (l *DefaultLogger) LogEmerg(format string, params ...interface{}) {
	l.Log(logger.LevelEmerg, format, params...)
}

// LogAlter log by Alter
func (l *DefaultLogger) LogAlter(format string, params ...interface{}) {
	l.Log(logger.LevelAlter, format, params...)
}

// LogCrit log by Crit
func (l *DefaultLogger) LogCrit(format string, params ...interface{}) {
	l.Log(logger.LevelCrit, format, params...)
}

// LogError log by error
func (l *DefaultLogger) LogError(format string, params ...interface{}) {
	l.Log(logger.LevelError, format, params...)
}

// LogWarning log by warning
func (l *DefaultLogger) LogWarning(format string, params ...interface{}) {
	l.Log(logger.LevelWarning, format, params...)
}

// LogNotice log by notice
func (l *DefaultLogger) LogNotice(format string, params ...interface{}) {
	l.Log(logger.LevelNotice, format, params...)
}

// LogInfo log by info
func (l *DefaultLogger) LogInfo(format string, params ...interface{}) {
	l.Log(logger.LevelInfo, format, params...)
}

// LogDebug log by debug
func (l *DefaultLogger) LogDebug(format string, params ...interface{}) {
	l.Log(logger.LevelDebug, format, params...)
}
