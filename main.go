package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(os.Args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	rawBaseURL := os.Args[1]
	fmt.Printf("starting crawl of %s\n", rawBaseURL)

	res, err := getHtml(rawBaseURL)
	if err != nil {
		fmt.Println("Error fetching HTML:", err)
		os.Exit(1)
	}

	fmt.Println(res)
}
