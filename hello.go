package main


import (
    "fmt"
    "text/template"
    "io/ioutil"
    "os"
    "path/filepath"
    "bytes"
    _"strings"
	"github.com/gobeam/stringy"
)

func renderTemplates(templatePath, outputPath string, data interface{}, force bool, namespace string) error {
    // Parse the outputPath to replace variables
    parsedOutputPath, err := replaceVariablesInString(outputPath, data)
    if err != nil {
        return err
    }

	if err := createOutputDirectory(parsedOutputPath); err != nil {
		return err
	}

	templateFiles, err := ioutil.ReadDir(templatePath)
	if err != nil {
		return err
	}

	for _, file := range templateFiles {
        println("Rendering: ", templatePath, file.Name())
		if file.IsDir() {
			subdir := filepath.Join(templatePath, file.Name())
			suboutput := filepath.Join(parsedOutputPath, file.Name())
			var subNamespace = namespace + "\\" + file.Name()

            if err := renderTemplates(subdir, suboutput, data, force, subNamespace); err != nil {
				return err
			}
			continue
		}

		if err := renderTemplateFile(templatePath, parsedOutputPath, file, data, force, namespace); err != nil {
			return err
		}
	}

	return nil
}

func createOutputDirectory(outputPath string) error {
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		return os.MkdirAll(outputPath, os.ModePerm)
	}
	return nil
}

func renderTemplateFile(templatePath, outputPath string, file os.FileInfo, data interface{}, force bool, namespace string) error {
    var newData = data.(map[string]interface{})
    newData["namespace"] = namespace
    
    templateContent, err := ioutil.ReadFile(filepath.Join(templatePath, file.Name()))
	if err != nil {
		return err
	}

	tmpl, err := parseTemplate(string(templateContent))
	if err != nil {
		return err
	}

	templateName := string(file.Name())
	newTemplateName, err := replaceVariablesInString(templateName, newData)
	if err != nil {
		return err
	}

	if !force {
		if _, err := os.Stat(filepath.Join(outputPath, newTemplateName)); !os.IsNotExist(err) {
			fmt.Println("File exists, skipping: ", newTemplateName)
			return nil
		}
	}

	outputFile, err := os.Create(filepath.Join(outputPath, newTemplateName))
	if err != nil {
		return err
	}
	defer outputFile.Close()

	if err := tmpl.Execute(outputFile, data); err != nil {
		return err
	}

	return nil
}

func parseTemplate(templateContent string) (*template.Template, error) {
	return template.New("template").Funcs(getFuncMap()).Parse(templateContent)
}

//create new template and return string from template
func replaceVariablesInString(input string, data interface{}) (string, error) {
	tmpl, err := parseTemplate(input)
    if err != nil {
        return "", err
    }

    var output bytes.Buffer

    if err := tmpl.Execute(&output, data); err != nil {
        return "", err
    }

    return output.String(), nil
}

func SnakeCase(s string) string {
    return stringy.New(s).SnakeCase().ToLower()
}

func CamelCase(s string) string {
    var camelCase = stringy.New(s).CamelCase()
    camelCase = stringy.New(camelCase).LcFirst()
    return camelCase
}

func PascalCase(s string) string {
    return stringy.New(s).CamelCase()
}

func KebabCase(s string) string {
    return stringy.New(s).KebabCase().Get()
}


func getFuncMap() template.FuncMap {
    return template.FuncMap{
        "SnakeCase":  SnakeCase,
        "CamelCase":  CamelCase,
        "KebabCase":  KebabCase,
        "PascalCase": PascalCase,
    }
}

func main() {
    pluginnamespace := "Swag\\MyFancyPlugin" //@todo: get from composer.json

    userInput := "SwagMyFancyEntity"
    entityName := PascalCase(userInput)

    templatePath := "templates/test"
    outputPath := "testoutput/Content/Entity/" + entityName

    namespace := pluginnamespace + "\\Content\\Entity\\" + entityName //@todo should bew generic or come from user input
    
    data := map[string]interface{}{
        "userInput": userInput,
        "entityName": entityName,
        "tableName": SnakeCase(userInput),
        "parentClassNamespace": namespace,
    }

    fmt.Println("Rendering templates...")

    if err := renderTemplates(templatePath, outputPath, data, true, namespace); err != nil {
        panic(err)
    }
}