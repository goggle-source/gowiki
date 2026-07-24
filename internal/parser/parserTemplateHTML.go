package parser

import (
	"html/template"
	"path/filepath"
)

type GlobalTemplate struct {
	Templates *template.Template
}

func Init(pathDirTemplates string) *GlobalTemplate {

	g := GlobalTemplate{
		Templates: loadTemplates(pathDirTemplates),
	}

	return &g
}

func loadTemplates(pathDirTemplates string) *template.Template {
	pattern := filepath.Join(pathDirTemplates, "*.html")
	templates, err := template.ParseGlob(pattern)
	if err != nil {
		panic(err)
	}
	return templates
}

func (g *GlobalTemplate) GetTemplate(nameTemplate string) *template.Template {
	if g.Templates != nil {
		result := g.Templates.Lookup(nameTemplate)
		return result
	}
	return nil
}
