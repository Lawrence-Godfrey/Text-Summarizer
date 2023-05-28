package utils

import (
	"path/filepath"
	"runtime"
)

// GetRootDir Returns the full path to the root directory of the project.
func GetRootDir() string {
	// Get the full path to this file.
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}

	// Get the directory of this file.
	thisDir := filepath.Dir(filename)

	// Get the directory of the models directory, which is two directories up.
	rootDir := filepath.Dir(filepath.Dir(thisDir))

	return rootDir
}

// GetModelsDir Returns the full path to the `models` directory.
func GetModelsDir() string {
	// Get the directory of the models directory by joining the root directory and "models".
	rootDir := GetRootDir()
	ModelsDir := filepath.Join(rootDir, "models")

	return ModelsDir
}

// Flatten flattens a 2D slice into a 1D slice.
func Flatten[T any](slice [][]T) []T {
	var flattened []T
	for _, s := range slice {
		flattened = append(flattened, s...)
	}
	return flattened
}
