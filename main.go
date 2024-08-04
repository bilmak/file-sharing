package main

import (
	"exchangingFiles/uploads"
	"fmt"
	"net/http"
	"os"
)

func main() {

	err := os.MkdirAll(uploads.UploadPath, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}
	fs := http.FileServer(http.Dir(uploads.UploadPath))
	http.Handle("/uploads", http.StripPrefix("/uploads", fs))

	http.HandleFunc("/", uploads.UploadForm)
	http.HandleFunc("/upload", uploads.UploadFile)
	http.HandleFunc("/download", uploads.DownloadFile)

	http.ListenAndServe("0.0.0.0:6767", nil)

}
