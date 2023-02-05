package twofer

import "strings"

func ShareWith(name string) string {
	if len(strings.TrimSpace(name)) == 0 {
		return "One for you, one for me."
	}
	return "One for " + name + ", one for me."
}
