package main

import "flag"

type Flags struct {
	Unpack bool

	Input      string
	Output     string
	Compressor string
	Level      string
	Size       int64
	DirExt     string

	FuzzSafe  int
	CheckCRC  int
	Verbosity int
	Decode    int
}

var flags = NewFlags()

func NewFlags() *Flags {
	return &Flags{}
}

//nolint:gochecknoinits // wontfix
func init() {
	registerFlags(flag.CommandLine, flags)
	flag.Parse()
}

// registerFlags registers all flags to the flagset.
func registerFlags(fs *flag.FlagSet, f *Flags) {
	fs.BoolVar(&f.Unpack, "u", true, "unpack embedded library")
	fs.StringVar(&f.Input, "i", "", "input file")
	fs.StringVar(&f.Output, "o", "", "output file")
	fs.StringVar(&f.Compressor, "c", "kraken", "compressor")
	fs.StringVar(&f.Level, "l", "optimal", "compression level")
	fs.Int64Var(&f.Size, "s", -1, "original uncompressed size")
	fs.StringVar(&f.DirExt, "e", ".oodle", "directory extension")
	fs.IntVar(&f.FuzzSafe, "fs", 0, "fuzz safe")
	fs.IntVar(&f.CheckCRC, "crc", 0, "check crc")
	fs.IntVar(&f.Verbosity, "v", 0, "verbosity")
	fs.IntVar(&f.Decode, "d", 3, "decode")
}
