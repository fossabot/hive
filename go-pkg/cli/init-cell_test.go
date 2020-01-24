package cli

import (
	"github.com/benka-me/hive/go-pkg/library"
	"github.com/benka-me/hive/go-pkg/yaml"
	"testing"
)

func Test_createProtobuf(t *testing.T) {
	type args struct {
		cell library.Cell
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "hive server api",
			args: args{cell:library.Cell{
				Name:        "hive-server-api",
				PkgName:     "api",
				Repo:        "github.com/benka-me/hive-server-api",
				Author:      "benka",
				AuthorEmail: "benka.node@gmail.com",
				Port:        8021,
				Public: false,
			}},
		},
		{
			name: "hive server core",
			args: args{cell:library.Cell{
				Name:        "hive-server-core",
				PkgName:     "core",
				Repo:        "github.com/benka-me/hive-server-core",
				Author:      "benka",
				AuthorEmail: "benka.node@gmail.com",
				Port:        8020,
				Public: false,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := yaml.SaveCell(tt.args.cell)
			if err != nil {
				t.Error(err)
			}

			err = cellCreateFiles(tt.args.cell)
			if err != nil {
				t.Error(err)
			}
		})
	}
}