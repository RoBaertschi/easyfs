package main

import (
	"fmt"
	"log"
	"os"

	"robaertshi.xyz/easyfs/config"
)

func NewLogger(name string) *log.Logger {
	return log.New(os.Stdout, fmt.Sprintf("[%v] ", name), log.Lshortfile|log.Ltime|log.Ldate)
}

var mainLogger = NewLogger("main")

func main() {
	mainLogger.Println("Hello World")

	_, err := config.ReadConfig("./easyfs.toml")
	if err != nil {
		mainLogger.Fatalf("Could not read config, error: %s", err.Error())
	}

}
