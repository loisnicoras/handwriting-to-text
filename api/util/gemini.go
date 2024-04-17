package util

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

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

	corText := genai.Text(correctText)
	incorrectText := genai.Text(genText)

	words := strings.Fields(string(correctText))
	numWords := len(words)

	scoreForWrongWord := ((1 * 100) / numWords) / 2
	
	scoreForOneWord := genai.Text(strconv.Itoa((1 * 100) / numWords))
	wrongWord := genai.Text(strconv.Itoa(scoreForWrongWord))

	prompt := genai.Text("Could you provide a score (the final score, just score without any words) to assess the similarity between two texts, ranging from 0 to 100? 1. Calculate the points deducted for missing words in the second text that exist in the first text. 2. Calculate the points deducted for extra words in the second text that does not exist in the first text. 3. Calculate the points deducted for missing or extra letters in each word of the second text you need to calculate the points deducted for missing or extra letters in each word of the second text. 4. Sum up all the deducted points to get the total points deducted. And give me an integer number. 5. Subtract the total points deducted from 100 to get the final score. All points deducted will be integers. The scoring should consider the following criteria: 1. Compare each word from the second text with the corresponding word from the first text. 2. Deduct " + wrongWord + " points from the total score for missing/extra letters in each word of the second text compared to the first. 3. Deduct " + scoreForOneWord + " points from the total score for missing words in the second text that exist in the first text. 4. Deduct " + scoreForOneWord + " points from the total score for extra words in the second text. The correct text is: `" + corText + "`. The incorrect text is: `" + incorrectText + "`. Give me just the final score")
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
