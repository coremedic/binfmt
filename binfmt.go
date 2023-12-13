package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/atotto/clipboard"
)

var format string

func init() {
	flag.StringVar(&format, "f", "c", "format of the output (c, go, py)")
}

func main() {
	// Parse format flag
	flag.Parse()

	// Check for correct usage
	if len(flag.Args()) != 1 {
		log.Fatalf("Usage: %s [-f format] <binary_file>", os.Args[0])
	}

	// Read binary file
	bin := flag.Arg(0)
	binBytes, err := os.ReadFile(bin)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	// Allocate buffer
	var buf bytes.Buffer

	// Generate output based on reuqested format
	switch format {
	case "c": // Format in C/C++ mode
		// C/C++ format as unsigned character array
		_, err := buf.WriteString("unsigned char payload[] = {")
		if err != nil {
			log.Fatalf("Failed to write to buffer: %v", err)
		}
		for i, b := range binBytes {
			// Add new line every 12 bytes
			if i%12 == 0 {
				_, err := buf.WriteString("\n    ")
				if err != nil {
					log.Fatalf("Failed to write to buffer: %v", err)
				}
			}
			// Write byte to buffer in hex format
			_, err := buf.WriteString(fmt.Sprintf("0x%02x, ", b))
			if err != nil {
				log.Fatalf("Failed to write to buffer: %v", err)
			}
		}
		// Close of char array on new line
		buf.WriteString("\n}")
	default: // Invalid format
		log.Fatalf("Invalid format specified. Use 'c', 'go', or 'py'.")
	}

	// Write formatted output to clipboard
	outString := buf.String()
	if err := clipboard.WriteAll(outString); err != nil {
		log.Fatalf("Failed to copy to clipboard: %s", err)
	}

	fmt.Println(outString)
	fmt.Println("Formatted binary copied to clipboard")
}
