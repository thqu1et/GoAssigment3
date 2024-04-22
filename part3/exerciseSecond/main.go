package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.SetLevel(logrus.DebugLevel)

	logger.WithFields(logrus.Fields{
		"app": "exercise-app",
	}).Info("Starting application...")

	result := performOperation(5, 0)
	logger.WithFields(logrus.Fields{
		"operation": "division",
		"result":    result,
	}).Info("Operation performed")

	if err := performOperation(10, 0); err != nil {
		logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("Error performing operation")
	}

	logger.WithFields(logrus.Fields{
		"app": "exercise-app",
	}).Info("Application finished.")
}

func performOperation(a, b int) error {
	if b == 0 {
		return ErrDivisionByZero
	}
	return nil
}

var ErrDivisionByZero = NewAppError("division by zero")

type AppError struct {
	message string
}

func NewAppError(message string) *AppError {
	return &AppError{message: message}
}

func (e *AppError) Error() string {
	return e.message
}
