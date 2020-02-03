package initier

import (
	"fmt"
	"github.com/benka-me/hive-server-core/go-pkg/core"
	"testing"
)

func Test1(t *testing.T) {
	//Go := core.Go{}
	//fmt.Println(Go.GoString())
}

func TestGenerator (t *testing.T) {
	tests := []struct {
		name string
		cell core.Cell
	}{
		// TODO: Add test cases.
		{
			name: "test number 1",
			cell: core.Cell{
				Name:        "test-number-1",
				PkgName:     "number1",
				Repo:        "github.com/benka-me/test-number-1",
				Author:      "benka",
				AuthorEmail: "benka.node@gmail.com",
				Port:        8030,
				Public:		 true,
			},
		},
	}

	for _, tt := range tests {
		tt.cell.FillMeta()
		t.Run(tt.name, func(t *testing.T) {
			gens := tt.cell.Languages.ToGenerators()
			fmt.Println(gens)
		})
	}
}
func TestGenerateCell(t *testing.T) {
	tests := []struct {
		name string
		cell core.Cell
	}{
		// TODO: Add test cases.
		{
			name: "test number 1",
			cell: core.Cell{
				Name:        "test-number-1",
				PkgName:     "number1",
				Repo:        "github.com/benka-me/test-number-1",
				Author:      "benka",
				AuthorEmail: "benka.node@gmail.com",
				Port:        8030,
				Public:		 true,
			},
		},
	}

	for _, tt := range tests {
		tt.cell.FillMeta()
		t.Run(tt.name, func(t *testing.T) {
			err := tt.cell.GenerateFiles()
			if err != nil {
				t.Error(err)
			}

			err = tt.cell.SaveYaml()
			if err != nil {
				t.Error(err)
			}
		})
	}
}