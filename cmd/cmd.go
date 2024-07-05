package main

import (
	"fmt"
	"log"
	"os"

	"robaertshi.xyz/easyfs/config"
	"robaertshi.xyz/easyfs/server"
)

func NewLogger(name string) *log.Logger {
	return log.New(os.Stdout, fmt.Sprintf("[%v] ", name), log.Lshortfile|log.Ltime|log.Ldate)
}

var mainLogger = NewLogger("main")

func main() {
	mainLogger.Println("Hello World")

	conf, err := config.ReadConfig("./easyfs.toml")
	if err != nil {
		mainLogger.Fatalf("Could not read config, error: %s", err.Error())
	}

	err = server.StartServer(conf, NewLogger("server"))
	if err != nil {
		mainLogger.Fatalf("Error while running server: %s", err.Error())
	}
	mainLogger.Println("Finished")
}
