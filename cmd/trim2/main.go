package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

// TrimLeft is deleting space at the left side of a string.
func TrimLeft(out func(string)) (in func(string)) {
	in = func(s string) {
		t := strings.TrimLeft(s, " \t\r\n")
		out(t)
	}
	return
}

// TrimRight is removing space at the right side of a string.
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

// StringFilter is a 'filter' for strings.
type StringFilter func(out func(string)) (in func(string))

// Trim is deleting space at both sides of a string.
func Trim(out func(string), f StringFilter) (in func(string)) {
	trIn := TrimRight(out)
	fIn := f(trIn)
	in = TrimLeft(fIn)
	return
}

// LogString logs a string.
func LogString(out func(string)) (in func(string)) {
	in = func(s string) {
		log.Println(s)
		out(s)
	}
	return
}

func main() {
	in := Trim(
		func(s string) { fmt.Printf("Trim: `%s`\n", s) },
		LogString,
	)
	in(" abc ")
	in(" a ")
}
