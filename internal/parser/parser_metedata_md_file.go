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
	count_str := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "---" {
			count_str += 1
			continue
		}
		if count_str == 2 {
			break
		}
		if count_str >= 1 && count_str < 2 {
			strs := strings.Split(text, ":")
			result[strs[0]] = strs[1]
		}
	}
	if scanner.Err() != nil {
		return map[string]any{}, scanner.Err()
	}

	return result, nil
}
