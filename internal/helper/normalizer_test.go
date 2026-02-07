package helper

import (
	"testing"
)

func Test_Normalizer(t *testing.T) {
	tests := map[string]struct {
		url      string
		expected string
	}{
		"correct url stays correct": {
			url:      "https://example.com",
			expected: "https://example.com",
		},
		"lower case host": {
			url:      "https://EXAMPLE.COM",
			expected: "https://example.com",
		},
		"trims trailing slash": {
			url:      "https://example.com/",
			expected: "https://example.com",
		},
		"trims space left": {
			url:      "  https://example.com",
			expected: "https://example.com",
		},
		"trims space right": {
			url:      "https://example.com  ",
			expected: "https://example.com",
		},
		"trims new line": {
			url:      "https://example.com\n",
			expected: "https://example.com",
		},
		"keeps port": {
			url:      "https://example.com:80",
			expected: "https://example.com:80",
		},
		"lowercases host only": {
			url:      "https://exAMPle.com/pAge?x=1#SeCtIon-2",
			expected: "https://example.com/pAge?x=1#SeCtIon-2",
		},
		"keeps fragment": {
			url:      "https://example.com#Top",
			expected: "https://example.com#Top",
		},
		"keeps order'": {
			url:      "https://example.com?b=2&a=1",
			expected: "https://example.com?b=2&a=1",
		},
		"invalid url is returned as-is": {
			url:      "://not-a-url",
			expected: "://not-a-url",
		},
	}
	for name, test := range tests {
		name := name
		test := test

		t.Run(name, func(t *testing.T) {
			t.Parallel()
			result := NormalizeUrl(test.url)
			if result != test.expected {
				t.Errorf("expected '%s', got '%s'", test.expected, result)
			}
		})
	}
}
