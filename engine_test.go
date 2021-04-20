package gobatis

import (
	"github.com/koyeo/gobatis/bundle"
	"os"
	"path/filepath"
	"testing"
)

func TestEngine(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		t.Error(err)
		return
	}
	engine, err := NewEngine(bundle.Dir(filepath.Join(pwd, "test")))
	if err != nil {
		t.Error(err)
		return
	}
	engine.GetSQL("")
}
