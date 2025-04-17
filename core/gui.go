package core

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

func Separator() {
	// get terminal width
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Println(strings.Repeat("-", 80))
		return
	}

	fmt.Println(strings.Repeat("-", width))
}

func PrintList(list []string) {
	for _, item := range list {
		fmt.Println(item)
	}
}
