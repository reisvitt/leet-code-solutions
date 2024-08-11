package main

import "testing"

type TestStruct struct {
	input    string
	expected bool
}

func TestValidPalindrome(t *testing.T) {
	tests := []TestStruct{
		{
			input:    "A man, a plan, a canal: Panama",
			expected: true,
		},
		{
			input:    "race a car",
			expected: false,
		},
		{
			input:    " ",
			expected: true,
		},
	}

	for _, test := range tests {
		result := isPalindrome(test.input)
		if result != test.expected {
			t.Errorf("Expected %v, got %v in word '%s'", test.expected, result, test.input)
		}
	}
}
