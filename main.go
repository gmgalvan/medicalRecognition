package main

import (
	"fmt"
	"io/ioutil"
	"medicalRecognition/awsapi"
	"medicalRecognition/gcpapi"
	"net/http"
	"time"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)

	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	textFromImage := gcpapi.TextInPicture(fileBytes)
	drug := awsapi.DetectEntities(textFromImage)
	path := drug.Host + ".html"
	http.Redirect(w, r, path, 301)
}

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./public"))
	mux.Handle("/", fs)
	mux.HandleFunc("/crazyai", uploadFile)
	srv := http.Server{
		Addr:           ":8081",
		Handler:        mux,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 8175,
	}

	fmt.Println("Listen on port " + srv.Addr)
	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
