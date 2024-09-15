package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		panic(fmt.Errorf("usage: %s repos.gmf", os.Args[0]))
	}

	cfg := os.Args[1]
	gmf := NewGitmeta()
	f, err := os.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = gmf.AddGMF(f)
	if err != nil {
		panic(err)
	}

	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	gitDir := filepath.Join(usr.HomeDir, "git")
	err = os.Chdir(gitDir)
	if err != nil {
		panic(err)
	}

	for _, c := range gmf.AllClonables() {
		err := cloneOrUpdate(c, gitDir)
		if err != nil {
			panic(err)
		}
	}
}
