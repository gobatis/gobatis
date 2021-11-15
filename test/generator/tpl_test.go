package generator

import "testing"

func TestRenderEntity(t *testing.T) {
	RenderEntity("./entity.make.go2", GoHeader{
		Package: "tpl",
		Imports: []string{"fmt"},
	},
		[]*Entity{{
			Name: "GoType",
			Params: []*Param{
				{
					Name: "TBigint",
					Type: "int64",
					Tag:  "t_bigint",
				},
			},
		}})
}
