package upload

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	util "github.com/loisnicoras/handwriting-to-text/util"
)

const (
	maxFileSize = 20 * 1024 * 1024
	fileField   = "image"
	uploadDir   = "./uploads"
)

func UploadHandler(apiKey *string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		err := r.ParseMultipartForm(maxFileSize)
		if err != nil {
			log.Printf("Error parsing image: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		file, handler, err := r.FormFile(fileField)
		if err != nil {
			fmt.Println("Error retrieving file:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		if !util.IsImage(file) {
			http.Error(w, "", http.StatusUnsupportedMediaType)
			return
		}

		if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
			err := os.Mkdir(uploadDir, os.ModePerm)
			if err != nil {
				log.Printf("Error creating 'uploads' directory: %v", err)
				http.Error(w, "Error creating 'uploads' directory:", http.StatusInternalServerError)
				return
			}
		}

		filePath := filepath.Join(uploadDir, handler.Filename)
		if err := util.SaveFile(file, filePath); err != nil {
			log.Printf("Error saving file: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		imageText, err := util.ExtractTextFromImage(filePath, *apiKey)
		if err != nil {
			log.Printf("Error extracting text from image: %v", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		
		responseData := map[string]interface{}{
			"text": imageText,
		}
		
		jsonResponse, err := json.Marshal(responseData)
		if err != nil {
			log.Printf("Error marshaling JSON: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(jsonResponse)
		if err != nil {
			log.Printf("Error creating response: %v", err)
		}
	}
}
