package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

const (
	folderAssets = "./assets" // Folder for up- and downloading.
	knockExt     = ".acsm"    // Extension for which knock function is required
	port         = ":4500"    // Port to access app
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(folderAssets)))
	http.HandleFunc("/books", handlerBooks)
	http.HandleFunc("/upload", handlerUpload)
	http.HandleFunc("/delete/", handlerDelete)
	srv := &http.Server{
		Addr:         port,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

// knock takes a path to a .acsm file, tries to convert it using the binairy
// file 'knock' in path folder. It returns the error from the binairy (if).
func knock(path string) error {
	output, err := exec.Command("./bin/knock", path).Output()
	if err != nil {
		return err
	}
	log.Println(string(output))
	return nil
}

// handlerbook lists the books on the server.
func handlerBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	s := "<h1>Download books</h1><table cellpadding='5'>"
	for _, fname := range files(folderAssets) {
		if fname != "index.html" {
			s += fmt.Sprintf("<tr><td><a href='./%s'>%s</a></td><td>&emsp;&emsp;&emsp;<i><a href='./delete/%s'>delete</a></i></td></tr>", fname, fname, fname)
		}
	}
	s += "</table><p><a href='/'>Click here to go back</a></p>"
	io.WriteString(w, s)
}

// files takes a path to a directory and returns a slice containing all files in that directory.
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

// handlerDelete deletes the specified file from the server.
func handlerDelete(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Path[len("/delete/"):]
	if file == "" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	err := os.Remove(fmt.Sprintf("%s/%s", folderAssets, file))
	if err != nil {
		log.Printf("Unable to delete file '%v': %v", file, err)
		http.Error(w, fmt.Sprint("Unable to delete file", file), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, "File succesfully deleted<br><br><a href='/books'>Click here to go back</a>")
}
