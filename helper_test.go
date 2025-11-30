package main

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
			result := isValidSHA1(test.sha1_string)
			if result != test.expected {
				t.Errorf("expected '%t', got '%t'", test.expected, result)
			}
		})
	}
}

func Test_StringIsCorrectlyHashed(t *testing.T) {
	t.Skip()
}

func Test_StringIsCorrectUrl(t *testing.T) {
	t.Skip()
}
