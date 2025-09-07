package main

import (
	"net/url"
	"strings"

	"fmt"
)

func (cfg *config) crawlPage(rawCurrentURL string) {

	cfg.concurrencyControl <- struct{}{} // acquire a slot
	defer func() {
		<-cfg.concurrencyControl // release the slot
		cfg.wg.Done()
	}()

	if cfg.pagesLen() >= cfg.maxPages {
		return
	}

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		logError("parsing current URL %s: %v", rawCurrentURL, err)
		return
	}

	if cfg.baseURL.Host != currentURL.Host {
		return
	}
	lowerPath := strings.ToLower(currentURL.Path)

	normCurrent, err := normalizeURL(rawCurrentURL)
	if err != nil {
		logError("error normalizing current URL %s: %v", rawCurrentURL, err)
		return
	}

	isFirst := cfg.addPageVisits(normCurrent)
	if !isFirst {
		// already visited this page or hit max page limit
		return
	}
	switch { // skip URLs likely to be non-HTML resources
	case
		strings.HasSuffix(lowerPath, ".xml"),
		strings.HasSuffix(lowerPath, ".pdf"),
		strings.HasSuffix(lowerPath, ".jpg"),
		strings.HasSuffix(lowerPath, ".jpeg"),
		strings.HasSuffix(lowerPath, ".png"),
		strings.HasSuffix(lowerPath, ".gif"),
		strings.HasSuffix(lowerPath, ".svg"),
		strings.HasSuffix(lowerPath, ".ico"),
		strings.HasSuffix(lowerPath, ".mp3"),
		strings.HasSuffix(lowerPath, ".mp4"),
		strings.HasSuffix(lowerPath, ".atom"),
		strings.HasSuffix(lowerPath, ".zip"),
		strings.HasSuffix(lowerPath, ".css"),
		strings.HasSuffix(lowerPath, ".json"):
		fmt.Printf("non-HTML resource, halted crawling %s\n", rawCurrentURL)
		return
	}

	// fetch the HTML for the page
	fmt.Printf("crawling %s\n", rawCurrentURL) // current url being crawled
	currentHTML, err := getHTML(rawCurrentURL)
	if err != nil {
		logError("error getting HTML for URL %s: %v", rawCurrentURL, err)
		return
	}
	fmt.Printf("fetched %d bytes of HTML\n", len(currentHTML))

	// get all the URLs from the response body HTML
	nextURLs, err := getURLsFromHTML(currentHTML, rawCurrentURL)
	if err != nil {
		logError("error getting URLs from HTML for URL %s: %v", rawCurrentURL, err)
		return
	}
	// recursively crawl each URL on the page
	for _, nURL := range nextURLs {
		cfg.wg.Add(1)
		go cfg.crawlPage(nURL)
	}
}
