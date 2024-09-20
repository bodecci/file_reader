package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Create a new scanner to read the file
	scanner := bufio.NewScanner(file)

	// Read each line of the file and print it
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Check for errors that occurred during the scan
	if scanner.Err() != nil {
		log.Fatal(err)
	}
}
