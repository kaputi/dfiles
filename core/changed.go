package core

import "strings"

func GetChanged() []string {

	// TODO : add files that are inside of tracked directories but not tracked

	cmd := gitCmd("status", "--porcelain")

	changed := strings.Split(string(cmd), "\n")
	return changed[:len(changed)-1]
}
