package main

import (
	"embed"
	"os"
	"path"
	"path/filepath"

	"github.com/ricochhet/oodle/pkg/oodle"
)

//go:embed libs
var fs embed.FS

// unpack unpacks the embedded oodle libraries.
func (f *Flags) unpack() error {
	_, err := oodle.LoadLib.ResolveLibPath()
	if err == nil {
		return nil
	}

	b, err := fs.ReadFile(path.Join("libs", oodle.LoadLib.Name))
	if err != nil {
		return err
	}

	return write(oodle.LoadLib.Name, b)
}

// write writes to the specified path with the provided data.
func write(path string, data []byte) error {
	err := os.MkdirAll(filepath.Dir(path), 0o755)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, data, 0o644)
	if err != nil {
		return err
	}

	return nil
}
