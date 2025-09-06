package main

import "testing"

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		expected string
	}{
		{
			name:     "remove scheme",
			inputURL: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove trailing slash",
			inputURL: "blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "no change",
			inputURL: "blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "uppercase to lowercase",
			inputURL: "HTTPS://BLOG.BOOT.DEV/PATH",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "http scheme and trailing slash",
			inputURL: "http://example.com/",
			expected: "example.com",
		},
		{
			name:     "only domain with http scheme",
			inputURL: "http://example.com",
			expected: "example.com",
		},
		{
			name:     "only domain with https scheme and trailing slash",
			inputURL: "https://example.com/",
			expected: "example.com",
		},
		{
			name:     "only domain with https scheme",
			inputURL: "https://example.com",
			expected: "example.com",
		},
		{
			name:     "IP address with https scheme and trailing slash",
			inputURL: "https://192.168.1.1/",
			expected: "192.168.1.1",
		},
		{
			name:     "IP address with http scheme",
			inputURL: "http://192.168.1.1",
			expected: "192.168.1.1",
		},
		{
			name:     "IP address with https scheme",
			inputURL: "https://192.168.1.1",
			expected: "192.168.1.1",
		},
		// add more test cases here
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}

/* TEST CASES FOR FRAGMENT, QUERY AND PORT VERSION
{
	name:     "remove scheme",
	inputURL: "https://blog.boot.dev/path",
	expected: "blog.boot.dev/path",
},
{
	name:     "remove trailing slash",
	inputURL: "blog.boot.dev/path/",
	expected: "blog.boot.dev/path",
},
{
	name:     "no change",
	inputURL: "blog.boot.dev/path",
	expected: "blog.boot.dev/path",
},
{
	name:     "uppercase to lowercase",
	inputURL: "HTTPS://BLOG.BOOT.DEV/PATH",
	expected: "blog.boot.dev/PATH",
},
{
	name:     "http scheme and trailing slash",
	inputURL: "http://example.com/",
	expected: "example.com",
},
{
	name:     "only domain with http scheme",
	inputURL: "http://example.com",
	expected: "example.com",
},
{
	name:     "only domain with https scheme and trailing slash",
	inputURL: "https://example.com/",
	expected: "example.com",
},
{
	name:     "only domain with https scheme",
	inputURL: "https://example.com",
	expected: "example.com",
},
{
	name:     "domain with path and query",
	inputURL: "https://example.com/path?query=1",
	expected: "example.com/path?query=1",
},
{
	name:     "domain with path, query, and trailing slash",
	inputURL: "https://example.com/path/?query=1",
	expected: "example.com/path?query=1",
},
{
	name:     "domain with path, query, and fragment",
	inputURL: "https://example.com/path/?query=1#section",
	expected: "example.com/path?query=1#section",
},
{
	name:     "domain with path, query, fragment, and trailing slash",
	inputURL: "https://example.com/path/?query=1#section/",
	expected: "example.com/path?query=1#section/",
},
{
	name:     "domain with port and path",
	inputURL: "http://example.com:8080/path",
	expected: "example.com:8080/path",
},
{
	name:     "domain with port, path, and trailing slash",
	inputURL: "http://example.com:8080/path/",
	expected: "example.com:8080/path",
},
{
	name:     "IP address with https scheme and trailing slash",
	inputURL: "https://192.168.1.1/",
	expected: "192.168.1.1",
},
{
	name:     "IP address with http scheme",
	inputURL: "http://192.168.1.1",
	expected: "192.168.1.1",
},
{
	name:     "IP address with https scheme",
	inputURL: "https://192.168.1.1",
	expected: "192.168.1.1",
},
*/
