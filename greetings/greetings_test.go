package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Hyunjae"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Hyunjae")

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Hyunjae") = %q, %v, want match for %#q, nil`, msg, err, want)
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
