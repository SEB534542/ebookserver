package main

/* file-upload.go is based on the repository from
https://github.com/Freshman-tech/file-upload
and is used under the MIT license with copyright 2020 to Ayooluwa Isaiah*/

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const maxUploadSize = 1024 * 1024 * 20 // 1024 * 1024 = 1 MB

// allowedFtypes contains the filetypes allowed for upload.
var allowedFtypes = []string{
	"application/pdf",
	"application/epub+zip",
	"application/zip",
	"text/plain; charset=utf-8",
}

// Progress is used to track the progress of a file upload.
// It implements the io.Writer interface so it can be passed
// to an io.TeeReader()
type Progress struct {
	TotalSize int64
	BytesRead int64
}

// Write is used to satisfy the io.Writer interface.
// Instead of writing somewhere, it simply aggregates
// the total bytes on each read
func (pr *Progress) Write(p []byte) (n int, err error) {
	n, err = len(p), nil
	pr.BytesRead += int64(n)
	pr.Print()
	return
}

// Print displays the current progress of the file upload
func (pr *Progress) Print() {
	if pr.BytesRead == pr.TotalSize {
		log.Println("DONE!")
		return
	}
	log.Printf("File upload in progress: %d\n", pr.BytesRead)
}

func handlerUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// 32 MB is the default used by FormFile
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// get a reference to the fileHeaders
	files := r.MultipartForm.File["file"]
	for _, fileHeader := range files {
		if fileHeader.Size > maxUploadSize {
			http.Error(w, fmt.Sprintf("The uploaded file is too big: %s. Please use a file less than %sMB in size", fileHeader.Filename, maxUploadSize/1024/1024), http.StatusBadRequest)
			return
		}
		file, err := fileHeader.Open()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()
		buff := make([]byte, 512)
		_, err = file.Read(buff)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		filetype := http.DetectContentType(buff)
		var allowed bool
		for _, v := range allowedFtypes {
			if filetype == v {
				allowed = true
			}
		}
		if !allowed {
			http.Error(w, fmt.Sprintf("The provided file format %v is not allowed", filetype), http.StatusBadRequest)
			return
		}
		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = os.MkdirAll(folderAssets, os.ModePerm)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		path := fmt.Sprintf("%s/%s", folderAssets, fileHeader.Filename)
		f, err := os.Create(path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer f.Close()
		pr := &Progress{
			TotalSize: fileHeader.Size,
		}
		_, err = io.Copy(f, io.TeeReader(file, pr))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if filepath.Ext(fileHeader.Filename) == knockExt {
			err := knock(path)
			if err != nil {
				log.Printf("Unable to convert .acsm file to epub (%v)", err)
			}
		}
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, "Upload successful<br><br><a href='./books'>Click here to list all books available for download</a>")
	// fmt.Fprintf(w, "Upload successful")
}
