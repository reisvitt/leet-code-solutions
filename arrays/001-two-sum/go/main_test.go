package main

import "testing"

type TestCase struct {
	nums     []int
	target   int
	expected []int
}

var tests = []TestCase{
	{
		nums:     []int{2, 7, 11, 15},
		target:   9,
		expected: []int{0, 1},
	},
	{
		nums:     []int{3, 2, 4},
		target:   6,
		expected: []int{1, 2},
	},
	{
		nums:     []int{3, 3},
		target:   6,
		expected: []int{0, 1},
	},
}

func TestTwoSum(t *testing.T) {
	for _, test := range tests {
		result := twoSum(test.nums, test.target)

		if result == nil || !sliceEqual(result, test.expected) {
			t.Errorf("Expected %v, got %v for nums %v and target %v", test.expected, result, test.nums, test.target)
		}
	}
}

func sliceEqual(a, b []int) bool {
	return a[0] == b[0] && a[1] == b[1]
}
