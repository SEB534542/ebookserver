package main

import (
	"fmt"
	"io"
	"net/http"
	"os/exec"
)

func knock(path string) error {
	output, err := exec.Command("./bin/knock", "./assets/test.acsm").Output()
	if err != nil {
		return fmt.Sprint("ERROR:", err.Error())
	}
	return (string(output))
}

func server() {
	http.Handle("/", http.FileServer(http.Dir("./assets")))
	http.HandleFunc("/download", handlerMain)
	http.ListenAndServe(":8080", nil)
}

func handlerMain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `Main branch (works)<br><br><a href="./test.pdf">link text</a>`)
}
