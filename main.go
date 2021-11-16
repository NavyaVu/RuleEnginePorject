package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"reflect"
	"ruleEngineProject/handlers"
	"strings"

	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {

	lambda.Start(HandleRequest)
	//controller.StartServer()
}

func HandleRequest(input interface{}) (interface{}, error) {
	log.Println("type:", reflect.TypeOf(input))

	var req events.APIGatewayProxyRequest
	var err error

	inputRequest, err := json.Marshal(input)
	if err != nil {
		log.Println("unable to marshal input to json")
	}
	//log.Println(string(inputRequest))
	inputRequestAsString := string(inputRequest)
	containsHeader := strings.Contains(inputRequestAsString, "headers")

	if !containsHeader {
		log.Println("unable to unmarshal input from json to APIGatewayRequest type or contains header is ", containsHeader)
		apiResponse, err := performRuleCheck(inputRequestAsString)
		if err != nil {
			return nil, err
		} else {
			return apiResponse, err
		}
	} else if containsHeader {
		err = json.Unmarshal(inputRequest, &req)
		if err != nil {
			log.Println("unable to unmarshal input from json")
		}
		log.Println("request path", req.Path)
		if req.HTTPMethod == http.MethodPost {
			apiResponse, err := performRuleCheck(req.Body)
			if err != nil {
				return events.APIGatewayProxyResponse{Body: "Couldn't perform rule check", StatusCode: http.StatusOK, IsBase64Encoded: false, Headers: nil}, nil
			} else {
				return events.APIGatewayProxyResponse{Body: apiResponse, StatusCode: http.StatusOK, IsBase64Encoded: false, Headers: nil}, nil
			}

		} else {
			err = errors.New("method not allowed NavyA" + req.Headers["Httpmethod"] + ":lambda ")
		}

	} else {
		err = errors.New("nothing executed")
	}

	if containsHeader {
		return events.APIGatewayProxyResponse{Body: "Nothing Executed", StatusCode: http.StatusOK, IsBase64Encoded: false, Headers: nil}, err
	} else {
		return "Nothing Executed", err
	}

}

func performRuleCheck(request string) (string, error) {
	l := log.New(os.Stdout, "Rules-API", log.LstdFlags)
	rc := handlers.NewRuleChecker(l)
	apiResponse := string(rc.GetRuleCheck(request))

	return apiResponse, nil
}
