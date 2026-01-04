package oodle

import (
	"errors"
	"unsafe"
)

const (
	FuzzSafeNo  = 0
	FuzzSafeYes = 1
)

const (
	CheckCRCNo      = 0
	CheckCRCYes     = 1
	CheckCRCForce32 = 0x40000000
)

const (
	VerbosityNone    = 0
	VerbosityMinimal = 1
	VerbositySome    = 2
	VerbosityLots    = 3
	VerbosityForce32 = 0x40000000
)

const (
	DecodeThreadPhase1   = 1
	DecodeThreadPhase2   = 2
	DecodeThreadPhaseAll = 3
	DecodeUnthreaded     = DecodeThreadPhaseAll
)

type Decompressor struct {
	FuzzSafe  uintptr
	CheckCRC  uintptr
	Verbosity uintptr
	Decode    uintptr
}

// NewDefaultDecompressor creates a default Decompressor with predefined values.
func NewDefaultDecompressor() *Decompressor {
	return &Decompressor{
		FuzzSafe:  FuzzSafeNo,
		CheckCRC:  CheckCRCNo,
		Verbosity: VerbosityNone,
		Decode:    DecodeThreadPhaseAll,
	}
}

// Decompress decompresses the input buffer with the specified size.
func (d *Decompressor) Decompress(input []byte, bufSize int64) ([]byte, error) {
	_, err := lib.load()
	if err != nil {
		return nil, err
	}

	size := len(input)
	buf := make([]byte, bufSize)

	var (
		decBufBase        uintptr
		decBufSize        uintptr
		fpCallback        uintptr
		callbackUserData  uintptr
		decoderMemory     uintptr
		decoderMemorySize uintptr
	)

	r1 := decompress(
		unsafe.Pointer(&input[0]),
		size,
		unsafe.Pointer(&buf[0]),
		bufSize,
		d.FuzzSafe,
		d.CheckCRC,
		d.Verbosity,
		decBufBase,
		decBufSize,
		fpCallback,
		callbackUserData,
		decoderMemory,
		decoderMemorySize,
		d.Decode,
	)

	if r1 == 0 {
		return nil, errors.New("decompress failure")
	}

	return buf, nil
}
