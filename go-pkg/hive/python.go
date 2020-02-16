package hive

import (
	"fmt"
)

func (py Python) Protoc() {
	for _, f := range py.Setup.Files {
		py.oneProtoc(f)
	}
}

func (py Python) oneProtoc (file string) {
	fmt.Println("python.go oneProtoc must be implemented")
	//args := []string{
	//	"-I=.",
	//	fmt.Sprintf("-I=%s/src/", gopath),
	//	fmt.Sprintf("-I=%s/src/github.com/gogo/protobuf/protobuf", gopath),
	//	fmt.Sprintf("-I=%s/src/%s/protobuf", gopath, py.Proto.Repo),
	//	fmt.Sprintf("--gogoslick_out=plugins=grpc:%s/src", gopath),
	//	file,
	//}
	//cmd := exec.Command("protoc", args...)
	//var out bytes.Buffer
	//var errs bytes.Buffer
	//cmd.Stdout = &out
	//cmd.Stderr= &errs
	//fmt.Println(cmd.Args)
	//err := cmd.Run()
	//if err != nil {
	//	fmt.Println(errs.String())
	//}
	//fmt.Printf("%s\n", out.String())
}

func (py Python) PublicFiles (bee *Bee) error {
	return nil
}
