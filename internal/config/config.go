package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Cfg struct {
	PathMdFile       string `yaml:"path_for_md_file"`
	PathForReadyHTML string `yaml:"path_for_ready_html_file"`
	PathTemplateHTML string `yaml:"path_for_template_html"`
	Port             int    `yaml:"port"`
}

func MustLoad() *Cfg {
	data, err := os.ReadFile("./config.yml")
	if err != nil {
		panic(err)
	}
	var cfg Cfg
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		panic(err)
	}

	return &cfg
}
