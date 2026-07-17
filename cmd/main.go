package main

import (
	"fmt"

	"github.com/goggle-source/gowiki/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)
}
