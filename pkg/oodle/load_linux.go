package oodle

import (
	"sync"

	"github.com/ebitengine/purego"
)

const libName = "liboo2corelinux64.so.9"

var once struct {
	sync.Once

	handle uintptr
	err    error
}

// load loads the library.
func (l *Lib) load() (uintptr, error) {
	once.Do(func() {
		lib, err := l.resolveLibPath()
		if err != nil {
			once.err = err
			return
		}

		handle, err := purego.Dlopen(lib, purego.RTLD_NOW|purego.RTLD_GLOBAL)
		if err != nil {
			once.err = err
			return
		}

		once.handle = handle

		purego.RegisterLibFunc(&compress, once.handle, "OodleLZ_Compress")
		purego.RegisterLibFunc(&decompress, once.handle, "OodleLZ_Decompress")
		purego.RegisterLibFunc(
			&compressionLevelGetName,
			once.handle,
			"OodleLZ_CompressionLevel_GetName",
		)
		purego.RegisterLibFunc(&compressorGetName, once.handle, "OodleLZ_Compressor_GetName")
	})

	return once.handle, once.err
}
