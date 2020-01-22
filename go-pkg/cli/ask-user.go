package cli

import (
	"errors"
	"fmt"
	"github.com/benka-me/hive/go-pkg/library"
	"github.com/benka-me/hive/go-pkg/yaml"
	"github.com/go-playground/validator"
	"regexp"
	"strconv"
	"strings"
)

var v *validator.Validate

func fmtError (err string) string {
	switch err {
	case "required":
		return "Empty field is required"
	case "gte":
		return "Too short"
	case "lte":
		return "Too long"
	case "alpha":
		return "Alphabetic characters only"
	case "excludes":
		return "please avoid 'cell' or 'hive' word on your name"
	case "number":
		return "please enter valid number"
	case "max":
		return "too high"
	case "min":
		return "too low"
	case "email":
		return "please enter valid email"
	default:
		return ""
	}
}

type stringValid func (s string) error

func isAlphaDash (s string) error {
	var rgx = regexp.MustCompile(`^[a-zA-Z-]+$`)
	if rgx.MatchString(s) {
		return nil
	} else {
		return errors.New("Alphanumeric and dash only")
	}
}

func isAlphanumDash (s string) error {
	var rgx = regexp.MustCompile(`^[a-zA-Z0-9-_]+$`)
	if rgx.MatchString(s) {
		return nil
	} else {
		return errors.New("Alphanumeric and dash only")
	}
}

func Step(title, validate string, custom stringValid) string {
	var tmp string

	fmt.Println(title)
	_, _ = fmt.Scanf("%s", &tmp)

	err := custom(tmp)
	if err != nil {
		fmt.Println(err)
		return Step(title, validate, custom)
	} else if err := v.Var(tmp, validate); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			fmt.Println(fmtError(e.Tag()))
		}
		return Step(title, validate, custom)
	}
	return tmp
}

func askUser() (library.Cell, error) {
	cell := library.Cell{}
	v = validator.New()

	cell.Name = strings.ToLower(Step("Name of the new cell microservice: ", "required,lte=20,gte=3,excludes=cell,excludes=hive", isAlphanumDash))
	cell.PkgName = strings.ToLower(Step("Package name (2 - 4 chars long) for packages and types building", "required,lte=5,gte=2,alpha,excludes=cell,excludes=hive", isAlphaDash))
	cell.Repo = Step("Repository address", "required,gte=5", func(s string) error {return nil})
	cell.Author = Step("Author", "required,alphanumunicode,excludesrune= ", func(s string) error {return nil})
	cell.AuthorEmail = Step("Email", "required,email", func(s string) error {return nil})
	cell.Port, _ = strconv.Atoi(Step("Port", "required,number", func(s string) error {return nil}))
	cell.PkgNameCamel = strings.Title(kebabToCamelCase(cell.PkgName))

	err := yaml.SaveCell(cell)
	if err != nil {
		return cell, err
	}

	return cell, nil
}
