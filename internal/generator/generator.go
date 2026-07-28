package generator

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/goggle-source/gowiki/internal/parser"
)

func GenerateReadyHTML(pathDirContent, pathDirTemplates, pathDirForReadyHTML string, log *slog.Logger) error {
	mdFiles, err := parser.GetHTMLContent(pathDirContent)
	if err != nil {
		return err
	}

	for _, file := range mdFiles {
		metadata, err := parser.GetMetadataMdFile(file.Path)
		if err != nil {
			log.Info("err get metadataMdFile", slog.Any("err", err))
			continue
		}
		if metadata["name_template"].(string) == "" {
			log.Info("name_template is not found")
			continue
		}
		if metadata["name_ready_html"].(string) == "" {
			log.Info("name_ready_html is not found")
			continue
		}

		ServicTemplate := parser.Init(pathDirTemplates)
		if ServicTemplate.Templates == nil {
			log.Error("err init ServicTemplate")
			return fmt.Errorf("app error")
		}
		template := ServicTemplate.GetTemplate(metadata["name_template"].(string))
		path := filepath.Join(pathDirForReadyHTML, metadata["name_ready_html"].(string))
		fileForReadyHTML, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0773)
		if err != nil {
			panic(fmt.Errorf("%s:%w", "err openfile", err))
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

	metadata["Content"] = template.HTML(data.String())

	for key, value := range metadata {
		if _, ok := value.(string); !ok {
			continue
		}
		metadata[key] = template.HTML(value.(string))
	}

	err := templateHTML.Execute(w, metadata)
	if err != nil {
		return err
	}

	return nil
}
