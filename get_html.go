package main

import (
	"fmt"
	"io"
	"mime"
	"net/http"
)

func getHTML(rawURL string) (string, error) {

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
	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch HTML: %s", response.Status)
	}

	ct := response.Header.Get("Content-Type")
	mt, _, err := mime.ParseMediaType(ct)
	if err != nil || mt != "text/html" {
		return "", fmt.Errorf("expected text/html but got %s: %v", mt, err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
