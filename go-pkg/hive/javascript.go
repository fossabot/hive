package hive

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func (js Javascript) Protoc() {
	//for _, f := range js.Setup.Files {
	//	js.oneProtoc(f)
	//}
	repoPath := fmt.Sprintf("%s/src/%s", gopath, js.Setup.Repo)
	bin := fmt.Sprintf("%s/src/github.com/benka-me/hive/js-pkg/node_modules/.bin/protoc-gen-ts", gopath)
	jsOut := fmt.Sprintf("%s/src", gopath)
	args := make([]string, 5)
	args = []string{
		fmt.Sprintf("-I=%s/src/", gopath),
		fmt.Sprintf("-I=%s/src/github.com/gogo/protobuf/protobuf", gopath),
		fmt.Sprintf("--plugin=protoc-gen-ts=%s", bin),
		fmt.Sprintf("--ts_out=service=true:%s", jsOut),
		fmt.Sprintf("--js_out=import_style=commonjs,binary:%s", jsOut),
	}

	for _, f := range js.Setup.Files {
		args = append(args, fmt.Sprintf("%s/protobuf/%s", repoPath, f))
	}

	cmd := exec.Command("protoc", args...)
	var out bytes.Buffer
	var errs bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr= &errs
	fmt.Println(cmd.Args)
	err := cmd.Run()
	if err != nil {
		fmt.Println(errs.String())
	}
	fmt.Printf("%s\n", out.String())
}

func (js Javascript) oneProtoc (file string) {
	repoPath := fmt.Sprintf("%s/src/%s", gopath, js.Setup.Repo)
	bin := fmt.Sprintf("%s/src/hive/js-pkg/node_modules/.bin/protoc-gen-ts", gopath)
	jsOut := fmt.Sprintf("%s/js-pkg", repoPath)
	args := []string{
		"-I=.",
		fmt.Sprintf("-I=%s/protobuf", repoPath),
		fmt.Sprintf("--plugin=protoc-gen-ts=%s", bin),
		fmt.Sprintf("--ts_out=service=true:%s", jsOut),
		fmt.Sprintf("--js_out=import_style=commonjs,binary:%s", jsOut),
		file,
	}
	cmd := exec.Command("protoc", args...)
	var out bytes.Buffer
	var errs bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr= &errs
	fmt.Println(cmd.Args)
	err := cmd.Run()
	if err != nil {
		fmt.Println(errs.String())
	}
	fmt.Printf("%s\n", out.String())
}

func (js Javascript) PublicFiles (bee *Bee) error {
	repo := js.Setup.Repo
	repoPath := fmt.Sprintf("%s/src/%s", gopath, repo)
	var perm os.FileMode = 0777 //TODO change file mode

	//create directories
	err := os.MkdirAll(fmt.Sprintf("%s/js-pkg/proto", repoPath), perm)
	if err != nil {
		return err
	}
	return nil
}
