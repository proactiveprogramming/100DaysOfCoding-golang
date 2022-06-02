/ File: greetingstest/greetingstest_test.go
// Version: 0.1
// Note: the "_test" part of the filename indicates that this file is a test
//       file.  The "greetingstest" is just our project name.
//       If this was a real project, the file would most likely be named "greetings_test.go"
//       With the corresponding library file called "greetings.go".
// -------------------------------------------------------
package greetingstest

import (
	"testing"
	"regexp"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	name := "Gladys"
	// \b are word boundaries.
	want := regexp.MustCompile(`\b`+name+`\b`)
	msg, err := Hello("Gladys")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
