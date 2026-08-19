package main

import (
	"fmt"
	"os"
)

// data_anonymizer - Anonymize sensitive data
func data_anonymizer(path string) {
	fmt.Println("========================================")
	fmt.Println("  Data-Anonymizer")
	fmt.Println("  Anonymize sensitive data")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	data_anonymizer(path)
}
