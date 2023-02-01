package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"time"
)

const (
	folderAssets = "./assets" // Folder for up- and downloading.
	knockExt     = ".acsm"    // Extension for which knock function is required
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(folderAssets)))
	http.HandleFunc("/books", handlerBooks)
	http.HandleFunc("/upload", handlerUpload)
	srv := &http.Server{
		Addr:         ":4500",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

// knock takes a path to a .acsm file, tries to convert is using the binairy
// knock in path folder. It returns the error from the binairy (if).
func knock(path string) error {
	output, err := exec.Command("./bin/knock", path).Output()
	if err != nil {
		return err
	}
	log.Println(string(output))
	return nil
}

func handlerBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	s := "<h1>Download books</h1>"
	for _, fname := range files(folderAssets) {
		if fname != "index.html" {
			s += fmt.Sprintf("<a href='./%s'>%s</a><br><br>", fname, fname)
		}
	}
	io.WriteString(w, s)
}

func files(path string) []string {
	files, err := ioutil.ReadDir(folderAssets)
	if err != nil {
		return []string{}
	}
	output := make([]string, len(files))
	for i, file := range files {
		output[i] = file.Name()
	}
	return output
}
