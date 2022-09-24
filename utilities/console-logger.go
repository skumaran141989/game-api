package utilities

import "fmt"

type ConsoleLogger struct{}

func (*ConsoleLogger) Error(msg string, exception error) {
	fmt.Println("Exception:: " + exception.Error())
	fmt.Println("Error-message:: " + msg)
}

func (*ConsoleLogger) Info(msg string) {
	fmt.Println("Info:: " + msg)
}

func (*ConsoleLogger) Debug(msg string) {
	fmt.Println("Debug:: " + msg)
}

var consoleLogger *ConsoleLogger

func GetConsoleLoggerInstance() *ConsoleLogger {
	if consoleLogger == nil {
		consoleLogger = &ConsoleLogger{}
	}
	return consoleLogger
}
