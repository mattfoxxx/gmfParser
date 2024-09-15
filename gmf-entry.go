package main

type GMFEntry struct{}

func NewGMFEntry() GMFEntry {
	return GMFEntry{}
}

func (g GMFEntry) Applicable(e GitMetaEntry) bool {
	return e.Type == ""

}

func (g GMFEntry) Expand(e GitMetaEntry) ([]Clonable, error) {
	return []Clonable{
		{URL: e.URL, Dir: e.Dir},
	}, nil
}
