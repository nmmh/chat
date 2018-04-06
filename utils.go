package main

import (
	"fmt"
	"strings"
)

func formatUserList(usernameSlice []string) string {
	ul := "UserList:{"
	for i := 0; i < len(usernameSlice); i++ {
		ul += fmt.Sprintf("%s, ", usernameSlice[i:i+1])
	}
	ul = strings.TrimSuffix(ul, ", ") + fmt.Sprintf("} Total:[%d]", len(usernameSlice))
	return ul
}

func usernameInUse(s []string, search string) bool {
	return stringInSlice(s, search)
}

func stringInSlice(s []string, search string) bool {
	for _, val := range s {
		if val == search {
			return true
		}
	}
	return false
}
