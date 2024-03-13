package util

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

const (
	apiEndpoint = "https://vision.googleapis.com/v1/images:annotate"
	featureType = "DOCUMENT_TEXT_DETECTION"
)

func ExtractTextFromImage(filePath, apiKey string) (string, error) {
	// Read and Base64 encode the image file
	imageData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read image file: %w", err)
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
						"type": featureType,
					},
				},
			},
		},
	}

	// Convert request payload to JSON
	requestJSON, err := json.Marshal(requestData)
	if err != nil {
		return "", fmt.Errorf("failed to encode JSON: %w", err)
	}

	// Send the HTTP POST request
	apiURL := fmt.Sprintf("%s?key=%s", apiEndpoint, apiKey)
	response, err := http.Post(apiURL, "application/json", bytes.NewBuffer(requestJSON))
	if err != nil {
		return "", fmt.Errorf("failed to send request to API: %w", err)
	}
	defer response.Body.Close()

	// Read and parse the response
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read API response body: %w", err)
	}

	// Extract text from the response
	var responseData map[string]interface{}
	if err := json.Unmarshal(responseBody, &responseData); err != nil {
		return "", fmt.Errorf("failed to decode response JSON: %w", err)
	}

	// Extracted text might be empty if no text was detected
	extractedText := ""
	if annotations, ok := responseData["responses"].([]interface{}); ok && len(annotations) > 0 {
		if annotation, ok := annotations[0].(map[string]interface{}); ok {
			if fullTextAnnotation, ok := annotation["fullTextAnnotation"].(map[string]interface{}); ok {
				if text, ok := fullTextAnnotation["text"].(string); ok {
					extractedText = text
				}
			}
		}
	}

	return extractedText, nil
}

func IsImage(file multipart.File) bool {
	// Read the first 512 bytes to determine the file type
	buffer := make([]byte, 512)
	_, err := file.Read(buffer)
	if err != nil && err != io.EOF {
		return false
	}

	// Reset the file pointer to the beginning
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return false
	}

	// Check for image file types
	fileType := http.DetectContentType(buffer)

	// Check supported image formats
	switch fileType {
	case "image/jpeg", "image/png", "image/gif", "image/bmp", "image/webp", "image/tiff", "image/x-icon":
		return true
	default:
		return false
	}
}

func SaveFile(source io.Reader, destination string) error {
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
