package main

import (
	"bytes"
	"os"
	"testing"
)

func SayHelloTo(name string) {
	fmt.Printf("Hello %s\n", name)
}

func TestSayHelloTo_4a6fd0c56c(t *testing.T) {
	cases := []struct {
		name     string
		expected string
	}{
		{
			name:     "John",
			expected: "Hello John\n",
		},
		{
			name:     "",
			expected: "Hello \n",
		},
	}

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	for _, c := range cases {
		SayHelloTo(c.name)

		w.Close()
		var buf bytes.Buffer
		os.Stdout = old
		buf.ReadFrom(r)
		result := buf.String()

		if result != c.expected {
			t.Errorf("SayHelloTo(%q) == %q, expected %q", c.name, result, c.expected)
		} else {
			t.Logf("SayHelloTo(%q) == %q, passed", c.name, result)
		}

		r, w, _ = os.Pipe()
		os.Stdout = w
	}
}
