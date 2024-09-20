package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	// Check if a file path was provided as a command-line argument
	if len(os.Args) < 2 {
		log.Fatal("Please provide a file path")
	}

	// Open the file specified as the first command-line argument
	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open the file: %v", err)
	}
	defer file.Close() // Ensure the file is closed after the function returns

	// Create a new scanner to read the file
	scanner := bufio.NewScanner(file)

	// Read each line of the file and print it
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// Check for errors that occurred during the scan
	if scanner.Err() != nil {
		log.Fatalf("Error during scanning: %v", err)
	}
}
