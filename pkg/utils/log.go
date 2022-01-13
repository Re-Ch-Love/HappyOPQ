package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type logLevel string

//goland:noinspection GoUnusedConst
const (
	LDebug logLevel = "[D]"
	LInfo           = "[I]"
	LWarn           = "[W]"
	LError          = "[E]"
	LFatal          = "[F]"
)

type logFlag int

//goland:noinspection GoUnusedConst
const (
	FTime logFlag = 1 << iota
)

type Logger struct {
	Format string
	Tag    string
}

func NewDefaultLogger() Logger {
	return Logger{
		Format: "[[Tag]] [YYYY]-[MM]-[MM]-[DD]-[SS].[NS_] [[level]] [context]",
		Tag:    "Default",
	}
}

func (l *Logger) Printf(level logLevel, format string, v ...interface{}) {
	var entry string
	strings.ReplaceAll(l.Format, "[Tag]", l.Tag)
	strings.ReplaceAll(l.Format, "[level]", string(level))
	strings.ReplaceAll(l.Format, "[YYYY]", strconv.Itoa(time.Now().Year()))
	strings.ReplaceAll(l.Format, "[MM]", fmt.Sprintf("%02d", int(time.Now().Month())))
	strings.ReplaceAll(l.Format, "[DD]", fmt.Sprintf("%02d", time.Now().Day()))
	strings.ReplaceAll(l.Format, "[SS]", strconv.Itoa(time.Now().Second()))
	strings.ReplaceAll(l.Format, "[NS_]", strconv.Itoa(time.Now().Nanosecond()/1000000))
	strings.ReplaceAll(l.Format, "[context]", fmt.Sprintf(format, v...))
	fmt.Println(entry)
}

//goland:noinspection GoUnusedExportedFunction,SpellCheckingInspection
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.Printf(LDebug, format, v...)
}

//goland:noinspection GoUnusedExportedFunction,SpellCheckingInspection
func (l *Logger) Debug(format string, v ...interface{}) {
	l.Printf(LDebug, format, v...)
}

//goland:noinspection GoUnusedExportedFunction,SpellCheckingInspection
func (l *Logger) Infof(format string, v ...interface{}) {
	l.Printf(LInfo, format, v...)
}

//goland:noinspection GoUnusedExportedFunction,SpellCheckingInspection
func (l *Logger) Info(format string, v ...interface{}) {
	l.Printf(LInfo, format, v...)
}

//goland:noinspection GoUnusedExportedFunction,SpellCheckingInspection
func (l *Logger) Warnf(format string, v ...interface{}) {
	l.Printf(LWarn, format, v...)
}

//goland:noinspection GoUnusedExportedFunction,SpellCheckingInspection
func (l *Logger) Warn(format string, v ...interface{}) {
	l.Printf(LWarn, format, v...)
}

//goland:noinspection GoUnusedExportedFunction,SpellCheckingInspection
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.Printf(LError, format, v...)
}

//goland:noinspection GoUnusedExportedFunction,SpellCheckingInspection
func (l *Logger) Error(format string, v ...interface{}) {
	l.Printf(LError, format, v...)
}

//goland:noinspection GoUnusedExportedFunction,SpellCheckingInspection
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.Printf(LFatal, format, v...)
	os.Exit(1)
}

//goland:noinspection GoUnusedExportedFunction,SpellCheckingInspection
func (l *Logger) Fatal(format string, v ...interface{}) {
	l.Printf(LFatal, format, v...)
	os.Exit(1)
}

//goland:noinspection GoUnusedExportedFunction,SpellCheckingInspection
func (l *Logger) PanicErr(v interface{}) {
	err, ok := v.(error)
	if ok {
		l.Printf(LFatal, err.Error())
	} else {
		l.Printf(LFatal, "Panic Error: %+err", v)
	}
	os.Exit(1)
}
