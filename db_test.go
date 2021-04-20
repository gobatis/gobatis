package gobatis

import (
	"github.com/koyeo/gobatis/bundle"
	"os"
	"path/filepath"
	"testing"
)

func TestNewDB(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		t.Error(err)
		return
	}
	db := NewPostgresql("postgresql://postgres:postgres@127.0.0.1:54322/angel?connect_timeout=10&sslmode=disable")
	db.SetBundle(bundle.Dir(filepath.Join(pwd, "test")))
	err = db.Init()
	if err != nil {
		t.Error("DB init error:", err)
		return
	}
	err = db.Ping()
	if err != nil {
		t.Error("DB ping error:", err)
		return
	}
	defer func() {
		err = db.Close()
		if err != nil {
			t.Error("DB close error:", err)
			return
		}
	}()
}

func TestReflect(t *testing.T) {



}
