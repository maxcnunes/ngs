package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
)

const (
	colorCyan  = "\x1b[36m"
	colorGreen = "\x1b[32m"
	colorReset = "\x1b[0m"
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
	fmt.Println("===>"+colorCyan, gitPath, colorReset)

	gitDir := "--git-dir=" + gitPath
	workTree := "--work-tree=" + gitPath + "/.."

	args := []string{gitDir, workTree, "status", "-s"}
	cmd := exec.Command("git", args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func isGitRepository(gitPath string) bool {
	if _, err := os.Stat(gitPath); err != nil {
		return false
	}
	return true
}
