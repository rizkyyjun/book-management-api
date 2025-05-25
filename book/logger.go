package book

import (
	"log"
	"os"
)

var logChan = make(chan string, 100)

func StartLogger() {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal("failed to initialize log file", err)
	}
	file.Close()

	go func() {
		f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("failed to open log file: %v", err)
		}
		defer f.Close()

		for msg := range logChan {
			f.WriteString(msg + "\n")
		}
	}()
}

func Log(msg string) {
	logChan <- msg
}