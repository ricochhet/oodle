package oodle

import "C"

import (
	"errors"
	"fmt"
	"strings"
	"unsafe"
)

const (
	CompressorInvalid = -1
	CompressorNone    = 3

	CompressorKraken    = 8
	CompressorLeviathan = 13
	CompressorMermaid   = 9
	CompressorSelkie    = 11
	CompressorHydra     = 12

	CompressorBitKnit = 10
	CompressorLZB16   = 4
	CompressorLZNA    = 7
	CompressorLZH     = 0
	CompressorLZHLW   = 1
	CompressorLZNIB   = 2
	CompressorLZBLW   = 5
	CompressorLZA     = 6

	CompressorCount   = 14
	CompressorForce32 = 0x40000000
)

const (
	CompressionLevelNone      = 0
	CompressionLevelSuperFast = 1
	CompressionLevelVeryFast  = 2
	CompressionLevelFast      = 3
	CompressionLevelNormal    = 4

	CompressionLevelOptimal1 = 5
	CompressionLevelOptimal2 = 6
	CompressionLevelOptimal3 = 7
	CompressionLevelOptimal4 = 8
	CompressionLevelOptimal5 = 9

	CompressionLevelHyperFast1 = -1
	CompressionLevelHyperFast2 = -2
	CompressionLevelHyperFast3 = -3
	CompressionLevelHyperFast4 = -4

	CompressionLevelHyperFast = CompressionLevelHyperFast1
	CompressionLevelOptimal   = CompressionLevelOptimal2
	CompressionLevelMax       = CompressionLevelOptimal5
	CompressionLevelMin       = CompressionLevelHyperFast4

	CompressionLevelForce32 = 0x40000000
	CompressionLevelInvalid = CompressionLevelForce32
)

type Compressor struct {
	Name  int
	Level int
}

// NewDefaultCompressor creates a default Decompressor with predefined values.
func NewDefaultCompressor() *Compressor {
	return &Compressor{
		Name:  CompressorKraken,
		Level: CompressionLevelOptimal,
	}
}

// Compress compresses the input buffer.
func (c *Compressor) Compress(input []byte) ([]byte, error) {
	_, err := lib.load()
	if err != nil {
		return nil, err
	}

	size := len(input)
	buf := make([]byte, size*2)

	var options uintptr
	var dictionaryBase uintptr
	var lrm uintptr
	var scratchMem uintptr
	var scratchSize uintptr

	r1 := compress(
		c.Name,
		unsafe.Pointer(&input[0]),
		size,
		unsafe.Pointer(&buf[0]),
		c.Level,
		options,
		dictionaryBase,
		lrm,
		scratchMem,
		scratchSize,
	)

	if r1 == 0 {
		return nil, errors.New("compress failure")
	}

	data := make([]byte, r1)
	copy(data, buf)

	return data, nil
}

// CompressionLevelGetName gets the name of the compression level integer.
func CompressionLevelGetName(level int) (string, error) {
	_, err := lib.load()
	if err != nil {
		return "", err
	}

	r1 := compressionLevelGetName(level)
	if r1 == 0 {
		return "", errors.New("error getting compression level name")
	}

	return C.GoString((*C.char)(unsafe.Pointer(r1))), nil
}

// CompressorGetName gets the name of the compressor integer.
func CompressorGetName(name int) (string, error) {
	_, err := lib.load()
	if err != nil {
		return "", err
	}

	r1 := compressorGetName(name)
	if r1 == 0 {
		return "", errors.New("error getting compressor name")
	}

	return C.GoString((*C.char)(unsafe.Pointer(r1))), nil
}

// CompressorToInt gets the integer of the compressor name.
func CompressorToInt(name string) (int, error) {
	switch strings.ToLower(name) {
	case "none":
		return CompressorNone, nil
	case "kraken":
		return CompressorKraken, nil
	case "leviathan":
		return CompressorLeviathan, nil
	case "mermaid":
		return CompressorMermaid, nil
	case "selkie":
		return CompressorSelkie, nil
	case "hydra":
		return CompressorHydra, nil
	default:
		return CompressorInvalid, fmt.Errorf("unknown compressor: %s", name)
	}
}

// CompressorLevelToInt gets the integer of the compressor level.
func CompressorLevelToInt(name string) (int, error) {
	switch strings.ToLower(name) {
	case "none":
		return CompressionLevelNone, nil
	case "superfast":
		return CompressionLevelSuperFast, nil
	case "veryfast":
		return CompressionLevelVeryFast, nil
	case "fast":
		return CompressionLevelFast, nil
	case "normal":
		return CompressionLevelNormal, nil
	case "optimal":
		return CompressionLevelOptimal, nil
	case "optimal1":
		return CompressionLevelOptimal1, nil
	case "optimal2":
		return CompressionLevelOptimal2, nil
	case "optimal3":
		return CompressionLevelOptimal3, nil
	case "optimal4":
		return CompressionLevelOptimal4, nil
	case "optimal5":
		return CompressionLevelOptimal5, nil
	case "hyperfast1":
		return CompressionLevelHyperFast1, nil
	case "hyperfast2":
		return CompressionLevelHyperFast2, nil
	case "hyperfast3":
		return CompressionLevelHyperFast3, nil
	case "hyperfast4":
		return CompressionLevelHyperFast4, nil
	default:
		return CompressionLevelInvalid, fmt.Errorf("unknown compression level: %s", name)
	}
}
