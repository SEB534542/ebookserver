package main

import (
	"fmt"
	"io"
	"net/http"
	"os/exec"
)

func main() {
	output, err := exec.Command("./bin/knock", "./assets/test.acsm").Output()
	if err != nil {
		fmt.Println("Error", err.Error())
	}
	fmt.Println(string(output))

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
