package generator

import (
	"bytes"
	"fmt"
	"github.com/benka-me/hive/go-pkg/hive"
	"os"
	"os/exec"
)

type Javascript hive.Javascript

func (js Javascript) ClientsFiles(bee *hive.Bee) error {
	return nil
}

func (js Javascript) EntryPointFiles(bee *hive.Bee) error {
	repo := bee.Repo
	repoPath := fmt.Sprintf("%s/src/%s", gopath, repo)
	var perm os.FileMode = 0777 //TODO change file mode

	//create directories
	err := os.MkdirAll(fmt.Sprintf("%s/js-pkg/proto", repoPath), perm)
	if err != nil {
		return err
	}
	return nil
}


protoc \
-I=/Users/benka/go/src/github.com/benka-me/test-repo \
-I=/Users/benka/go/src/github.com/gogo/protobuf/protobuf \
--plugin=protoc-gen-ts=/Users/benka/go/src/github.com/benka-me/hive/js-pkg/node_modules/.bin/protoc-gen-ts \
--ts_out=service=true:./js-pkg \
--js_out=import_style=commonjs,binary:./js-pkg \
protobuf/trepo.proto protobuf/rpc-trepo.proto
func (js Javascript) Protoc(bee *hive.Bee) {
	repoPath := fmt.Sprintf("%s/src/%s", gopath, bee.Repo)
	bin := fmt.Sprintf("%s/src/github.com/benka-me/hive/js-pkg/node_modules/.bin/protoc-gen-ts", gopath)
	jsOut := fmt.Sprintf("./js-pkg")
	args := make([]string, 5)
	args = []string{
		fmt.Sprintf("-I=%s", repoPath),
		fmt.Sprintf("-I=%s/src/github.com/gogo/protobuf/protobuf", gopath),
		fmt.Sprintf("--plugin=protoc-gen-ts=%s", bin),
		fmt.Sprintf("--ts_out=service=true:%s", jsOut),
		fmt.Sprintf("--js_out=import_style=commonjs,binary:%s", jsOut),
	}

	for _, f := range bee.ProtoSetup.Files {
		args = append(args, fmt.Sprintf("protobuf/%s", f))
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

