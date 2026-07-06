package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("message.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	currentLine := ""
	for {
		data := make([]byte, 8)
		n, err := f.Read(data)
		if err != nil {
			break
		}

		parts := strings.Split(string(data[:n]), "\n")
		for i := 0; i < len(parts); i++ {
			currentLine += parts[i]
			if i < len(parts)-1 {
				fmt.Printf("read: %s\n", currentLine)
				currentLine = ""
			}
		}
	}

	if currentLine != "" {
		fmt.Printf("read: %s\n", currentLine)
	}
}
