package main

import (
	"io"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./assets")))
	http.HandleFunc("/test", handlerMain)
	http.ListenAndServe(":8080", nil)
}

func handlerMain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `hallo<br><a href="./test.pdf">link text</a>`)
}
