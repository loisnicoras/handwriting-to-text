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

// type Response struct {
// 	Status  int         `json:"status"`
// 	Message string      `json:"message"`
// 	Data    interface{} `json:"data"`
// }

// func createResponse(status int, message string, data interface{}) []byte {
// 	resp := Response{
// 		Status:  status,
// 		Message: message,
// 		Data:    data,
// 	}

// 	jsonData, _ := json.Marshal(resp)

// 	return jsonData
// }

func UploadHandler(apiKey *string) http.HandlerFunc {
	// Return a function compatible with http.HandlerFunc
	return func(w http.ResponseWriter, r *http.Request) {

		// Parse the multipart form with a 20MB file size limit
		err := r.ParseMultipartForm(20 * 1024 * 1024)
		if err != nil {
			fmt.Println("couldn't parse image: %w", err)
			w.WriteHeader(http.StatusInternalServerError)
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
			fmt.Println("couldn't save file: %w", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		imageText, err := extractTextFromImage(filePath, *apiKey)
		if err != nil {
			fmt.Println("couldn't exctract text from image: %w", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		data := map[string]interface{}{
			"text":    imageText,
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			fmt.Printf("could not marshal json: %s\n", err)
			return
		}
	
		_, err = w.Write(jsonData)
		if err != nil {
			fmt.Println("error create response %w", err)
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

func extractTextFromImage(filePath string, apiKey string) (string, error) {
	// Read and Base64 encode the image file
	imageData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("error reading image file: %w", err)
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
		return "", fmt.Errorf("error encoding JSON: %w", err)
	}

	// Send the HTTP POST request
	apiEndpoint := fmt.Sprintf("https://vision.googleapis.com/v1/images:annotate?key=%s", apiKey)
	response, err := http.Post(apiEndpoint, "application/json", bytes.NewBuffer(requestJSON))
	if err != nil {
		return "", fmt.Errorf("error sending request to api: %w", err)
	}
	defer response.Body.Close()

	// Read and parse the response
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("error reading api response body: %w", err)
	}

	// Extract text from the response
	var responseData map[string]interface{}
	if err := json.Unmarshal(responseBody, &responseData); err != nil {
		return "", fmt.Errorf("error decoding response JSON: %w", err)
	}

	// return the extracted text
	extractedText := responseData["responses"].([]interface{})[0].(map[string]interface{})["fullTextAnnotation"].(map[string]interface{})["text"].(string)

	return extractedText, nil
}
