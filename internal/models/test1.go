package models

import (
	"html/template"
)

type Test1 struct {
	Title     template.HTML
	Name      template.HTML
	Email     template.HTML
	Content   template.HTML
	CreatedAt template.HTML
}
