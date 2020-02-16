package hive

import "reflect"

type LangGenerator interface {
	Protoc()
	PublicFiles(*Bee) error
}

type LangGenerators *[]LangGenerator

func (lgs *Languages) ToGenerators() LangGenerators {
	var langs = make([]LangGenerator, 0)
	arr := reflect.ValueOf(*lgs)
	for i := 0; i < arr.NumField() ; i++ {
		f := arr.Field(i)
		inter := f.Interface().(LangGenerator)
		langs = append(langs, inter)
	}
	return &langs
}
