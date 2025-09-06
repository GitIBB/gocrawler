package main

import (
	"fmt"
	"os"
)

func main() {

	urlArg := os.Args[1:]
	if len(urlArg) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(urlArg) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	if len(urlArg) == 1 {
		fmt.Printf("starting crawl of %s\n", urlArg[0])
	}

	fmt.Println("Hello, World!")
}
