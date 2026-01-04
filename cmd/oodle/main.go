package main

import (
	"flag"
	"fmt"
	"os"

	oodle "github.com/ricochhet/oodle/pkg/oodle"
)

var (
	buildDate string
	gitHash   string
	buildOn   string
)

func version() string {
	return fmt.Sprintf(
		"Oodle\n\tBuild Date: %s\n\tGit Hash: %s\n\tBuilt On: %s\n",
		buildDate, gitHash, buildOn,
	)
}

func usage() {
	flag.Usage()
	os.Exit(0)
}

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}

	if flags.Unpack {
		_ = flags.unpack()
	}

	if flag.NArg() == 0 {
		usage()
	}

	switch flag.Args()[0] {
	case "compress":
		_ = flags.compress()
	case "decompress":
		_ = flags.decompress()
	case "version":
		fmt.Fprint(os.Stdout, version())
		os.Exit(0)
	default:
		usage()
	}
}

// compress command.
func (f *Flags) compress() error {
	if f.Input == "" || f.Output == "" {
		usage()
	}

	input, err := os.ReadFile(f.Input)
	if err != nil {
		return err
	}

	name, err := oodle.CompressorToInt(f.Compressor)
	if err != nil {
		return err
	}

	level, err := oodle.CompressorLevelToInt(f.Level)
	if err != nil {
		return err
	}

	compressor := oodle.Compressor{
		Name:  name,
		Level: level,
	}

	output, err := compressor.Compress(input)
	if err != nil {
		return err
	}

	return os.WriteFile(f.Output, output, 0o644)
}

// decompress command.
func (f *Flags) decompress() error {
	if f.Input == "" || f.Output == "" || f.Size <= 0 {
		usage()
	}

	input, err := os.ReadFile(f.Input)
	if err != nil {
		return err
	}

	decompressor := oodle.Decompressor{
		FuzzSafe:  uintptr(flags.FuzzSafe),
		CheckCRC:  uintptr(flags.CheckCRC),
		Verbosity: uintptr(flags.Verbosity),
		Decode:    uintptr(flags.Decode),
	}

	output, err := decompressor.Decompress(input, f.Size)
	if err != nil {
		return err
	}

	return os.WriteFile(f.Output, output, 0o644)
}
