package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_updateMessage(t *testing.T) {
	wg.Add(1)

	go updateMessage("Hello")

	wg.Wait()

	if msg != "Hello" {
		t.Error("Exepcted to find Hello, but not found.")
	}
}

func Test_printMessage(t *testing.T) {
	stdOut := os.Stdout

	read, write, _ := os.Pipe()
	os.Stdout = write

	msg = "hello"
	printMessage()

	_ = write.Close()

	result, _ := io.ReadAll(read)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "hello") {
		t.Error("Expected hello, but not found")
	}

}

func Test_main(t *testing.T) {
	stdOut := os.Stdout

	read, write, _ := os.Pipe()
	os.Stdout = write

	main()

	_ = write.Close()

	result, _ := io.ReadAll(read)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "Update One") {
		t.Error("Expected Update One, but not found")
	}

	if !strings.Contains(output, "Update Two") {
		t.Error("Expected Update Two, but not found")
	}

	if !strings.Contains(output, "Update Three") {
		t.Error("Expected Update Three, but not found")
	}
}
