// Read compser json file

package repo

import (
	"strings"
)

type Composer struct {
	name        string
	libraryType string
	extra       [] string
}

func SearchComposer(file string) bool {
	if strings.Contains(file, ".json") {
		return true
	}

	return false
}