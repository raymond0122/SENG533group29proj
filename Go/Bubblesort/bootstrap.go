package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type JSONRequest struct {
	Numbers []int `json:"numbers"`
}

func main() {
	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	inputType := request.QueryStringParameters["input_type"] // "json" or "txt"
	var numbers []int

	switch inputType {
	case "json":
		var jsonReq JSONRequest
		err := json.Unmarshal([]byte(request.Body), &jsonReq)
		if err != nil {
			return failResponse("Invalid JSON input")
		}
		numbers = jsonReq.Numbers

	case "txt":
		decoded, err := base64.StdEncoding.DecodeString(request.Body)
		if err != nil {
			return failResponse("Invalid base64 .txt content")
		}
		parts := strings.Split(string(decoded), ",")
		for _, part := range parts {
			n, err := strconv.Atoi(strings.TrimSpace(part))
			if err != nil {
				return failResponse("Invalid number in .txt content")
			}
			numbers = append(numbers, n)
		}

	default:
		return failResponse("Invalid input_type. Use 'json' or 'txt'")
	}

	linearSort(numbers)

	body := fmt.Sprintf("%v", numbers)
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       body,
		Headers:    map[string]string{"Content-Type": "text/plain"},
	}, nil
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
