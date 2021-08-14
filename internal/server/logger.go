package server

import "log"

type Logger struct{}

func (l Logger) Println(v ...interface{}) {
	log.Println(v)
}

func (l Logger) Info(v ...interface{}) {
	l.Println("(II)", v)
}

func (l Logger) Warning(v ...interface{}) {
	l.Println("(WW)", v)
}

