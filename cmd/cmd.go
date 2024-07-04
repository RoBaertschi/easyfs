package main

import (
	"fmt"
	"log"
	"os"
)

func NewLogger(name string) *log.Logger {
	return log.New(os.Stdout, fmt.Sprintf("[%v] ", name), log.Lshortfile|log.Ltime|log.Ldate)
}

var mainLogger = NewLogger("main")

func main() {
	mainLogger.Println("Hello World")
}
