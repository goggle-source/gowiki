package generator

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"

	"github.com/goggle-source/gowiki/internal/models"
	"github.com/goggle-source/gowiki/internal/parser"
)

func GenerateReadyHTML(pathDirContent, pathDirTemplates, pathDirForReadyHTML string) error {
	mdFiles, err := parser.GetHTMLContent(pathDirContent)
	if err != nil {
		return err
	}

	for _, file := range mdFiles {
		metadata, err := parser.GetMetadataMdFile(file.Path)
		if err != nil {
			panic(fmt.Errorf("%s:%w", "getmetadata", err))
		}
		ServicTemplate := parser.Init(pathDirTemplates)
		template := ServicTemplate.GetTemplate(metadata["name_template"].(string))
		path := filepath.Join(pathDirForReadyHTML, metadata["name_ready_html"].(string))
		fileForReadyHTML, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0773)
		if err != nil {
			panic(fmt.Errorf("%s:%w", "openfile", err))
		}

		if metadata["name_template"].(string) == "test1.html" {
			err := genereteHTMLTest1(template, file.Data, metadata, fileForReadyHTML)
			if err != nil {
				panic(fmt.Errorf("%s:%w", "generateReadyHTML", err))
			}
		} else {
			fmt.Println("error")
		}
	}
	return nil
}

func genereteHTMLTest1(templateHTML *template.Template, data bytes.Buffer, metadata map[string]any, w io.Writer) error {

	model := models.Test1{
		Name:      template.HTML(metadata["Name"].(string)),
		Title:     template.HTML(metadata["Title"].(string)),
		Email:     template.HTML(metadata["Email"].(string)),
		Content:   template.HTML(data.String()),
		CreatedAt: template.HTML(metadata["CreatedAt"].(string)),
	}

	err := templateHTML.Execute(w, model)
	if err != nil {
		return err
	}

	return nil
}
