package commands

import (
	"github.com/goggle-source/gowiki/internal/config"
	"github.com/goggle-source/gowiki/internal/handler"
	"github.com/spf13/cobra"
)

func InitServe() *cobra.Command {
	command := &cobra.Command{
		Use:   "serve",
		Short: "start http server",
		RunE: func(cmd *cobra.Command, args []string) error {

			cfg := config.MustLoad()

			server, err := handler.InitServe(cfg.Port, cfg.PathForReadyHTML)

			if err != nil {
				return err
			}

			if err := server.ListenAndServe(); err != nil {
				return err
			}

			return nil
		},
	}

	return command
}
