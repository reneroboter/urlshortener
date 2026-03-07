package application

import "testing"

func Test_StringIsCorrectlyHashed(t *testing.T) {
	tests := map[string]struct {
		url      string
		expected string
	}{
		"valid url https://www.google": {
			url:      "https://www.google",
			expected: "f148c8a0a0900a3f0f154eb891e4229e228f1dd4",
		},
		"empty string": {
			url:      "",
			expected: "da39a3ee5e6b4b0d3255bfef95601890afd80709",
		},
		"one char difference: abc": {
			url:      "abc",
			expected: "a9993e364706816aba3e25717850c26c9cd0d89d",
		},
		"one char difference: abcd": {
			url:      "abcd",
			expected: "81fe8bfe87576c3ecb22426f8e57847382917acf",
		},
		"with unicode: https://www.google/😀/äüö": {
			url:      "https://www.google/😀/äüö",
			expected: "686147b483eb023e07b9500b7d2687a1bceede1f",
		},
		"with unicode: https://例子.测试": {
			url:      "https://例子.测试",
			expected: "ff41d6e9c83d931da7bb42cd02591da31cdbba9a",
		},
		"same url, but different cases 1: https://example.com": {
			url:      "https://example.com",
			expected: "327c3fda87ce286848a574982ddd0b7c7487f816",
		},
		// todo need to be enabled later -> missing normalization layer
		//"same url, but different cases 2: https://example.com": {
		//	url:      "https://EXAMPLE.COM",
		//		expected: "f148c8a0a0900a3f0f154eb891e4229e228f1dd4",
		//},
		//"same url, but different cases 3: https://example.com": {
		//	url:      "https://example.com/",
		//	expected: "f148c8a0a0900a3f0f154eb891e4229e228f1dd4",
		//}//,

	}
	for name, test := range tests {
		name := name
		test := test

		t.Run(name, func(t *testing.T) {
			t.Parallel() // speed up tests
			result := HashUrl(test.url)
			if result != test.expected {
				t.Errorf("expected '%s', got '%s'", test.expected, result)
			}
		})
	}
}
