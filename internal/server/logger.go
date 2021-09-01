package server

import (
	"fmt"
	"log"
	"os"
)

type Logger struct {
	logFile string
}

func (l Logger) Print(v ...interface{}) {
	f, err := os.OpenFile(l.logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(fmt.Sprint(v...)); err != nil {
		log.Println(err)
	}
}

func (l Logger) Info(v ...interface{}) {
	l.Print("(II) ", fmt.Sprintln(v...))
}

func (l Logger) Warn(v ...interface{}) {
	l.Print("(WW) ", fmt.Sprintln(v...))
}
