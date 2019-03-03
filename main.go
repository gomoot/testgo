package main

import (
	"fmt"
	"log"
	"os"

	"github.com/natefinch/lumberjack"
	//"github.com/sirupsen/logrus"
)

var errLog *log.Logger

func main() {

	e, err := os.OpenFile("./tt.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Printf("error opening file: %v", err)
		os.Exit(1)
	}
	errLog = log.New(e, "", log.Ldate|log.Ltime)
	errLog.SetOutput(&lumberjack.Logger{
		Filename:   "./foo.log",
		MaxSize:    100,   // megabytes after which new file is created
		MaxBackups: 10000, // number of backups
		MaxAge:     28,    //days
	})

	for i := 0; i <= 100000000; i++ {
		if i % 4 == 0 {
		errLog.Printf("ddddddddddddddddddddddddddddddddTest Test\n")
		} else {
			errLog.Printf("ddddddddddddddddddddddddddiohdfjjjjjjjjjjjjjjjjjjjjjjjjjjjjj;lkdj kaddddddTest Test\n")
			
		}
	}
}
