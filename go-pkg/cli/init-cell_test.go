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
				Name:        "testnumber-1",
				PkgName:     "asdf",
				Repo:        "github.com/benka-me/test-number-1",
				Author:      "benka",
				AuthorEmail: "test@gmail.com",
				Port:        3456,
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