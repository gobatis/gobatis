package engine

import (
	"os"
	"path/filepath"
)

func (p *Engine) join(paths ...string) string {
	path, _ := os.Getwd()
	return filepath.Join(path, filepath.Join(paths...))
}
