package xorml

import (
	"fmt"
	"strings"

	"github.com/gzjjyz/logger"
	"xorm.io/xorm/log"
)

type Logger struct {
	level   log.LogLevel
	showSQL bool
}

var _ log.Logger = &Logger{}

func assembleLog(v ...interface{}) (string, []interface{}) {
	if len(v) == 1 {
		return replaceFormatPlaceholder(fmt.Sprintf("%s", v[0])), nil
	}
	return fmt.Sprintf("%s", v[0]), v[1:]
}

func replaceFormatPlaceholder(format string) string {
	return strings.ReplaceAll(format, "?", "%s")
}

func NewLogger(defLogger log.LogLevel, showSQL bool) *Logger {
	return &Logger{
		level:   defLogger,
		showSQL: showSQL,
	}
}

func (c *Logger) Debug(v ...interface{}) {
	if c.level <= log.LOG_DEBUG {
		format, args := assembleLog(v)
		logger.LogDebug(format, args...)
	}
	return
}

func (c *Logger) Debugf(format string, v ...interface{}) {
	format = replaceFormatPlaceholder(format)
	if c.level <= log.LOG_DEBUG {
		logger.LogDebug(format, v...)
	}
	return
}

func (c *Logger) Error(v ...interface{}) {
	if c.level <= log.LOG_ERR {
		format, args := assembleLog(v)
		logger.LogError(format, args...)
	}
	return
}

func (c *Logger) Errorf(format string, v ...interface{}) {
	format = replaceFormatPlaceholder(format)
	if c.level <= log.LOG_ERR {
		logger.LogError(format, v...)
	}
	return
}

func (c *Logger) Info(v ...interface{}) {
	if c.level <= log.LOG_INFO {
		format, args := assembleLog(v)
		logger.LogInfo(format, args...)
	}
	return
}

func (c *Logger) Infof(format string, v ...interface{}) {
	format = replaceFormatPlaceholder(format)
	if c.level <= log.LOG_INFO {
		logger.LogInfo(format, v...)
	}
	return
}

func (c *Logger) Warn(v ...interface{}) {
	if c.level <= log.LOG_WARNING {
		format, args := assembleLog(v)
		logger.LogWarn(format, args...)
	}
	return
}

func (c *Logger) Warnf(format string, v ...interface{}) {
	format = replaceFormatPlaceholder(format)
	if c.level <= log.LOG_WARNING {
		logger.LogWarn(format, v...)
	}
	return
}

func (c *Logger) Level() log.LogLevel {
	return c.level
}

func (c *Logger) SetLevel(l log.LogLevel) {
	c.level = l
	return
}

func (c *Logger) ShowSQL(show ...bool) {
	if len(show) == 0 {
		c.showSQL = true
		return
	}
	c.showSQL = show[0]
}

func (c *Logger) IsShowSQL() bool {
	return c.showSQL
}
