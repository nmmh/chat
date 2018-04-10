package main

import (
	"testing"

	"github.com/nmmh/chat/utils"
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

func TestUsernameInUse(t *testing.T) {
	cm := NewCM()
	for _, c := range tests {
		got, err := cm.UsernameInUse(c.s, c.search)
		utils.Ok(t, err)
		utils.Equals(t, c.want, got)
	}
}

func TestFormatUserList(t *testing.T) {
	cm := NewCM()
	var tests1 = []struct {
		s    []string
		want string
	}{
		{[]string{"neil", "matt", "adam"}, "UserList:{adam, matt, neil} Total:[3]"},
		{[]string{"neil", "matt", "adam", "linda"}, "UserList:{adam, linda, matt, neil} Total:[4]"},
		{[]string{}, "UserList:{} Total:[0]"},
	}
	for _, c := range tests1 {
		got, err := cm.FormatUserList(c.s)
		utils.Ok(t, err)
		utils.Equals(t, c.want, got)
	}
}
