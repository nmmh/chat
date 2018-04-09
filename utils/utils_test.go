package utils

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
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

//Test d
func TestStringInSlice(t *testing.T) {
	for _, c := range tests {
		got, err := stringInSlice(c.s, c.search)
		ok(t, err)
		equals(t, c.want, got)
	}
}
func TestUsernameInUse(t *testing.T) {
	for _, c := range tests {
		got, err := usernameInUse(c.s, c.search)
		ok(t, err)
		equals(t, c.want, got)
	}
}

func TestFormatUserList(t *testing.T) {
	var tests1 = []struct {
		s    []string
		want string
	}{
		{[]string{"neil", "matt", "adam"}, "UserList:{adam, matt, neil} Total:[3]"},
		{[]string{"neil", "matt", "adam", "linda"}, "UserList:{adam, linda, matt, neil} Total:[4]"},
	}
	for _, c := range tests1 {
		got, err := formatUserList(c.s)
		ok(t, err)
		equals(t, c.want, got)
	}
}

/////////////////////////////////////////////////////////////////////////////////////////

// assert fails the test if the condition is false.
func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d: "+msg+"\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// ok fails the test if an err is not nil.
func ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d: unexpected error: %s\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}
