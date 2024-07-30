package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/brunobmello25/educationalsp/src/rpc"
)

func main() {
	logger := getLogger("/Users/bruno.mello/dev/personal/educationalsp/log.txt")
	logger.Println("Started")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Text()
		handleMessage(logger, msg)
	}
}

func handleMessage(logger *log.Logger, msg any) {
	logger.Println(msg)
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic(fmt.Sprintf("error opening log file:\n%s", err))
	}

	return log.New(logfile, "[educationalsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
