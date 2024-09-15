package main

import (
	"log"
	"os"
	"os/exec"
	"path"
)

func cloneOrUpdate(c Clonable, gitPath string) error {
	fullPath := path.Join(gitPath, c.Dir)
	_, err := os.Stat(fullPath)
	if os.IsNotExist(err) {
		cmd := exec.Command("git", "clone", c.URL, fullPath)
		stdout, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		if len(stdout) > 0 {
			log.Printf("%s is New => %s", c.Dir, stdout)
		}
	}

	err = os.Chdir(fullPath)
	if err != nil {
		return err
	}

	cmd := exec.Command("git", "pull")
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	if len(stdout) > 0 {
		log.Printf("%s is %s", c.Dir, stdout)
	}
	return nil
}
