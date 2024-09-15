package main

import (
	"errors"
	"os"
)

type Clonable struct {
	URL string
	Dir string
}

type Gitmeta struct {
	Clonables []Clonable
}

type PluginIf interface {
	Applicable(e GitMetaEntry) bool
	Expand(e GitMetaEntry) ([]Clonable, error)
}

func (g *Gitmeta) FindPlugin(m GitMetaEntry) (PluginIf, error) {
	plugins := []PluginIf{
		NewGMFEntry(),
		NewGMFGithubUser(),
		NewGMFSSH(),
	}
	for _, plugin := range plugins {
		if plugin.Applicable(m) {
			return plugin, nil
		}
	}
	return nil, errors.New("no applicable plugin found")
}

func NewGitmeta() *Gitmeta {
	return &Gitmeta{
		Clonables: []Clonable{},
	}
}

func (g *Gitmeta) AddGMF(f *os.File) error {
	entries, err := g.parseGMF(f)
	if err != nil {
		return err
	}
	for _, e := range entries {
		p, err := g.FindPlugin(e)
		if err != nil {
			return err
		}
		clonables, err := p.Expand(e)
		if err != nil {
			return err
		}
		for _, clonable := range clonables {
			g.Clonables = append(g.Clonables, clonable)
		}
	}
	return nil
}

func (g *Gitmeta) AllClonables() []Clonable {
	return g.Clonables
}
