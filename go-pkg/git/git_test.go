package git

import (
	"fmt"
	"testing"
)


func TestClone(t *testing.T) {
	tests := []struct {
		name string
		url string
	}{
		// TODO: Add test cases.
		{name: "github.com/benka-me/test-repo", url:"https://github.com/benka-me/test-repo"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := Clone(tt.url)
			if err != nil {
				t.Error(err)
			} else {
				fmt.Println(r)
			}
		})
	}
}