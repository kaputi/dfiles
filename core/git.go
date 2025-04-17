package core

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// comand
// git --git-dir=repoPath --work-tree=$HOME'
var repoPath = os.Getenv("HOME") + "/.local/share/dfiles/repo"
var gitDirFlag = "--git-dir=" + repoPath
var workTreeFlag = "--work-tree=" + os.Getenv("HOME")

func gitCmd(args ...string) []byte {
	checkInit()
	allArgs := append([]string{gitDirFlag, workTreeFlag}, args...)
	cmd := exec.Command("git", allArgs...)
	output, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatal(err)
	}

	return output
}

func checkDir() {
	if _, err := os.Stat(repoPath); os.IsNotExist(err) {
		err := os.MkdirAll(repoPath, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func isGitRepo() (bool, error) {
	info, err := os.Stat(repoPath)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}

	return info.IsDir(), nil
}

// execute git init --bare and config
func gitInit() {
	if ok, err := isGitRepo(); err != nil {
		log.Fatal(err)

	} else if !ok {
		checkDir()
		initCmd := exec.Command("git", "init", "--bare", repoPath)
		output, err := initCmd.CombinedOutput()

		gitCmd("config", "--local", "status.showUntrackedFiles", "no")

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(output))
	}
}

func checkInit() {
	if ok, err := isGitRepo(); err != nil {
		log.Fatal(err)
	} else if !ok {
		gitInit()
	}
}
