// Test generated by RoostGPT for test go-unit-algo-string using AI Type Open AI and AI Model gpt-4

package Factorial

import (
	"testing"
)

func TestFactorialRecursive_178a7b8974(t *testing.T) {
	// Test case 1: Testing with a positive number
	num := 5
	expectedResult := 120
	result := FactorialRecursive(num)
	if result != expectedResult {
		t.Errorf("Test case 1 failed, expected: %d, got: %d", expectedResult, result)
	} else {
		t.Logf("Test case 1 success, expected: %d, got: %d", expectedResult, result)
	}

	// Test case 2: Testing with 0
	num = 0
	expectedResult = 1
	result = FactorialRecursive(num)
	if result != expectedResult {
		t.Errorf("Test case 2 failed, expected: %d, got: %d", expectedResult, result)
	} else {
		t.Logf("Test case 2 success, expected: %d, got: %d", expectedResult, result)
	}

	// Test case 3: Testing with a negative number
	num = -5
	expectedResult = -120
	result = FactorialRecursive(num)
	if result != expectedResult {
		t.Errorf("Test case 3 failed, expected: %d, got: %d", expectedResult, result)
	} else {
		t.Logf("Test case 3 success, expected: %d, got: %d", expectedResult, result)
	}
}
