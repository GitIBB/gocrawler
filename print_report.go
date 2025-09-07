package main

import (
	"fmt"
	"net/url"
	"sort"
)

type PageCount struct {
	URL   string
	Count int
}

func printReport(sortedPages []PageCount, baseURL string) {
	fmt.Printf("\n=============================\n REPORT for %s \n=============================\n", baseURL)

	parsedBase, err := url.Parse(baseURL)
	if err != nil {
		fmt.Printf("Error parsing base URL: %v\n", err)
		return
	}
	scheme := parsedBase.Scheme

	// print one line per page
	for _, page := range sortedPages {

		normalizedURL, err := normalizeURL(page.URL) // normalize the URL but keep the scheme
		if err != nil {
			fmt.Printf("Found %d internal links to %s\n", page.Count, page.URL) // fallback
		} else {
			// add scheme back to normalized URL
			fullURL := scheme + "://" + normalizedURL
			fmt.Printf("Found %d internal links to %s\n", page.Count, fullURL)
		}
	}
}

func sortPagesByCount(pages map[string]int) []PageCount { // takes a map of URLs to visit counts, returns a sorted slice of PageCount structs
	pageSlice := make([]PageCount, 0, len(pages))
	for url, count := range pages {
		pageSlice = append(pageSlice, PageCount{URL: url, Count: count})
	}

	sort.Slice(pageSlice, func(i, j int) bool {
		if pageSlice[i].Count == pageSlice[j].Count { // sort by count (descending)
			return pageSlice[i].URL < pageSlice[j].URL // if counts are equal, sort by URL (ascending)
		}
		return pageSlice[i].Count > pageSlice[j].Count
	})

	return pageSlice
}
