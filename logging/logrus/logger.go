package logrus

import (
	"os"

	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

type Logger struct {
	logger *log.Logger
}

// Fields type, used to pass to `WithFields`.
type Fields map[string]interface{}

func New(isDebug bool) *Logger {

	logLevel := log.InfoLevel
	if isDebug {
		log.SetReportCaller(true)
		logLevel = log.DebugLevel
	}

	logger := &log.Logger{
		Out: os.Stderr,
		Formatter: &prefixed.TextFormatter{
			DisableColors:   false,
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
			ForceFormatting: true,
		},
		Hooks: make(log.LevelHooks),
		Level: logLevel,
	}

	// if isDebug {
	// 	logger.SetReportCaller(true)
	// }

	return &Logger{logger: logger}
}

func (l *Logger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *Logger) WithFields(fields Fields) *log.Entry {
	return l.logger.WithFields(log.Fields(fields))
}

func (l *Logger) Trace(args ...interface{}) {
	l.logger.Trace(args...)
}

func (l *Logger) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l *Logger) Print(args ...interface{}) {
	l.logger.Print(args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l *Logger) Warning(args ...interface{}) {
	l.logger.Warning(args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

func (l *Logger) Panic(args ...interface{}) {
	l.logger.Panic(args...)
}

// Tracef logs a message at level Trace on the standard logger.
func (l *Logger) Tracef(format string, args ...interface{}) {
	l.logger.Tracef(format, args...)
}

// Debugf logs a message at level Debug on the standard logger.
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

// Printf logs a message at level Info on the standard logger.
func (l *Logger) Printf(format string, args ...interface{}) {
	l.logger.Printf(format, args...)
}

// Infof logs a message at level Info on the standard logger.
func (l *Logger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

// Warnf logs a message at level Warn on the standard logger.
func (l *Logger) Warnf(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

// Warningf logs a message at level Warn on the standard logger.
func (l *Logger) Warningf(format string, args ...interface{}) {
	l.logger.Warningf(format, args...)
}

// Errorf logs a message at level Error on the standard logger.
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

// Panicf logs a message at level Panic on the standard logger.
func (l *Logger) Panicf(format string, args ...interface{}) {
	l.logger.Panicf(format, args...)
}

// Fatalf logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}

// Traceln logs a message at level Trace on the standard logger.
func (l *Logger) Traceln(args ...interface{}) {
	l.logger.Traceln(args...)
}

// Debugln logs a message at level Debug on the standard logger.
func (l *Logger) Debugln(args ...interface{}) {
	l.logger.Debugln(args...)
}

// Println logs a message at level Info on the standard logger.
func (l *Logger) Println(args ...interface{}) {
	l.logger.Println(args...)
}

// Infoln logs a message at level Info on the standard logger.
func (l *Logger) Infoln(args ...interface{}) {
	l.logger.Infoln(args...)
}

// Warnln logs a message at level Warn on the standard logger.
func (l *Logger) Warnln(args ...interface{}) {
	l.logger.Warnln(args...)
}

// Warningln logs a message at level Warn on the standard logger.
func (l *Logger) Warningln(args ...interface{}) {
	l.logger.Warningln(args...)
}

// Errorln logs a message at level Error on the standard logger.
func (l *Logger) Errorln(args ...interface{}) {
	l.logger.Errorln(args...)
}

// Panicln logs a message at level Panic on the standard logger.
func (l *Logger) Panicln(args ...interface{}) {
	l.logger.Panicln(args...)
}

// Fatalln logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func (l *Logger) Fatalln(args ...interface{}) {
	l.logger.Fatalln(args...)
}
