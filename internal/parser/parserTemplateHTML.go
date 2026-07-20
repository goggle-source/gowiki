package parser

import (
	"html/template"
)

type GlobalTemplate struct {
	Templates *template.Template
}

func Init(pathDirTemplates string) *GlobalTemplate {

	g := GlobalTemplate{
		Templates: nil,
	}
	g.loadTemplates(pathDirTemplates)

	return &g
}

func (g *GlobalTemplate) loadTemplates(pathDirTemplates string) error {
	templates, err := template.ParseGlob(pathDirTemplates + "/" + "*.html")
	if err != nil {
		return err
	}
	g.Templates = templates
	return nil
}

func (g *GlobalTemplate) GetTemplate(nameTemplate string) *template.Template {
	result := g.Templates.Lookup(nameTemplate)

	return result
}
