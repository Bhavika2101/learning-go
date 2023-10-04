package main

import (
	"testing"
	"os"
	"io/ioutil"
	"strings"
)

// TODO: Replace with actual implementation of generatePassword
// func generatePassword(length int) string {
//	return "password"
// }

func TestGeneratePassword(t *testing.T) {
	length := 10
	password := generatePassword(length)

	if len(password) != length {
		t.Errorf("Expected password of length %d, but got %d", length, len(password))
	}

	// test uniqueness by generating another password
	anotherPassword := generatePassword(length)
	if password == anotherPassword {
		t.Errorf("Expected a unique password, but got two identical ones")
	}
}

func TestMainOutput(t *testing.T) {
	// Capture the output of the main function
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	// Check if the output is a string of length 10
	if len(strings.TrimSpace(string(out))) != 10 {
		t.Errorf("main() should print a string of length 10")
	}
}
