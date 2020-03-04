package generator

import (
	"github.com/benka-me/hive/go-pkg/hive"
)

type LangGenerator interface {
	Protoc(*hive.Bee)
	ClientsFile(*hive.Bee) error
	ServerFiles(*hive.Bee) error
}

type LangGenerators *[]LangGenerator

func GetLangs(lgs *hive.Languages) (LangGenerators, error) {
	var langs = make([]LangGenerator, 2)
	langs[0] = Go(*lgs.GetGo())
	langs[1] = Javascript(*lgs.GetJavascript())

	return &langs, nil
}

func GetDevLang(b *hive.Bee) LangGenerator {
	var EnumLang = map[hive.DevLang]LangGenerator{
		0: Go(*b.Languages.Go),
		1: Javascript(*b.Languages.Javascript),
	}
	return EnumLang[b.DevLang]
}
