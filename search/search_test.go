package search

import (
	"regexp"
	"testing"
)

func TestSearchName(t *testing.T) {
	name := "Alex"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Search("Alex")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Search("Alex") = %q,%v, want match fir %#q,nil`, msg, err, want)
	}
}

func TestSearchEmpty(t *testing.T) {
	msg, err := Search("")
	if msg != "" || err == nil {
		t.Fatalf(`Search("") = %q, %v, want "", error`, msg, err)
	}
}
