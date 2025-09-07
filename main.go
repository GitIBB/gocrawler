package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("No .env file found, proceeding with environment variables")
	}
	if len(os.Args) < 4 {
		fmt.Println("not enough arguments provided")
		fmt.Println("usage: crawler <base_url> [maxConcurrency] [maxPages]")
		os.Exit(1)
	}
	if len(os.Args) > 4 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	rawBaseURL := os.Args[1]
	fmt.Printf("\n=============================\n Starting crawl of %s \n=============================\n", rawBaseURL)

	maxConcurrencyString := os.Args[2]
	maxPagesString := os.Args[3]

	maxConcurrency, err := strconv.Atoi(maxConcurrencyString)
	if err != nil || maxConcurrency < 1 {
		fmt.Printf("invalid maxConcurrency value: %s\n", maxConcurrencyString)
		os.Exit(1)
	}

	maxPages, err := strconv.Atoi(maxPagesString)
	if err != nil || maxPages < 1 {
		fmt.Printf("invalid maxPages value: %s\n", maxPagesString)
		os.Exit(1)
	}

	cfg, err := configure(rawBaseURL, maxConcurrency, maxPages)
	if err != nil {
		fmt.Printf("error configuring crawler: %v\n", err)
		os.Exit(1)
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait() // wait for all goroutines to finish

	sortedPages := sortPagesByCount(cfg.pages)
	printReport(sortedPages, rawBaseURL)

	fmt.Printf("\n=============================\ncrawl complete, %d pages found\n=============================\n", len(cfg.pages))
}
