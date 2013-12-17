package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

const (
	UPLOAD_DIR = "./uploads"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	s := "<html><form method=\"POST\" action=\"upload\" " +
		"enctype=\"multipart/form-data\">   " +
		" Choose an image to upload: " +
		"<input name=\"image\" type=\"file\"/>" +
		"<input type=\"submit\" value=\"Upload\"/>" +
		"</form></html>"
	if r.Method == "GET" {
		io.WriteString(w, s)
		return
	}

	if r.Method == "POST" {
		handlePost(w, r)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	imagePath := UPLOAD_DIR + "/" + imageId
	if exists := isExists(imagePath); !exists {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-type", "image")
	http.ServeFile(w, r, imagePath)
}

func isExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	f, h, err := r.FormFile("image")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal("image error")
		return
	}
	filename := h.Filename
	defer f.Close()
	t, err := os.Create(UPLOAD_DIR + "/" + filename)
	if err != nil {
		http.Error(
			w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer t.Close()
	if _, err := io.Copy(t, f); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view?id="+filename, http.StatusFound)
}

func main() {
	http.HandleFunc("/view", viewHandler)
	http.HandleFunc("/upload", uploadHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
}
