# E-book server
E-book server hosts a file server to upload ebooks (i.e. epub, pdf, etc) 
to and download them directly from your (kobo) e-reader.

It also accepts .acsm files, which are Adobe Content Server Manager. 
Currently there is no Linux support from Adobe to add this content to 
your e-reader, which is why knock is used from 
github.com/BentonEdmondson/knock to convert it to a epub.

## Functionalities
Functionalities in a quick overview:
- Upload ebooks to local fileserver
- Convert .acsm to .epub
- Download ebooks (to your e-reader)

## How to run
- Clone/Download this repository.
- In command-line, go to the folder where you've cloned/unpacked the repository.
- Run `go build` to create executable.
- Run the executable (in Linux: `./ebookserver`, in windows `ebookserver.exe`).
- Open the cookbook in your browser at `localhost:4500`.
- You can also use the dockerfile to build an image and run it (with port 4500 exposed).
