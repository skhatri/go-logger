package logging

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

var _cache map[string]*Logger

func init() {
  log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(logLevel())
	_cache = make(map[string]*Logger, 0)
}

func logLevel() log.Level {
	var logLevel log.Level
  lgLevel := os.Getenv("LOG_LEVEL")

	switch lgLevel {
	case "debug":
		logLevel = log.DebugLevel
	case "error":
		logLevel = log.ErrorLevel
	case "warn":
		logLevel = log.WarnLevel
	default:
		logLevel = log.InfoLevel
	}
	return logLevel
}

func NewLogger(source string) *Logger {
	if cachedLogger, ok := _cache[source]; ok {
		return cachedLogger
	} else {
		logger := &Logger{
			source: source,
		}
		_cache[source] = logger
		return logger
	}
}

type Logger struct {
	source string
}

func (lgr *Logger) WithTask(taskName string) *LoggerTask {
	task := LoggerTask{}
	task._logger = log.WithFields(log.Fields{
		"source": lgr.source,
		"task":   taskName,
	})
	return &task
}

type LoggerTask struct {
	_logger *log.Entry
}

func (task *LoggerTask) WithAttributes(data map[string]interface{}) *LoggerTask {
	task._logger = task._logger.WithFields(data)
	return task
}

func (task *LoggerTask) WithAttribute(key string, value interface{}) *LoggerTask {
	task._logger = task._logger.WithField(key, value)
	return task
}

func (task *LoggerTask) Info() {
	task._logger.Info()
}

func (task *LoggerTask) WithInfo(msg string) {
	task._logger.Info("msg", msg)
}

func (task *LoggerTask) WithMessage(msg string, args ...any) *LoggerTask {
	task._logger = task._logger.WithField("msg", fmt.Sprintf(msg, args...))
	return task
}

func (task *LoggerTask) Warn() {
	task._logger.Warn()
}

func (task *LoggerTask) WithWarn(warning string) {
	task._logger.Warn("warning ", warning)
}

func (task *LoggerTask) Error() {
	task._logger.Error()
}

func (task *LoggerTask) WithError(err error) {
	task._logger.Error("error ", err.Error())
}

func (task *LoggerTask) Debug() {
	task._logger.Debug()
}

func (task *LoggerTask) Fatalf(s string, args ...any) {
	task._logger.Fatalf(s, args...)
}

func (task *LoggerTask) Fatal(err error) {
	task._logger.Fatal("fatal error ", err.Error())
}


