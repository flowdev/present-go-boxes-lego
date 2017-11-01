package main

import (
	"errors"
	"fmt"
	"strings"
)

// TrimLeft is trimming space at the left side of a string.
func TrimLeft(out func(string)) (in func(string)) {
	in = func(s string) {
		t := strings.TrimLeft(s, " \t\r\n")
		out(t)
	}
	return
}

// TrimRight is trimming space at the right side of a string.
func TrimRight(out func(string)) (in func(string)) {
	in = func(s string) {
		t := strings.TrimRight(s, " \t\r\n")
		if len(t) < 2 {
			panic(errors.New("show us the stack trace"))
		}
		out(t)
	}
	return
}

// Trim is trimming space at both sides of a string.
func Trim(out func(string)) (in func(string)) {
	trIn := TrimRight(out)
	in = TrimLeft(trIn)
	return
}

func main() {
	in := Trim(func(s string) { fmt.Printf("Trim: `%s`\n", s) })
	in(" abc ")
	in(" a ")
}
