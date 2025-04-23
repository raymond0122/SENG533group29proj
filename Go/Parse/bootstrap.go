package main

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	testString := "hellolambdaaws"

	letterCounts := countLetters(testString)

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

func countLetters(s string) [26]int {
	var counts [26]int
	for i := 0; i < len(s); i++ {
		index := s[i] - 'a' // ASCII math: 'a' = 97
		counts[index]++
	}
	return counts
}
