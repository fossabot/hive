package generator

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"text/template"
)

func (code Code) Generate() error {
	dat, err := ioutil.ReadFile(code.Template)
	if err != nil {
		return err
	}

	tmpl, err := template.New(code.Name).Parse(string(dat))
	if err != nil {
		return err
	}

	f, err := os.Create(code.Target)
	if err != nil {
		return err
	}

	err = tmpl.Execute(f, code.Interface)
	if err != nil {
		return err
	}
	return nil
}

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
