package main

import "log"

func main() {
	log.Println("Starting application...")

	result := performOperation(5, 2)
	log.Printf("Result of operation: %d\n", result)

	err := simulateError()
	if err != nil {
		log.Printf("Error occurred: %v\n", err)
	}

	log.Println("Application finished.")
}

func performOperation(a, b int) int {
	log.Println("Performing operation...")
	return a + b
}

func simulateError() error {
	log.Println("Simulating error...")
	return nil
}
