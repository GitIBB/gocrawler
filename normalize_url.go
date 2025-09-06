package main

import (
	"net/url"
	"strings"
)

func normalizeURL(rawURL string) (string, error) {
	pURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	fullPath := pURL.Host + pURL.Path

	fullPath = strings.ToLower(fullPath)

	fullPath = strings.TrimSuffix(fullPath, "/")

	return fullPath, nil

	/* Version which includes port, query, and fragment
	scheme := strings.ToLower(pURL.Scheme)
	host := strings.ToLower(pURL.Hostname())
	path := pURL.Path
	port := pURL.Port()
	query := pURL.RawQuery
	frag := pURL.Fragment

	switch {
	case port == "":
		port = ""
	case (scheme == "http" && port == "80") || (scheme == "https" && port == "443"):
		port = ""
	default:
		port = ":" + port
	}

	for len(path) > 1 && strings.HasSuffix(path, "/") {
		path = strings.TrimSuffix(path, "/")
	}
	if path == "/" {
		path = ""
	}
	if frag != "" {
		frag = "#" + frag
	}
	if query != "" {
		path += "?" + query
	}
	suffix := path + frag

	return host + port + suffix, nil
	*/
}
