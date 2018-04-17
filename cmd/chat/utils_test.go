package main

import (
	"testing"
)

var tests = []struct {
	s      []string
	search string
	want   bool
}{
	{[]string{"neil", "matt", "adam"}, "neil", true},
	{[]string{"neil", "matt", "adam"}, "neil1", false},
	{[]string{"neil", "matt", "adam"}, "adam", true},
}

func TestStringInSlice(t *testing.T) {
	for _, c := range tests {
		got, err := StringInSlice(c.s, c.search)
		Ok(t, err)
		Equals(t, c.want, got)
	}
}
