package utility

import "path/filepath"

func GetCurrentPath(path string, dirs []string) string {
	main := filepath.Join(path, "release", "current", "schema")

	for _, dir := range dirs {
		main = filepath.Join(main, dir)
	}

	return main
}