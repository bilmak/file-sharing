package uploads

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const (
	UploadPath = "Downloads/"
)

func UploadForm(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("templates/upload.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	temp.Execute(w, nil)

}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	filePath := filepath.Join(UploadPath, handler.Filename)
	outfile, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Error saving the file", http.StatusInternalServerError)
		return
	}
	defer outfile.Close()

	_, err = io.Copy(outfile, file)
	if err != nil {
		http.Error(w, "Error saving the file", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "File uploaded successfully")
}

func DownloadFile(w http.ResponseWriter, r *http.Request) {
	//http://127.0.0.1:6767/download?file=pexels-vasanth-babu-797797.jpg
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	filename := r.URL.Query().Get("file")
	image, err := os.Open("Downloads/" + filename)
	if err != nil {
		http.Error(w, "Can not open image", http.StatusBadRequest)
		return
	}
	io.Copy(w, image)
}
