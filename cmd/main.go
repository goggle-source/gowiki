package main

import (
	"fmt"
	"os"

	"github.com/goggle-source/gowiki/internal/config"
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

	g := parser.Init(cfg.PathTemplateHTML)

	temp := g.GetTemplate("test1.html")

	file, err := os.OpenFile("test2.html", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}

	temp.Execute(file, 0)
}
