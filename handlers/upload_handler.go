package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func UploadHandler(apiKey *string) http.HandlerFunc {
	// Return a function compatible with http.HandlerFunc
	return func(w http.ResponseWriter, r *http.Request) {

		// Parse the multipart form with a 20MB file size limit
		err := r.ParseMultipartForm(20 * 1024 * 1024)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Retrieve the uploaded file from the form
		file, handler, err := r.FormFile("image")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		// Create or use the "uploads" directory to store the uploaded files
		uploadDir := "./uploads"
		if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
			os.Mkdir(uploadDir, os.ModePerm)
		}

		// Create a file path for the uploaded file
		filePath := filepath.Join(uploadDir, handler.Filename)
		if err := saveFile(file, filePath); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response := extractTextFromImage(filePath, *apiKey)
		fmt.Fprintf(w, "Image %s uploaded successfully!\n", handler.Filename)
		fmt.Fprintf(w, "Extracted Text:\n%s", response)
	}
}

func saveFile(source io.Reader, destination string) error {
	outputFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	_, err = io.Copy(outputFile, source)
	return err
}

