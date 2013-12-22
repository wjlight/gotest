package main

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"runtime/debug"
)

const (
	UPLOAD_DIR   = "./uploads"
	TEMPLATE_DIR = "./view"
	ListDir      = 0x0001
)

var templates map[string]*template.Template

func tmplInit() {
	templates = make(map[string]*template.Template)
}

//该方法会先main函数之前被调，自动的
func init() {
	tmplInit()
	fileInfoArr, err := ioutil.ReadDir(TEMPLATE_DIR)

	check(err)

	var templateName, templatePath string
	for _, fileInfo := range fileInfoArr {
		templateName = fileInfo.Name()
		if ext := path.Ext(templateName); ext != ".html" {
			continue
		}
		templatePath = TEMPLATE_DIR + "/" + templateName
		log.Println("Loading template " + templateName + "...")
		t := template.Must(template.ParseFiles(templatePath))
		templates[templateName] = t
	}
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	// s := "<html><form method=\"POST\" action=\"upload\" " +
	// 	"enctype=\"multipart/form-data\">   " +
	// 	" Choose an image to upload: " +
	// 	"<input name=\"image\" type=\"file\"/>" +
	// 	"<input type=\"submit\" value=\"Upload\"/>" +
	// 	"</form></html>"
	if r.Method == "GET" {
		err := renderHtml(w, "upload", nil)
		check(err)
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
	check(err)
	filename := h.Filename
	defer f.Close()
	t, err := os.Create(UPLOAD_DIR + "/" + filename)
	check(err)
	defer t.Close()
	_, err = io.Copy(t, f)
	check(err)
	http.Redirect(w, r, "/view?id="+filename, http.StatusFound)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	fileInfoArr, err := ioutil.ReadDir("./uploads")
	check(err)
	locals := make(map[string]interface{})
	images := []string{}
	for _, fileInfo := range fileInfoArr {
		imageId := fileInfo.Name()
		images = append(images, imageId)
	}

	locals["images"] = images
	err = renderHtml(w, "list", locals)
	check(err)
}

func renderHtml(w http.ResponseWriter, tmpl string, locals map[string]interface{}) error {
	tmpl = tmpl + ".html"
	err := templates[tmpl].Execute(w, locals)
	return err
}

func safeHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e, ok := recover().(error); ok {
				http.Error(w, e.Error(), http.StatusInternalServerError)
				//可以自定义的错误显示
				//logging
				log.Println("WARN: panic in %v - %v", fn, e)
				log.Println(string(debug.Stack()))
			}
		}()
		fn(w, r)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func staticDirHandler(mux *http.ServeMux, prefix string, staticDir string, flag int) {
	mux.HandleFunc(prefix, func(w http.ResponseWriter, r *http.Request) {
		file := staticDir + r.URL.Path[len(prefix)-1:]
		if (flag & ListDir) == 0 {
			if exists := isExists(file); !exists {
				http.NotFound(w, r)
				return
			}
		}
		http.ServeFile(w, r, file)
	})
}

func main() {
	mux := http.NewServeMux()
	staticDirHandler(mux, "/assets", "./public", 0)
	http.HandleFunc("/", safeHandler(listHandler))
	http.HandleFunc("/view", safeHandler(viewHandler))
	http.HandleFunc("/upload", safeHandler(uploadHandler))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
}
