// Test generated by RoostGPT for test go-unit-algo-string using AI Type Open AI and AI Model gpt-4

package Fibonacci

import (
	"testing"
	"reflect"
)

func TestFibonacciSequence_bf4aa71a9c(t *testing.T) {
	// Test case 1: Positive number
	num1 := 5
	expectedResult1 := []int{0, 1, 1, 2, 3, 5}
	result1 := fibonacciSequence(num1)
	if !reflect.DeepEqual(result1, expectedResult1) {
		t.Errorf("Test Case 1 Failed: %v inputted and expected output %v, but got %v", num1, expectedResult1, result1)
	} else {
		t.Logf("Test Case 1 Passed: %v inputted and expected output %v", num1, expectedResult1)
	}

	// Test case 2: Zero
	num2 := 0
	expectedResult2 := []int{0}
	result2 := fibonacciSequence(num2)
	if !reflect.DeepEqual(result2, expectedResult2) {
		t.Errorf("Test Case 2 Failed: %v inputted and expected output %v, but got %v", num2, expectedResult2, result2)
	} else {
		t.Logf("Test Case 2 Passed: %v inputted and expected output %v", num2, expectedResult2)
	}

	// Test case 3: Negative number
	num3 := -1
	expectedResult3 := []int{}
	result3 := fibonacciSequence(num3)
	if !reflect.DeepEqual(result3, expectedResult3) {
		t.Errorf("Test Case 3 Failed: %v inputted and expected output %v, but got %v", num3, expectedResult3, result3)
	} else {
		t.Logf("Test Case 3 Passed: %v inputted and expected output %v", num3, expectedResult3)
	}
}