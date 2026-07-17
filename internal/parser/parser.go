package parser

import (
	"bytes"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/yuin/goldmark"
)

type HTMLReadyMarkDownFile struct {
	Data bytes.Buffer
	Path string
}

func GetHTMLContent(pathDirContent string) ([]HTMLReadyMarkDownFile, error) {
	pathFiles, err := GetMarkDownFiles(pathDirContent)
	if err != nil {
		return []HTMLReadyMarkDownFile{}, err
	}
	result := make([]HTMLReadyMarkDownFile, 1)
	for _, path := range pathFiles {
		data, err := os.ReadFile(path)
		if err != nil {
			continue
		}
		var byte bytes.Buffer

		if err := goldmark.Convert(data, &byte); err != nil {
			continue
		}
		result = append(result, HTMLReadyMarkDownFile{
			Data: byte,
			Path: path,
		})
	}

	return result, nil
}

func GetMarkDownFiles(pathDirContent string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(pathDirContent, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if !d.IsDir() && filepath.Ext(d.Name()) == ".md" {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
