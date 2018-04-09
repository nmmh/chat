package utils

import (
	"fmt"
	"sort"
	"strings"
)

//FormatUserList extracts,sorts adn returns a userlist string
func FormatUserList(usernames []string) (string, error) {
	//sort first
	sort.Strings(usernames)
	ul := "UserList:{"
	for _, username := range usernames {
		ul += fmt.Sprintf("%s, ", username)
	}
	ul = strings.TrimSuffix(ul, ", ") + fmt.Sprintf("} Total:[%d]", len(usernames))
	return ul, nil
}

//UsernameInUse looksup a username returns true if found
func UsernameInUse(usernames []string, search string) (bool, error) {
	return StringInSlice(usernames, search)
}

//StringInSlice looksup a string in a slice returns true if found
func StringInSlice(s []string, search string) (bool, error) {
	for _, val := range s {
		if val == search {
			return true, nil
		}
	}
	return false, nil
}
