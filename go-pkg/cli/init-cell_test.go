package cli

import (
	"github.com/benka-me/hive/go-pkg/library"
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
			name: "regular1",
			args: args{cell:library.Cell{
				Name:        "cell-test",
				PkgName:     "test",
				Repo:        "github.com/benka-me/test",
				Author:      "benka",
				AuthorEmail: "test@gmail.com",
				Port:        2000,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := createFiles(tt.args.cell)
			if err != nil {
				t.Error(err)
			}
		})
	}
}