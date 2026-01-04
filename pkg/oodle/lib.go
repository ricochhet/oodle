package oodle

import "C"

import (
	"fmt"
	"os"
	"unsafe"
)

type Lib struct {
	Name  string
	Paths []string
}

var LoadLib = NewLib()

// NewLib creates a new Lib with values defined per OS.
func NewLib() *Lib {
	return &Lib{Name: libName, Paths: []string{libName}}
}

//nolint:lll // wontfix
var (
	compress                func(int, unsafe.Pointer, int, unsafe.Pointer, int, uintptr, uintptr, uintptr, uintptr, uintptr) uintptr
	decompress              func(unsafe.Pointer, int, unsafe.Pointer, int64, uintptr, uintptr, uintptr, uintptr, uintptr, uintptr, uintptr, uintptr, uintptr, uintptr) uintptr
	compressionLevelGetName func(int) uintptr
	compressorGetName       func(int) uintptr
)

// ResolveLibPath resolves the possible paths for the library.
func (l *Lib) ResolveLibPath() (string, error) {
	for _, p := range l.Paths {
		_, err := os.Stat(p)
		if !os.IsNotExist(err) {
			return p, nil
		}
	}

	return "", fmt.Errorf("`%s` could not be resolved", libName)
}
