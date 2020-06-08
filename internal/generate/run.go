package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
)

// ReplaceTemplate is the value of filename text to replace
var ReplaceTemplate = ".template"

// DefaultTypes are the types to auto-generate for
var DefaultTypes = []string{
	"bool",
	"int",
	"int64",
	"uint",
	"uint64",
	"float64",
	"string",
	"time.Time",
	"time.Duration",
}

var nameTransforms = map[string]string{
	"interface": "generic",
	"value":     "generic",
}

// GenSlice auto-generates slices for types
var GenSlice = true

var path, types string

type data struct {
	Type       string
	Elem       string
	Name       string
	Title      string
	IsSlice    bool
	TakesValue bool
}

var fields []string

func init() {
	st := reflect.TypeOf(data{})
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		fields = append(fields, field.Name)
	}
}

func main() {
	flag.StringVar(&path, "path", ".", "The path to generate from")
	flag.StringVar(&types, "types", strings.Join(DefaultTypes, ","), "The subtypes used for the queue being generated")
	flag.BoolVar(&GenSlice, "slice", GenSlice, "Auto-generate slice types")
	flag.Parse()
	if types == "" {
		log.Fatal(fmt.Errorf("Type should be defined"))
	}
	if err := Generate(path); err != nil {
		log.Fatal(err)
	}
}

// Generate searches for files that match ReplaceTemplate and generates based on DefaultTypes and GenSlice
func Generate(path string) error {
	return filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if path != "." && strings.HasPrefix(info.Name(), ".") {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		if info.Name() == "vendor" && info.IsDir() {
			return filepath.SkipDir
		}

		if strings.Contains(info.Name(), ReplaceTemplate) {
			for _, t := range strings.Split(types, ",") {
				if err := genTemplate(path, t); err != nil {
					return err
				}
				if GenSlice && !strings.HasPrefix(t, "[]") {
					if err := genTemplate(path, "[]"+t); err != nil {
						return err
					}
				}
			}
		}

		return nil
	})
}

func genTemplate(path, typeString string) error {
	dir := filepath.Dir(path)
	base := filepath.Base(path)

	templateInfo := genTemplateInfo(path, typeString)

	replacementText := fmt.Sprintf(".zz_generated_%s", templateInfo.Name)
	generatedPath := strings.TrimPrefix(base, "_")
	generatedPath = strings.ReplaceAll(generatedPath, ReplaceTemplate, replacementText)
	generatedPath = filepath.Join(dir, generatedPath)

	fileContents, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	fileTemplate := regexp.MustCompile(`(?i)(`+strings.Join(fields, "|")+`)__`).ReplaceAllString(string(fileContents), "{{.$1}}")
	fileTemplate = regexp.MustCompile(`{{\..`).ReplaceAllStringFunc(fileTemplate, strings.ToUpper)

	fmt.Printf("Generating %s\n", generatedPath)

	t := template.Must(template.New(path).Parse(fileTemplate))
	f, err := os.Create(generatedPath)
	if err != nil {
		return err
	}
	if err := t.Execute(f, templateInfo); err != nil {
		return err
	}
	return nil
}

func genTemplateInfo(path, t string) data {
	packageName := strings.ReplaceAll(filepath.Dir(path), "/", ".") + "."
	t = strings.ReplaceAll(t, packageName, "")

	typeInfo := strings.TrimSpace(t)
	elemInfo := strings.TrimPrefix(typeInfo, "[]")

	nameInfo := elemInfo
	nameInfo = regexp.MustCompile(`^.*\.`).ReplaceAllString(nameInfo, "")
	nameInfo = strings.TrimSuffix(nameInfo, "{}")
	nameInfo = strings.ToLower(nameInfo)
	if v, ok := nameTransforms[nameInfo]; ok {
		nameInfo = v
	}

	isSliceInfo := false

	if strings.HasPrefix(typeInfo, "[]") {
		nameInfo = fmt.Sprintf("%sSlice", nameInfo)
		isSliceInfo = true
	}

	titleInfo := strings.Title(nameInfo)

	return data{
		Type:       typeInfo,
		Elem:       elemInfo,
		Name:       nameInfo,
		Title:      titleInfo,
		IsSlice:    isSliceInfo,
		TakesValue: elemInfo != "bool",
	}
}
