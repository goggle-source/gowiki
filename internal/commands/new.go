package commands

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/goggle-source/gowiki/internal/config"
	"github.com/spf13/cobra"
)

func InitNew() *cobra.Command {
	command := &cobra.Command{
		Use:   "new",
		Short: "copy md file or html file in directory project",
		RunE: func(cmd *cobra.Command, args []string) error {

			path, err := cmd.Flags().GetString("path")
			if err != nil {
				return err
			}
			if path == "" {
				return fmt.Errorf("path is required")
			}
			var baseDir string
			expansion := filepath.Ext(path)

			if expansion == ".md" {
				baseDir = config.MustLoad().PathMdFile
			}
			if expansion == ".html" {
				baseDir = config.MustLoad().PathTemplateHTML
			}
			if baseDir == "" {
				return fmt.Errorf("incorrect file extension: either .md or .html")
			}
			file, err := os.Open(path)
			if err != nil {
				return err
			}

			pathForCopyFile := filepath.Join(baseDir, filepath.Base(file.Name()))
			fmt.Println(pathForCopyFile)
			dstFile, err := os.Create(pathForCopyFile)
			if err != nil {
				return err
			}

			_, err = io.Copy(dstFile, file)
			if err != nil {
				return err
			}

			return nil
		},
	}
	command.Flags().String("path", "", "the full path to the file")
	return command
}
