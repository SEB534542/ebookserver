package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
	"time"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./assets")))
	// http.HandleFunc("/", handlerMain)
	http.HandleFunc("/upload", handlerUpload)
	http.HandleFunc("/download", handlerDownload)
	srv := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

// knock takes a path to a .acsm file, tries to convert is using the binairy
// knock in path folder. It returns the error from the binairy (if).
func knock(path string) error {
	output, err := exec.Command("./bin/knock", "./assets/test.acsm").Output()
	if err != nil {
		return err
	}
	fmt.Println(string(output))
	return nil
}

func handlerDownload(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `Main branch (works)<br><br><a href="./test.pdf">link text</a>`)
}

func handlerMain(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	http.ServeFile(w, r, "index.html")
}
