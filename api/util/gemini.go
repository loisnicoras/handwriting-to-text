package util

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"cloud.google.com/go/vertexai/genai"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

type SafetyRating struct {
	Category    int  `json:"Category"`
	Probability int  `json:"Probability"`
	Blocked     bool `json:"Blocked"`
}

type Candidate struct {
	Index   int `json:"Index"`
	Content struct {
		Role  string   `json:"Role"`
		Parts []string `json:"Parts"`
	} `json:"Content"`
	FinishReason     int            `json:"FinishReason"`
	SafetyRatings    []SafetyRating `json:"SafetyRatings"`
	FinishMessage    string         `json:"FinishMessage"`
	CitationMetadata interface{}    `json:"CitationMetadata"`
}

func CalculateScore(correctText, genText, projectId, region string) (int, error) {
	ctx := context.Background()
	// f, err := os.Open("moonlit-shadow-325207-72e8674d169e.json")
	// if err != nil {
	// 	log.Fatalf("unable to read file: %v", err)
	// }
	// buf := make([]byte, 1024)
	// fmt.Print(f.Read(buf))
	jsonKey, err := ioutil.ReadFile("../api/moonlit-shadow-325207-72e8674d169e.json")
	if err != nil {
		return 0, fmt.Errorf("Failed to read JSON: %w", err)
	}
	creds, err := google.CredentialsFromJSON(ctx, jsonKey, "https://www.googleapis.com/auth/cloud-platform")
	if err != nil {
		return 0, fmt.Errorf("Failed get credential from JSON: %w", err)
	}

	client, err := genai.NewClient(ctx, projectId, region, option.WithCredentials(creds))
	if err != nil {
		return 0, fmt.Errorf("Failed create new client: %w", err)
	}
	gemini := client.GenerativeModel("gemini-pro-vision")

	prompt := genai.Text("Can you give me a score (just the score no more words) that is between 0-100 from comparing the first text with the second. First is the correct text. The correct text is: " + correctText + " The incorrect text is " + genText)
	resp, err := gemini.GenerateContent(ctx, prompt)
	if err != nil {
		return 0, fmt.Errorf("Failed to generate content: %w", err)
	}
	rb, _ := json.MarshalIndent(resp, "", "  ")

	type Response struct {
		Candidates []Candidate `json:"Candidates"`
	}

	// Unmarshal the JSON response string into the Response struct
	var response Response
	err = json.Unmarshal([]byte(rb), &response)
	if err != nil {
		return 0, fmt.Errorf("Failed to unmarshal the json: %w", err)
	}

	// Access the "Parts" data from the first candidate
	parts := response.Candidates[0].Content.Parts

	// Convert the string value to float64
	floatValue, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return 0, fmt.Errorf("Failed to parse float value: %w", err)
	}

	// Convert float64 to integer
	intValue := int(floatValue)

	return intValue, nil
}
