package handlers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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
		outputFile, err := os.Create(filePath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer outputFile.Close()

		// Copy the contents of the uploaded file to the new file
		_, err = io.Copy(outputFile, file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response := extractTextFromImage(filePath, *apiKey)
		fmt.Fprintf(w, "Image %s uploaded successfully!\n", handler.Filename)
		fmt.Fprintf(w, "Extracted Text:\n%s", response)
	}
}

func extractTextFromImage(filePath string, apiKey string) string{

	// Read and Base64 encode the image file
	imageData, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading image file:", err)
		// return
	}
	encodedImage := base64.StdEncoding.EncodeToString(imageData)

	// Construct the request payload
	requestData := map[string]interface{}{
		"requests": []map[string]interface{}{
			{
				"image": map[string]string{
					"content": encodedImage,
				},
				"features": []map[string]string{
					{
						"type": "DOCUMENT_TEXT_DETECTION",
					},
				},
			},
		},
	}

	// Convert request payload to JSON
	requestJSON, err := json.Marshal(requestData)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		// return
	}

	// Send the HTTP POST request
	apiEndpoint := fmt.Sprintf("https://vision.googleapis.com/v1/images:annotate?key=%s", apiKey)
	response, err := http.Post(apiEndpoint, "application/json", bytes.NewBuffer(requestJSON))
	if err != nil {
		fmt.Println("Error sending request:", err)
		// return
	}
	defer response.Body.Close()

	// Read and parse the response
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		// return
	}

	// Check if the request was successful
	if response.StatusCode != http.StatusOK {
		fmt.Println("Error:", string(responseBody))
		// return
	}

	// Extract text from the response
	var responseData map[string]interface{}
	if err := json.Unmarshal(responseBody, &responseData); err != nil {
		fmt.Println("Error decoding response JSON:", err)
		// return
	}

	// Print the extracted text
	extractedText := responseData["responses"].([]interface{})[0].(map[string]interface{})["fullTextAnnotation"].(map[string]interface{})["text"].(string)

	return extractedText
}

