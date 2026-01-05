package fsutil

import "os"

const (
	None = iota
	Dir
	File
)

// Read reads the file at the specified path.
func Read(path string) ([]byte, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// Write writes the specified data to path.
func Write(path string, data []byte) error {
	return os.WriteFile(path, data, 0o644)
}

// IsDirOrFile returns 1 if the file is a directory, 2 if it's a file, and none if there's an error.
func IsDirOrFile(path string) (int, error) {
	i, err := os.Stat(path)
	if err != nil {
		return None, err
	}

	if i.IsDir() {
		return Dir, nil
	}

	if i.Mode().IsRegular() {
		return File, nil
	}

	return None, nil
}
