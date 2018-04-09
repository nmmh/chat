package utils

import (
	"fmt"
	"sort"
	"strings"
)

//FormatUserList extracts,sorts adn returns a userlist string
func FormatUserList(usernameSlice []string) (string, error) {
	//sort first
	sort.Strings(usernameSlice)
	ul := "UserList:{"
	for i := 0; i < len(usernameSlice); i++ {
		ul += fmt.Sprintf("%s, ", usernameSlice[i])
	}
	ul = strings.TrimSuffix(ul, ", ") + fmt.Sprintf("} Total:[%d]", len(usernameSlice))
	return ul, nil
}

//UsernameInUse looksup a username returns true if found
func UsernameInUse(s []string, search string) (bool, error) {
	return StringInSlice(s, search)
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
