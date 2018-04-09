package utils

import (
	"fmt"
	"sort"
	"strings"
)

func formatUserList(usernameSlice []string) (string, error) {
	//sort first
	sort.Strings(usernameSlice)
	ul := "UserList:{"
	for i := 0; i < len(usernameSlice); i++ {
		ul += fmt.Sprintf("%s, ", usernameSlice[i])
	}
	ul = strings.TrimSuffix(ul, ", ") + fmt.Sprintf("} Total:[%d]", len(usernameSlice))
	return ul, nil
}

func usernameInUse(s []string, search string) (bool, error) {
	return stringInSlice(s, search)
}

func stringInSlice(s []string, search string) (bool, error) {
	for _, val := range s {
		if val == search {
			return true, nil
		}
	}
	return false, nil
}
