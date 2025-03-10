package logger

type Logger interface {
	Error(format string, a ...any)
	Warning(format string, a ...any)
	Info(format string, a ...any)
	Debug(format string, a ...any)
	Trace(format string, a ...any)
}
