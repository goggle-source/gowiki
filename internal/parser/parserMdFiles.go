package parser

import (
	"bufio"
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
	result := make([]HTMLReadyMarkDownFile, 0)
	for _, path := range pathFiles {
		file, err := os.Open(path)
		if err != nil {
			file.Close()
			continue
		}
		scanner := bufio.NewScanner(file)
		count := 0
		flagIsReadOneStr := true
		data := make([]byte, 0)
		for scanner.Scan() {
			text := scanner.Text()
			if flagIsReadOneStr && text != "---" {
				data = append(data, scanner.Bytes()...)
				data = append(data, '\n')
			}
			if text == "---" {
				count += 1
				flagIsReadOneStr = false
				continue
			}
			if count >= 2 {
				data = append(data, scanner.Bytes()...)
				data = append(data, '\n')
			}
		}

		if scanner.Err() != nil {
			file.Close()
			return []HTMLReadyMarkDownFile{}, scanner.Err()
		}
		file.Close()
		var bufes bytes.Buffer
		if err := goldmark.Convert(data, &bufes); err != nil {
			return []HTMLReadyMarkDownFile{}, err
		}

		result = append(result, HTMLReadyMarkDownFile{
			Path: path,
			Data: bufes,
		})
	}

	return result, nil
}

func GetMarkDownFiles(pathDirContent string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(pathDirContent, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && filepath.Ext(d.Name()) == ".md" {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
