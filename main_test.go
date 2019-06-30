package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestMainProgram(t *testing.T) {
	input := "1+2"
	os.Args = []string{"main", input}
	result := captureStdout(main)
	actual := "3"
	if result != actual {
		t.Errorf("Parsing of %s was incorrect, got: %s, want: %s.", input, result, actual)
	}
}

// Source https://gist.github.com/mindscratch/0faa78bd3c0005d080bf
func captureStdout(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return strings.TrimSuffix(buf.String(), "\n")
}
