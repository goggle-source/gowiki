package main

import (
	"fmt"

	"github.com/goggle-source/gowiki/internal/config"
	"github.com/goggle-source/gowiki/internal/generator"
	"github.com/goggle-source/gowiki/internal/parser"
)

func main() {
	cfg := config.MustLoad()

	res, err := parser.GetHTMLContent(cfg.PathMdFile)
	if err != nil {
		panic(err)
	}

	for _, value := range res {
		fmt.Println(value.Data.String())
		fmt.Println(value.Path)
		maps, err := parser.GetMetadataMdFile(value.Path)
		if err != nil {
			fmt.Println(err)
		}

		for key, value := range maps {
			fmt.Println("key", key)
			fmt.Println("value", value)
		}
	}

	fmt.Println(res)

	err = generator.GenerateReadyHTML(cfg.PathMdFile, cfg.PathTemplateHTML, cfg.PathForReadyHTML)

	fmt.Println(err)

}
