package handler

import (
	"net/http"
	"strconv"
	"time"
)

func InitServe(port int, pathForReadyHTML string) (*http.Server, error) {
	strPort := strconv.Itoa(port)
	addr := ":" + strPort
	mux := http.NewServeMux()

	mux.HandleFunc("/files/", fileHandler)

	ser := &http.Server{
		Addr:         addr,
		Handler:      mux,
		WriteTimeout: 20 * time.Second,
		IdleTimeout:  40 * time.Second,
	}

	return ser, nil
}
