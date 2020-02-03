package generator

import (
	"os"
)

var gopath = os.Getenv("GOPATH")

type Langs interface {
	ExecLangProtoc()
}

type Protobuf struct {
	Repo string
	Files []string
}

type Code struct {
	Interface interface{}
	Template string
	Target string
	Name string
}
