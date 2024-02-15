package handlers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

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

