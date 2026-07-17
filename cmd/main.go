package main

import (
	"fmt"

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
	}
}
