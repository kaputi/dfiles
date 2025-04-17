package core

import "strings"

func GetTraked() []string {
	cmd := gitCmd("ls-files")
	tracked := strings.Split(string(cmd), "\n")
	return tracked[:len(tracked)-1]
}
