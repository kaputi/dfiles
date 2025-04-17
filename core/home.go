package core

import "os"

var home = os.Getenv("HOME")

func AddHome(path string) string {
	return home + "/" + path
}

func RemoveHome(path string) string {
	return path[len(home)+1:]
}
