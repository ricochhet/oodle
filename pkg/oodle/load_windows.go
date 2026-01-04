package oodle

import (
	"sync"

	"github.com/ebitengine/purego"
	"golang.org/x/sys/windows"
)

const libName = "oo2core_9_win64.dll"

var once struct {
	sync.Once

	handle uintptr
	err    error
}

// load loads the library.
func (l *Lib) load() (uintptr, error) {
	once.Do(func() {
		lib, err := l.ResolveLibPath()
		if err != nil {
			once.err = err
			return
		}

		handle, err := windows.LoadLibrary(lib)
		if err != nil {
			once.err = err
			return
		}

		once.handle = uintptr(handle)

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
