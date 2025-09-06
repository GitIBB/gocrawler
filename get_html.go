package main

import (
	"fmt"
	"io"
	"net/http"
)

func getHtml(rawURL string) (string, error) {

	client := &http.Client{}

	request, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		return "", err
	}
	request.Header.Set("User-Agent", "BootCrawler/1.0")

	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	if response.Header.Get("Content-Type") != "text/html" {
		return "", fmt.Errorf("expected text/html but got %s", response.Header.Get("Content-Type"))
	}
	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch HTML: %s", response.Status)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
