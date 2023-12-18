package logger

type Fields map[string]interface{}

type ILogger interface {
	Debug(args ...interface{})
}
