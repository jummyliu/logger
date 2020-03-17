package logger

// Logger logger interface
type Logger interface {
	Log(Level, string, ...interface{})
	LogEmerg(string, ...interface{})
	LogAlter(string, ...interface{})
	LogCrit(string, ...interface{})
	LogError(string, ...interface{})
	LogWarning(string, ...interface{})
	LogNotice(string, ...interface{})
	LogInfo(string, ...interface{})
	LogDebug(string, ...interface{})
}

// Level log level
type Level int64

// Log level
const (
	LevelEmerg Level = iota
	LevelAlter
	LevelCrit
	LevelError
	LevelWarning
	LevelNotice
	LevelInfo
	LevelDebug

	// LevelNum // count
)

// LogNameMap log name map
var LogNameMap = map[Level]string{
	LevelEmerg:   "emerg",
	LevelAlter:   "alert",
	LevelCrit:    "crit",
	LevelError:   "error",
	LevelWarning: "warning",
	LevelNotice:  "notice",
	LevelInfo:    "info",
	LevelDebug:   "debug",
}

const defaultLevel = LevelError

var GetLevelByName func(string) Level

func init() {
	GetLevelByName = LevelGetter()
}

func LevelGetter() func(string) Level {
	m := map[string]Level{}
	for k, v := range LogNameMap {
		m[v] = k
	}
	return func(levelName string) Level {
		if level, ok := m[levelName]; ok {
			return level
		}
		return defaultLevel
	}
}
