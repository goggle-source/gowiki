package commands

import (
	"fmt"

	"github.com/goggle-source/gowiki/internal/config"
	createenv "github.com/goggle-source/gowiki/internal/create_env"
	"github.com/spf13/cobra"
)

func InitCreateEnv() *cobra.Command {
	command := &cobra.Command{
		Use:   "create_env",
		Short: "create base directory from config",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := config.MustLoad()

			ok, err := createenv.IsDirExists(cfg.PathForReadyHTML)
			if err != nil || !ok {
				err := createenv.CreateDirectory(cfg.PathForReadyHTML)
				if err != nil {
					return fmt.Errorf("err create ready HTML: %w", err)
				}
			}

			ok, err = createenv.IsDirExists(cfg.PathMdFile)
			if err != nil || !ok {
				err := createenv.CreateDirectory(cfg.PathMdFile)
				if err != nil {
					return fmt.Errorf("err create md file: %w", err)
				}
			}

			ok, err = createenv.IsDirExists(cfg.PathTemplateHTML)
			if err != nil || !ok {
				err := createenv.CreateDirectory(cfg.PathTemplateHTML)
				if err != nil {
					return fmt.Errorf("err create template: %w", err)
				}
			}

			return nil
		},
	}

	return command
}
