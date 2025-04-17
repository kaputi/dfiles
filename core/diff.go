package core

import (
	"log"
	"os"
	"slices"

	"github.com/kaputi/dfiles/config"
)

// TODO: read the include and ignore files
// for wach include file
var include = []string{} // this are relative to $HOME
var exclude = []string{} // this are relative to $HOME
var tracked = []string{} // this are relative to $HOME

func Untracked() []string {
	config := config.ReadConfig()

	fileList := []string{}

	for _, path := range config.Include {
		absolutePath := AddHome(path)

		info, err := os.Stat(absolutePath)
		if err != nil {
			log.Fatalf("failed to stat path: %v", err)
		}
		if !info.IsDir() {
			if slices.Contains(tracked, RemoveHome(path)) {
				continue
			}
			fileList = append(fileList, path)
		} else {
			innerFiles, err := CollectFiles(absolutePath)
			if err != nil {
				log.Fatalf("failed to collect files: %v", err)
			}
			fileList = append(fileList, innerFiles...)
		}
	}

	return fileList
}
