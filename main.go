package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"regexp"
)

const (
	colorBlue  = "\x1b[1;34m"
	colorGreen = "\x1b[32m"
	colorReset = "\x1b[0m"

	okSymbol    = "\x1b[1;32m\xE2\x9c\x93\x1b[0m"
	wrongSymbol = "\x1b[1;31m\xE2\x9c\x97\x1b[0m"
	arrowSymbol = "\x1b[1;36m\xC2\xBB\x1b[0m"
)

func main() {
	dir := flag.String("dir", ".", "base directory")

	flag.Parse()

	files, _ := ioutil.ReadDir(*dir)
	for _, f := range files {
		gitPath := path.Join(*dir, f.Name(), ".git")

		if isGitRepository(gitPath) {
			gitStatus(gitPath)
		} else if f.IsDir() && f.Name() == ".git" {
			gitStatus(*dir + "/.git")
		}
	}
}

func gitStatus(gitPath string) {
	printRepositoryName(gitPath)

	gitDir := "--git-dir=" + gitPath
	workTree := "--work-tree=" + gitPath + "/.."

	args := []string{gitDir, workTree, "status", "-s"}
	cmd := exec.Command("git", args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func isGitRepository(gitPath string) bool {
	_, err := os.Stat(gitPath)
	return err == nil
}

func printRepositoryName(gitPath string) {

	reg := regexp.MustCompile(`^(.*)\.git$`)

	name := reg.ReplaceAllString(gitPath, "$1")
	fmt.Println(arrowSymbol, colorBlue, name, colorReset)
}
