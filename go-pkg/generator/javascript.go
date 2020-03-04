package generator

import (
	"fmt"
	"github.com/benka-me/hive/go-pkg/hive"
)

type Javascript hive.Javascript

func (js Javascript) ClientsFile(bee *hive.Bee) error {
	return nil
}

func (js Javascript) ServerFiles(bee *hive.Bee) error {
	//repo := bee.Repo
	//repoPath := fmt.Sprintf("%s/src/%s", gopath, repo)
	//var perm os.FileMode = 0777 //TODO change file mode
	return nil
}

func (js Javascript) Protoc(bee *hive.Bee) {
	repoPath := fmt.Sprintf("%s/src/%s", gopath, bee.Repo)
	bin := fmt.Sprintf("%s/src/github.com/benka-me/hive/js-pkg/node_modules/.bin/protoc-gen-ts", gopath)
	jsOut := fmt.Sprintf("./js-pkg/src/protobuf")
	args := make([]string, 5)
	args = []string{
		fmt.Sprintf("-I=%s/protobuf", repoPath),
		fmt.Sprintf("-I=%s/src/github.com/gogo/protobuf/protobuf", gopath),
		fmt.Sprintf("--plugin=protoc-gen-ts=%s", bin),
		fmt.Sprintf("--ts_out=service=true:%s", jsOut),
		fmt.Sprintf("--js_out=import_style=commonjs,binary:%s", jsOut),
	}

	for _, f := range bee.ProtoSetup.Files {
		args = append(args, fmt.Sprintf("%s", f))
	}
	runProtocCommand(args)
}

