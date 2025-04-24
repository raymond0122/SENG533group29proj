package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Hardcoded file
	filename := "string_input_25.json"

	str, err := loadStringFromFile(filename)
	if err != nil {
		return failResponse("Failed to load string: " + err.Error())
	}

	letterCounts := countLetters(str)

	// Return result as JSON
	respJSON, err := json.Marshal(letterCounts)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Failed to marshal result",
			Headers:    map[string]string{"Content-Type": "text/plain"},
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(respJSON),
		Headers:    map[string]string{"Content-Type": "application/json"},
	}, nil
}

func loadStringFromFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var data struct {
		Data string `json:"data"`
	}
	err = json.NewDecoder(file).Decode(&data)
	return data.Data, err
}

func countLetters(s string) [26]int {
	var counts [26]int
	for i := 0; i < len(s); i++ {
		index := s[i] - 'a'
		if index >= 0 && index < 26 {
			counts[index]++
		}
	}
	return counts
}

func failResponse(message string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusBadRequest,
		Body:       message,
		Headers:    map[string]string{"Content-Type": "text/plain"},
	}, nil
}
