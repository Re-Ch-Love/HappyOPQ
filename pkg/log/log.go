package log

import (
	"fmt"
	"log"
	"os"
)

type logLevel int

//goland:noinspection GoUnusedConst
const (
	NoneLevel logLevel = iota
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

var (
	DebugLevelPrefix = "[D]"
	InfoLevelPrefix  = "[I]"
	WarnLevelPrefix  = "[W]"
	ErrorLevelPrefix = "[E]"
	FatalLevelPrefix = "[F]"
)

func init() {
	Level = DebugLevel
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	log.SetOutput(os.Stdout)
}

var Level logLevel

//goland:noinspection GoUnusedExportedFunction
func DebugF(format string, v ...interface{}) {
	if Level <= DebugLevel {
		if Level <= DebugLevel {
			outputF(DebugLevelPrefix, format, v...)
		}
	}
}

//goland:noinspection GoUnusedExportedFunction
func Debug(v ...interface{}) {
	if Level <= DebugLevel {
		output(DebugLevelPrefix, v...)
	}
}

//goland:noinspection GoUnusedExportedFunction
func Info(v ...interface{}) {
	if Level <= InfoLevel {
		output(InfoLevelPrefix, v...)
	}
}

//goland:noinspection GoUnusedExportedFunction
func InfoF(format string, v ...interface{}) {
	if Level <= InfoLevel {
		outputF(InfoLevelPrefix, format, v...)
	}
}

//goland:noinspection GoUnusedExportedFunction
func Warn(v ...interface{}) {
	if Level <= WarnLevel {
		output(WarnLevelPrefix, v...)
	}
}

//goland:noinspection GoUnusedExportedFunction
func WarnF(format string, v ...interface{}) {
	if Level <= WarnLevel {
		outputF(WarnLevelPrefix, format, v...)
	}
}

//goland:noinspection GoUnusedExportedFunction
func Error(v ...interface{}) {
	if Level <= ErrorLevel {
		output(ErrorLevelPrefix, v...)
	}
}

//goland:noinspection GoUnusedExportedFunction
func ErrorF(format string, v ...interface{}) {
	if Level <= ErrorLevel {
		outputF(ErrorLevelPrefix, format, v...)
	}
}

//goland:noinspection GoUnusedExportedFunction
func Fatal(v ...interface{}) {
	if Level <= FatalLevel {
		output(FatalLevelPrefix, v...)
		os.Exit(1)
	}
}

//goland:noinspection GoUnusedExportedFunction
func FatalF(format string, v ...interface{}) {
	if Level <= FatalLevel {
		outputF(FatalLevelPrefix, format, v...)
	}
}

func output(prefix string, v ...interface{}) {
	log.Println(append([]interface{}{prefix}, v...)...)
}

func outputF(prefix string, format string, v ...interface{}) {
	log.Println(append([]interface{}{prefix}, fmt.Sprintf(format, v...))...)
}
