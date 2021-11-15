package generator

import (
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"path/filepath"
)

func resolve(path string) string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return filepath.Join(wd, path)
}

func writeFile(relativePath, content string) (err error) {
	err = ioutil.WriteFile(resolve(relativePath), []byte(content), 0644)
	if err != nil {
		return
	}
	return
}

func gofmt(buf string) string {
	formatted, err := format.Source([]byte(buf))
	if err != nil {
		panic(fmt.Errorf("%s\nOriginal code:\n%s", err.Error(), buf))
	}
	return string(formatted)
}
