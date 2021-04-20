package bundle

import "github.com/shurcooL/vfsgen"

type Option struct {
	vfsgen.Options
}

func Generate(dir string, option Option) error {
	return vfsgen.Generate(Dir(dir), vfsgen.Options{
		Filename:        option.Filename,
		PackageName:     option.PackageName,
		BuildTags:       option.BuildTags,
		VariableName:    option.VariableName,
		VariableComment: option.VariableComment,
	})
}
