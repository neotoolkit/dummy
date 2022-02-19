package logger

type Logger interface {
	Debug(msg string)
	Info(msg string)
	Infof(msg string, args ...interface{})
	Infow(msg string, w ...interface{})
	Warn(msg string)
	Error(msg string)
	Errorf(msg string, args ...interface{})
	Errorw(msg string, w ...interface{})
}
