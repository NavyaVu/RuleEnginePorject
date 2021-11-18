package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"log"
	"net/http"
	"os"
	"reflect"
	"ruleEngineProject/handlers"
	"ruleEngineProject/ruleEngine"
	"strings"

	//"ruleEngineProject/controller"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	ruleEngineInstance                      *engine.GruleEngine
	knowledgeBase                           *ast.KnowledgeBase
	knowledgeBaseName, knowledgeBaseVersion string
)

func init() {
	ruleEngineInstance = engine.NewGruleEngine()
	knowledgeBaseName = "Test"
	knowledgeBaseVersion = "0.0.1"
	knowledgeBase = ruleEngine.LoadRules(knowledgeBaseName, knowledgeBaseVersion)

}

func main() {

	lambda.Start(HandleRequest)

}

func HandleRequest(input interface{}) (interface{}, error) {
	log.Println("type:", reflect.TypeOf(input))

	var (
		req events.APIGatewayProxyRequest
		err error
	)

	inputRequest, err := json.Marshal(input)
	if err != nil {
		log.Println("unable to marshal input to json")
	}
	//log.Println(string(inputRequest))
	inputRequestAsString := string(inputRequest)
	containsHeader := strings.Contains(inputRequestAsString, "headers")
	err = json.Unmarshal(inputRequest, &req)
	if err != nil || !containsHeader {
		log.Println("unable to unmarshal input from json to APIGatewayRequest type or contains header is ", containsHeader)
		apiResponse, err := performRuleCheck(inputRequestAsString)
		if err != nil {
			return nil, err
		} else {
			fmt.Println("Response: ", apiResponse)
			return apiResponse, err
		}
	} else if containsHeader {
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

	return events.APIGatewayProxyResponse{Body: "Nothing Executed", StatusCode: http.StatusOK, IsBase64Encoded: false, Headers: nil}, err

}

func performRuleCheck(request string) (string, error) {
	l := log.New(os.Stdout, "Rules-API", log.LstdFlags)
	rc := handlers.NewRuleChecker(l)
	apiResponse, err := rc.GetRuleCheck(request, ruleEngineInstance, knowledgeBase)

	return string(apiResponse), err
}

//type MyEvent struct {
//	DepartureDateTimeInUtc string `json:"departure_date_time_in_utc"`
//	AirlineCode            string `json:"airline_code"`
//	BookingTimeInUtc       string `json:"booking_time_in_utc"`
//	Origin                 string `json:"origin"`
//	Destination            string `json:"destination"`
//	JourneyType            string `json:"journeyType"`
//}

//func HandleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
//	//rb := name
//
//	//if req.Headers["Httpmethod"] == http.MethodPost{
//	if req.HTTPMethod == http.MethodPost {
//		l := log.New(os.Stdout, "Rules-API", log.LstdFlags)
//		rc := handlers.NewRuleChecker(l)
//		apiResponse := string(rc.GetRuleCheck(req.Body))
//		return events.APIGatewayProxyResponse{Body: apiResponse, StatusCode: http.StatusOK, IsBase64Encoded: false, Headers: nil}, nil
//		//return events.APIGatewayV2HTTPResponse{Body: "{\"empname\":\"Rocky\",\"empid\":5454}", StatusCode: http.StatusOK, IsBase64Encoded: false, Headers: nil } , nil
//	} else {
//		err := errors.New("method not allowed NavyA" + req.Headers["Httpmethod"] + ":lambda ")
//		return events.APIGatewayProxyResponse{Body: "Method not ok", StatusCode: http.StatusBadRequest, IsBase64Encoded: false, Headers: nil}, err
//	}
//
//}

//func invokeRuleEngineLambda(request string) (response []byte, err error) {
//	sess := session.Must(session.NewSessionWithOptions(session.Options{
//		SharedConfigState: session.SharedConfigEnable,
//	}))
//
//	client := lambdaService.New(sess, &aws.Config{Region: aws.String("us-east-2")})
//
//	payload := []byte(request)
//
//	result, err := client.Invoke(&lambdaService.InvokeInput{FunctionName: aws.String("MyGetItemsFunction"), Payload: payload})
//	if err != nil {
//		fmt.Println("Error calling MyGetItemsFunction")
//		os.Exit(0)
//	}
//
//	return result.Payload, nil
//}
