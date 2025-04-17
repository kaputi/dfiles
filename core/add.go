package core

import (
	"log"
	"slices"

	"github.com/kaputi/dfiles/config"
)

func Add(path string) {
	config := config.ReadConfig()

	if slices.Contains(config.Exclude, path) {
		log.Printf("%v is in the exclude list\n", path)
		return
	}

	if slices.Contains(config.Include, path) {
		log.Printf("%v is already added\n", path)
		return
	}

	config.Include = append(config.Include, path)

	gitCmd("add", path)
}
