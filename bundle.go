package gobatis

import (
	"fmt"
	"github.com/shurcooL/vfsgen"
	"net/http"
	"os"
	"path/filepath"
)

type Bundle = http.FileSystem

type BundleOption struct {
	Filename        string
	PackageName     string
	BuildTags       string
	VariableName    string
	VariableComment string
}

func NewBundle(dir string) (bundle Bundle) {
	var err error
	defer func() {
		if err != nil {
			panic(fmt.Errorf("new bundle error: %s", err))
		}
	}()
	pwd, err := os.Getwd()
	if err != nil {
		return
	}
	path := filepath.Join(pwd, dir)
	info, err := os.Stat(path)
	if err != nil {
		return
	}
	if !info.IsDir() {
		err = fmt.Errorf("bundle path is not dir: %s", path)
		return
	}
	bundle = fsDir(path)
	return
}

func GenerateBundle(dir string, option BundleOption) error {
	return vfsgen.Generate(fsDir(dir), vfsgen.Options{
		Filename:        option.Filename,
		PackageName:     option.PackageName,
		BuildTags:       option.BuildTags,
		VariableName:    option.VariableName,
		VariableComment: option.VariableComment,
	})
}
