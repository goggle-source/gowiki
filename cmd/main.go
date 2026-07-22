package main

import (
	"fmt"

	"github.com/goggle-source/gowiki/internal/config"
	"github.com/goggle-source/gowiki/internal/generator"
	"github.com/goggle-source/gowiki/internal/handler"
)

func main() {
	cfg := config.MustLoad()

	err := generator.GenerateReadyHTML(cfg.PathMdFile, cfg.PathTemplateHTML, cfg.PathForReadyHTML)

	fmt.Println(err)

	ser, err := handler.InitServe(cfg.Port, cfg.PathForReadyHTML)
	if err != nil {
		panic(err)
	}

	if err := ser.ListenAndServe(); err != nil {
		panic(err)
	}
}
