package http

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func Download(store StoreService, fs FileService) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		//First of check if Get is set in the URL
		Filename := request.URL.Query().Get("file")
		if Filename == "" {
			//Get not set, send a 400 bad request
			http.Error(writer, "Get 'file' not specified in url.", 400)
			return
		}
		fmt.Println("Client requests: " + Filename)

		//Check if file exists and open
		Openfile, err := os.Open("main.go")
		defer Openfile.Close() //Close after function return
		if err != nil {
			//File not found, send 404
			http.Error(writer, "File not found.", 404)
			return
		}

		//File is found, create and send the correct headers

		//Get the Content-Type of the file
		//Create a buffer to store the header of the file in
		FileHeader := make([]byte, 512)
		//Copy the headers into the FileHeader buffer
		Openfile.Read(FileHeader)
		//Get content type of file
		FileContentType := http.DetectContentType(FileHeader)

		//Get the file size
		FileStat, _ := Openfile.Stat()                     //Get info from file
		FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string

		//Send the headers
		writer.Header().Set("Content-Disposition", "attachment; filename="+Filename)
		writer.Header().Set("Content-Type", FileContentType)
		writer.Header().Set("Content-Length", FileSize)

		//Send the file
		//We read 512 bytes from the file already so we reset the offset back to 0
		Openfile.Seek(0, 0)
		io.Copy(writer, Openfile) //'Copy' the file to the client
		return
	}
}

func Downloadv2(w http.ResponseWriter, r *http.Request) {
	//copy the relevant headers. If you want to preserve the downloaded file name, extract it with go's url parser.
	w.Header().Set("Content-Disposition", "attachment; filename=aFile")
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", r.Header.Get("Content-Length"))

	file, _, err := r.FormFile("uploadfile")
	defer file.Close()

	//stream the body to the client without fully loading it into memory
	io.Copy(w, file)

	//Check if file exists and open
	Openfile, err := os.Open("main.go")
	defer Openfile.Close() //Close after function return

	if err != nil {
		//File not found, send 404
		http.Error(w, "File not found.", 404)
		return
	}

	//File is found, create and send the correct headers

	//Get the Content-Type of the file
	//Create a buffer to store the header of the file in
	FileHeader := make([]byte, 512)
	//Copy the headers into the FileHeader buffer
	Openfile.Read(FileHeader)
	//Get content type of file
	FileContentType := http.DetectContentType(FileHeader)

	//Get the file size
	FileStat, _ := Openfile.Stat()                     //Get info from file
	FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string

	//Send the headers
	w.Header().Set("Content-Disposition", "attachment; filename="+"main.go")
	w.Header().Set("Content-Type", FileContentType)
	w.Header().Set("Content-Length", FileSize)

	//Send the file
	//We read 512 bytes from the file already so we reset the offset back to 0
	Openfile.Seek(0, 0)
	io.Copy(w, Openfile) //'Copy' the file to the client
	return
}
