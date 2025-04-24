package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Get ?size=25|50|100|500 from query string
	sizeParam := request.QueryStringParameters["size"]
	if sizeParam == "" {
		sizeParam = "500"
	}

	validSizes := map[string]bool{
		"25":  true,
		"50":  true,
		"100": true,
		"500": true,
	}

	if !validSizes[sizeParam] {
		return failResponse("Invalid size parameter. Use one of: 25, 50, 100, 500")
	}

	filename := fmt.Sprintf("pi_input_%s.json", sizeParam)

	numbers, err := loadFromFile(filename)
	if err != nil {
		return failResponse("Failed to load file: " + err.Error())
	}

	linearSort(numbers)

	body := fmt.Sprintf("%v", numbers)
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       body,
		Headers:    map[string]string{"Content-Type": "text/plain"},
	}, nil
}

func loadFromFile(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data struct {
		Numbers []int `json:"numbers"`
	}
	err = json.NewDecoder(file).Decode(&data)
	return data.Numbers, err
}

func linearSort(list []int) {
	i := 0
	for i < len(list)-1 {
		if list[i] > list[i+1] {
			list[i], list[i+1] = list[i+1], list[i]
			i = 0
		} else {
			i++
		}
	}
}

func failResponse(message string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusBadRequest,
		Body:       message,
		Headers:    map[string]string{"Content-Type": "text/plain"},
	}, nil
}
