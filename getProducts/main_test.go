package main

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandleRequest(t *testing.T) {
	testProduct := Product{
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

	request := events.APIGatewayProxyRequest{}

	ctx := context.Background()

	response, err := handleRequest(ctx, request)

	if err != nil {
		t.Fatalf("Error occurred: %v", err)
	}

	if response.StatusCode != 200 {
		t.Errorf("Expected status code 200, but got %d", response.StatusCode)
	}

	expectedBody, _ := json.Marshal(testProduct)

	if response.Body != string(expectedBody) {
		t.Errorf("Expected body %s, but got %s", string(expectedBody), response.Body)
	}
}