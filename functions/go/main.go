package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
)

type InputEvent struct {
	Name string `json:"name"`
}

type OutputEvent struct {
	Message string `json:"message"`
}

func handler(request InputEvent) (OutputEvent, error) {
	resp, err := http.Get("https://checkip.amazonaws.com")
	if err != nil {
		return OutputEvent{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return OutputEvent{}, err
	}
	fmt.Println(string(body))

	response := OutputEvent{
		Message: fmt.Sprintf("hello %s", request.Name),
	}
	return response, nil
}

func main() {
	lambda.Start(handler)
}
