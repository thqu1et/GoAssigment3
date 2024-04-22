package main

import (
	"log"
	"os"
)

type AppLogger struct {
	logger *log.Logger
}

func NewAppLogger() *AppLogger {
	return &AppLogger{logger: log.New(os.Stdout, "", 0)}
}

func (al *AppLogger) Info(msg string, fields map[string]interface{}) {
	al.log("INFO", msg, fields)
}

func (al *AppLogger) Error(msg string, fields map[string]interface{}) {
	al.log("ERROR", msg, fields)
}

func (al *AppLogger) log(level, msg string, fields map[string]interface{}) {
	// Add common metadata
	fields["level"] = level

	logEntry := make(map[string]interface{})
	for k, v := range fields {
		logEntry[k] = v
	}
	al.logger.Println(logEntry)
}

func main() {

	logger := NewAppLogger()
	logger.Info("Starting application", nil)

	result := performOperation(5, 2, logger)
	logger.Info("Result of operation", map[string]interface{}{"result": result})

	err := simulateError(logger)
	if err != nil {
		logger.Error("Error occurred", map[string]interface{}{"error": err.Error()})
	}

	logger.Info("Application finished", nil)
}

func performOperation(a, b int, logger *AppLogger) int {
	logger.Info("Performing operation", map[string]interface{}{"operands": []int{a, b}})
	return a + b
}

func simulateError(logger *AppLogger) error {
	logger.Info("Simulating error", nil)
	return nil
}
