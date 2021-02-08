package main

import (
	"fmt"
	"golang.org/x/mod/modfile"
	"io/ioutil"
)

func main() {
	b, err := ioutil.ReadFile("../go.mod")
	if err != nil {
		fmt.Println(err)
		return
	}
	file, err := modfile.Parse("../go.mod", b, func(path, version string) (s string, err error) {
		return "", nil
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(file.Module.Mod.Path)
}
