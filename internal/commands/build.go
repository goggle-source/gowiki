package commands

import (
	"github.com/goggle-source/gowiki/internal/config"
	"github.com/goggle-source/gowiki/internal/generator"
	"github.com/goggle-source/gowiki/internal/logger"
	"github.com/spf13/cobra"
)

func InitBuild() *cobra.Command {
	command := &cobra.Command{
		Use:   "build",
		Short: "assemble a ready-made HTML file",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := config.MustLoad()
			log := logger.InitLogger()
			err := generator.GenerateReadyHTML(cfg.PathMdFile, cfg.PathTemplateHTML, cfg.PathForReadyHTML, log)
			return err
		},
	}

	return command
}
