package parser

import (
	"bufio"
	"os"
	"strings"
)

func GetMetadataMdFile(pathForFile string) (map[string]any, error) {
	file, err := os.Open(pathForFile)
	if err != nil {
		return map[string]any{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := make(map[string]any)
	isTextEqualStart := true
	count_str := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text != "---" && isTextEqualStart {
			break
		}

		if text == "---" {
			count_str += 1
			if count_str == 2 {
				break
			}
			isTextEqualStart = false
			continue
		}

		strs := strings.Split(text, ":")
		if len(strs) == 1 {
			continue
		}
		resStrs := strings.TrimSpace(strs[1])
		keyStrs := strings.TrimSpace(strs[0])
		result[keyStrs] = resStrs
	}
	if scanner.Err() != nil {
		return map[string]any{}, scanner.Err()
	}

	return result, nil
}
