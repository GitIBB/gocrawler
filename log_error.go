package main

import "fmt"

func logError(format string, args ...interface{}) {
	fmt.Printf("ERROR: "+format+"\n", args...)
}
