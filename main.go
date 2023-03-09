package main

import (
	"github.com/sethvargo/go-githubactions"
	"golang.org/x/mod/modfile"
	"os"
)

const (
	MODFILE = "modfile"
)

var (
	version string
	module  string
)

func main() {
	action := githubactions.New()
	mod := action.GetInput(MODFILE)
	if mod == "" {
		mod = "go.mod"
	}
	action.Debugf("Using %s file", mod)
	data, err := os.ReadFile(mod)
	if err != nil {
		action.Fatalf("%e", err)
		return
	}
	m, err := modfile.Parse(mod, data, nil)
	if err != nil {
		action.Fatalf("%e", err)
		return
	}
	if m.Go != nil {
		version = m.Go.Version
	}

	if m.Module != nil {
		module = m.Module.Mod.Path
	}

	action.SetOutput("go_version", version)
	action.SetOutput("go_module", module)
}
