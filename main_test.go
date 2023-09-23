package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestReverseWithNoArguments(t *testing.T) {
	args := []string{}
	expectedOutput := "" // Replace with the expected output for this case

	output := captureOutput(func() {
		reverse(args)
	})

	if output != expectedOutput {
		t.Errorf("Expected: %s, Got: %s", expectedOutput, output)
	}
}

func TestReverseWithTooManyArguments(t *testing.T) {
	args := []string{"arg1", "arg2"}
	expectedOutput := "Too many arguments"

	output := captureOutput(func() {
		reverse(args)
	})

	if output != expectedOutput {
		t.Errorf("Expected: %s, Got: %s", expectedOutput, output)
	}
}

func TestReverseWithValidFile(t *testing.T) {
	args := []string{"--reverse=valid.txt"} // Replace with the name of a valid test file
	expectedOutput := "YourExpectedOutputHere" // Replace with the expected output for this case

	output := captureOutput(func() {
		reverse(args)
	})

	if output != expectedOutput {
		t.Errorf("Expected: %s, Got: %s", expectedOutput, output)
	}
}

func TestReverseWithInvalidFile(t *testing.T) {
	args := []string{"--reverse=invalid.txt"} // Replace with the name of a non-existent test file
	expectedOutput := "Could not read the content in the file due to " // This is the expected prefix of the error message

	output := captureOutput(func() {
		reverse(args)
	})

	if !strings.HasPrefix(output, expectedOutput) {
		t.Errorf("Expected output to start with: %s, Got: %s", expectedOutput, output)
	}
}

// Add more test functions for different scenarios

func captureOutput(f func()) string {
	// Redirect standard output to capture printed output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	// Reset standard output
	w.Close()
	os.Stdout = old

	// Read the captured output
	var buf bytes.Buffer
	io.Copy(&buf, r)

	return buf.String()
}
