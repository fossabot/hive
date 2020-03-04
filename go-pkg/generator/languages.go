package generator

import (
	"github.com/benka-me/hive/go-pkg/hive"
)

type LangGenerator interface {
	Protoc(*hive.Bee)
	ClientsFiles(*hive.Bee) error
	EntryPointFiles(*hive.Bee) error
}

type PkgsGenerators *[]LangGenerator

func ToPkgsGenerators(lgs *hive.Languages) (PkgsGenerators, error) {
	var langs = make([]LangGenerator, 2)
	langs[0] = Go(*lgs.GetGo())
	langs[1] = Javascript(*lgs.GetJavascript())

	return &langs, nil
}

func ServerLang(b *hive.Bee) LangGenerator {
	var EnumLang = map[hive.DevLang]LangGenerator{
		0: Go(*b.Languages.Go),
		1: Javascript(*b.Languages.Javascript),
	}
	return EnumLang[b.DevLang]
}
