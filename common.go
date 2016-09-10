package log

import (
	golog "log"
)

// Print calls log.Print()
func Print(v ...interface{}) {
	golog.Print(v...)
}

// Printf calls log.Printf()
func Printf(format string, v ...interface{}) {
	golog.Printf(format, v...)
}

// Println calls log.Println()
func Println(v ...interface{}) {
	golog.Println(v...)
}
