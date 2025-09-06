package main

import "testing"

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
			<html>
				<body>
					<a href="/path/one">Link</a>
						<span>Boot.dev</span>
					<a/>
					<a href="https://blog.boot.dev/path/two">Link</a>
					<a/>
				</body>
			</html>
			`,
			expected: []string{"https://blog.boot.dev/path/one", "https://blog.boot.dev/path/two"},
		},
		{
			name:     "no links",
			inputURL: "https://blog.boot.dev",
			inputBody: `
			<html>
				<body>
					<span>Boot.dev</span>
				</body>
			</html>
			`,
			expected: []string{},
		},
		{
			name:     "empty body",
			inputURL: "https://blog.boot.dev",
			inputBody: `
			`,
			expected: []string{},
		},
		{
			name:     "malformed html",
			inputURL: "https://blog.boot.dev",
			inputBody: `
			<html>
				<body>
					<a href="/path/one">Link
						<span>Boot.dev</span>
					<a/>
					<a href="https://blog.boot.dev/path/two">Link</a>
					<a/>
				</body>
			</html>
			`,
			expected: []string{"https://blog.boot.dev/path/one", "https://blog.boot.dev/path/two"},
		},
		// add more test cases here
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if len(actual) != len(tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
