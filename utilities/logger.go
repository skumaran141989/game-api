package utilities

type Logger interface {
	Error(msg string, exception error)
	Info(msg string)
	Debug(msg string)
}
