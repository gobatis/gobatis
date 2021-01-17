package engine

import (
	"path/filepath"
	"strings"
)

func (p *Engine) join(paths ...string) string {
	//path, _ := os.Getwd()
	//return filepath.Join(path, filepath.Join(paths...))
	return strings.TrimPrefix(filepath.Join(paths...), "/")
}

func (p *Engine) resolve(pkgPath string) string {
	return p.join(strings.TrimPrefix(pkgPath, p.module))
}
