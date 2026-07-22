package handler

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
)

const baseDir = "./public"

func fileHandler(w http.ResponseWriter, r *http.Request) {

	filename := strings.TrimPrefix(r.URL.Path, "/files/")
	if filename == "" {
		http.Error(w, "missing filename", http.StatusBadRequest)
		return
	}
	fmt.Println(filename)

	cleanName := filepath.Clean(filename)

	fullPath := filepath.Join(baseDir, cleanName)

	if !strings.HasPrefix(fullPath, filepath.Clean(baseDir)+string(filepath.Separator)) &&
		fullPath != filepath.Clean(baseDir) {
		http.Error(w, "forbidden path", http.StatusForbidden)
		return
	}

	fmt.Println(fullPath)
	http.ServeFile(w, r, fullPath)
}
