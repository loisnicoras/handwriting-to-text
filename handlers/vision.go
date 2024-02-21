package handlers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	apiEndpoint = "https://vision.googleapis.com/v1/images:annotate"
	featureType = "DOCUMENT_TEXT_DETECTION"
)

func extractTextFromImage(filePath, apiKey string) (string, error) {
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
						"type": featureType,
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
	apiURL := fmt.Sprintf("%s?key=%s", apiEndpoint, apiKey)
	response, err := http.Post(apiURL, "application/json", bytes.NewBuffer(requestJSON))
	if err != nil {
		return "", fmt.Errorf("error sending request to API: %w", err)
	}
	defer response.Body.Close()

	// Read and parse the response
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("error reading API response body: %w", err)
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

