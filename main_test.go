package main

import "testing"

func Test_calculate(t *testing.T) {
	testCases := []struct {
		filepath string
		expected int64
	}{
		{"./test/1.txt", 1},
		{"./test/2.txt", 3},
		{"./test/3.txt", 2},
	}
	for i, tt := range testCases {
		result, err := calculate(tt.filepath)
		if err != nil {
			t.Errorf("unexpected error: %s\n", err)
		}
		if result != tt.expected {
			t.Errorf("%d failed for %s: expected %d got %d\n", i, tt.filepath, tt.expected,
				result)
		}
	}
}
