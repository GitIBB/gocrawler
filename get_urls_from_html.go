package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	baseURL, err := url.Parse(rawBaseURL) // parses the base URL string into a URL object
	if err != nil {
		return nil, fmt.Errorf("couldn't parse base URL: %v", err)
	}
	htmlReader := strings.NewReader(htmlBody) // creates an io.Reader from the html string
	NodeTree, err := html.Parse(htmlReader)   // creates a tree of html.Node(s) from the io.Reader
	if err != nil {                           // if parsing fails, return the error
		return nil, fmt.Errorf("couldn't parse HTML: %v", err)
	}

	var urls []string
	var traverseNodes func(*html.Node)      // declares variable function f that will hold a function which takes a HTML node pointer
	traverseNodes = func(node *html.Node) { // assigns the function to f, allowing recursion
		if node.Type == html.ElementNode && node.Data == "a" { // checks if the current node is a HTML element AND an <a> anchor tag
			for _, attr := range node.Attr { // loops through all the attributes of the <a> tag
				if attr.Key == "href" {
					// Parse the href URL
					hrefURL, err := url.Parse(attr.Val)
					if err != nil {
						continue // Skip invalid URLs
					}
					// Resolve relative URLs against the base URL
					absoluteURL := baseURL.ResolveReference(hrefURL)
					urls = append(urls, absoluteURL.String())
					break
				}
			}
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			traverseNodes(child) // recursively calls traverseNodes on each child node of the current node
		}
	}
	traverseNodes(NodeTree) // initiates the recursive traversal starting from the root node

	return urls, nil

}
