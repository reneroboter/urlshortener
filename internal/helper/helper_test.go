package helper

import (
	"testing"
)

func Test_StringIsValidSHA1String(t *testing.T) {
	tests := map[string]struct {
		sha1_string string
		expected    bool
	}{
		"hashed 'asd123?#' sha1 string": {
			sha1_string: "12455cf13af0a7338400a7ed255eac6a7eb4fe9b",
			expected:    true,
		},
		"less than 40 char  'asd123?#' sha1 string": {
			sha1_string: "455cf13af0a7338400a7ed255eac6a7eb4fe9b",
			expected:    false,
		},
		"broken more than 40 char 'asd123?#' sha1 string": {
			sha1_string: "455cf13af0a7338400a7ed255eac6a7eb4fe9b455cf13af0a7338400a7ed255eac6a7eb4fe9b",
			expected:    false,
		},
		"hashed 'https://www.google' sha1 string": {
			sha1_string: "f148c8a0a0900a3f0f154eb891e4229e228f1dd4",
			expected:    true,
		},
		"less than 40 char 'https://www.google' sha1 string": {
			sha1_string: "f148c8a0900a3f0f154eb891e4229e228f1dd4",
			expected:    false,
		},
		"uppercase 'FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF' sha1 string": {
			sha1_string: "FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF",
			expected:    true,
		},
		"uppercase 'ABCDEF1234567890ABCDEF1234567890ABCDEF12' sha1 string": {
			sha1_string: "ABCDEF1234567890ABCDEF1234567890ABCDEF12",
			expected:    true,
		},
		"whitespace '' string": {
			sha1_string: "",
			expected:    false,
		},
		"unicode '測試12345678901234567890123456789012345678' string": {
			sha1_string: "測試12345678901234567890123456789012345678",
			expected:    false,
		},
		"unicode 'μ10e2821bbbea527ea02200352313bc059445190' string": {
			sha1_string: "μ10e2821bbbea527ea02200352313bc059445190",
			expected:    false,
		},
		"non-hex ASCII 'g10e2821bbbea527ea02200352313bc059445190' string": {
			sha1_string: "g10e2821bbbea527ea02200352313bc059445190",
			expected:    false,
		},
		"non-hex ASCII '12455cf13af0a7338400a7ed255eac6a7eb4fe9!' string": {
			sha1_string: "12455cf13af0a7338400a7ed255eac6a7eb4fe9!",
			expected:    false,
		},
	}
	for name, test := range tests {
		name := name
		test := test

		t.Run(name, func(t *testing.T) {
			t.Parallel() // speed up tests
			result := IsValidSHA1(test.sha1_string)
			if result != test.expected {
				t.Errorf("expected '%t', got '%t'", test.expected, result)
			}
		})
	}
}

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

func Test_StringIsCorrectUrl(t *testing.T) {
	t.Skip()
}
