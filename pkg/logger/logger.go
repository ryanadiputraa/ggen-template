package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Logger interface {
	Info(v ...any)
	Warn(v ...any)
	Error(v ...any)
	Fatal(v ...any)
}

type logger struct {
	logger *log.Logger
}

func New(location *time.Location, out *os.File) Logger {
	l := log.New(out, "", log.Ldate|log.Ltime)
	l.SetFlags(l.Flags() | log.LUTC)
	l.SetPrefix(fmt.Sprintf("[%v] ", location.String()))
	return &logger{logger: l}
}

func (l *logger) Info(v ...any) {
	l.logger.Println(append([]any{"[INFO]:"}, v...)...)
}

func (l *logger) Warn(v ...any) {
	l.logger.Println(append([]any{"[WARN]:"}, v...)...)
}

func (l *logger) Error(v ...any) {
	l.logger.Println(append([]any{"[ERROR]:"}, v...)...)
}

func (l *logger) Fatal(v ...any) {
	l.logger.Fatal(append([]any{"[FATAL]:"}, v...)...)
}
