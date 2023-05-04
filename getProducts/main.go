package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Variant struct {
	Size  string  `json:"size"`
}

type Product struct {
	Name     string    `json:"name"`
	Price    float64   `json:"price"`
	Variants []Variant `json:"variants"`
}

var product = Product{
    Name: "mschfx",
	Price: 1010.10,
    Variants: []Variant{
        {
            Size: "s/m",
        },
        {
            Size: "m/l",
        },
    },
}


func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// encode the product struct into JSON
	responseBody, err := json.Marshal(product)
	if err != nil {
		log.Fatal(err)
		return events.APIGatewayProxyResponse{}, err
	}

	// create a response object with the JSON-encoded product struct as the body
	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(responseBody),
	}

	return response, nil
}

func main() {
	lambda.Start(handleRequest)
}