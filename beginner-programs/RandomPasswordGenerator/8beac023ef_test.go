// Test generated by RoostGPT for test golang-level0-PassGen using AI Type Open AI and AI Model gpt-4-1106-preview

package main

import (
	"testing"
	"unicode"
	"unicode/ascii"
)

func TestGeneratePasswordPositive(t *testing.T) {
	length := 10
	password := GeneratePassword(length)
	
	if len(password) != length {
		t.Errorf("Expected password length of %d, but got %d", length, len(password))
	}
	
	digits := 0
	specials := 0
	letters := 0
	
	for _, r := range password {
		switch {
		case unicode.IsDigit(r):
			digits++
		case unicode.IsLetter(r) && ascii.IsPrint(r):
			letters++
		case unicode.IsPunct(r) || unicode.IsSymbol(r):
			specials++
		default:
			t.Errorf("Invalid character found: %r", r)
		}
	}
	
	if digits == 0 || specials == 0 || letters == 0 {
		t.Errorf("Password does not contain at least one digit, one special character, and one letter")
	}
}

func TestGeneratePasswordNegative(t *testing.T) {
	length := -1
	password := GeneratePassword(length)
	
	if password != "" {
		t.Errorf("Expected empty password for negative length, got %q", password)
	}
}

