package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	out := make(chan string, 1)

	go func() {
		defer f.Close()
		defer close(out)

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
					out <- currentLine
					currentLine = ""
				}
			}
		}

		if currentLine != "" {
			out <- currentLine
		}
	}()

	return out
}

func main() {
	f, err := os.Open("message.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := getLinesChannel(f)
	for line := range lines {
		fmt.Printf("read: %s\n", line)
	}
}
