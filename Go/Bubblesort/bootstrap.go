package main

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	numbers := []int{42, 17, 93, 8, 65, 23, 71, 4, 39, 56}
	linearSort(numbers)

	body := fmt.Sprintf("%v", numbers)
	response := events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       body,
		Headers:    map[string]string{"Content-Type": "text/plain"},
	}

	return response, nil
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
