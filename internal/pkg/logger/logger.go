package logger

type Logger interface {
	Debug(msg string)
	Info(msg string)
	Infof(msg string, args ...any)
	Infow(msg string, w ...any)
	Warn(msg string)
	Error(msg string)
	Errorf(msg string, args ...any)
	Errorw(msg string, w ...any)
}
