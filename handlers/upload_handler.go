package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const (
	maxFileSize = 20 * 1024 * 1024
	fileField   = "image"
	uploadDir   = "./uploads"
)

func UploadHandler(apiKey *string) http.HandlerFunc {
	// Return a function compatible with http.HandlerFunc
	return func(w http.ResponseWriter, r *http.Request) {

		// Parse the multipart form with a 20MB file size limit
		err := r.ParseMultipartForm(maxFileSize)
		if err != nil {
			log.Printf("Error parsing image: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Retrieve the uploaded file from the form
		file, handler, err := r.FormFile(fileField)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		// Create or use the "uploads" directory to store the uploaded files
		if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
			os.Mkdir(uploadDir, os.ModePerm)
		}

		// Create a file path for the uploaded file
		filePath := filepath.Join(uploadDir, handler.Filename)
		if err := saveFile(file, filePath); err != nil {
			log.Printf("Error saving file: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		imageText, err := extractTextFromImage(filePath, *apiKey)
		if err != nil {
			log.Printf("Error extracting text from image: %v", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		data := map[string]interface{}{
			"text": imageText,
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Printf("Error marshaling JSON: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		_, err = w.Write(jsonData)
		if err != nil {
			log.Printf("Error creating response: %v", err)
		}
	}
}

func saveFile(source io.Reader, destination string) error {
	outputFile, err := os.Create(destination)
	if err != nil {
		return fmt.Errorf("couldn't create destination file: %w", err)
	}
	defer outputFile.Close()

	if _, err := io.Copy(outputFile, source); err != nil {
		return fmt.Errorf("couldn't copy source file to destination file: %w", err)
	}
	return nil
}

