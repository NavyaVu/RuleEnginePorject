package main

import (
	"log"
	"net/http"
	"os"
	"ruleEngineProject/handlers"

	"errors"
	"github.com/aws/aws-lambda-go/events"

	//"ruleEngineProject/controller"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {

	lambda.Start(HandleRequest)
	//controller.StartServer()
}

//type MyEvent struct {
//	DepartureDateTimeInUtc string `json:"departure_date_time_in_utc"`
//	AirlineCode            string `json:"airline_code"`
//	BookingTimeInUtc       string `json:"booking_time_in_utc"`
//	Origin                 string `json:"origin"`
//	Destination            string `json:"destination"`
//	JourneyType            string `json:"journeyType"`
//}

func HandleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	//rb := name

	//if req.Headers["Httpmethod"] == http.MethodPost{
	if req.HTTPMethod == http.MethodPost {
		l := log.New(os.Stdout, "Rules-API", log.LstdFlags)
		rc := handlers.NewRuleChecker(l)
		apiResponse := string(rc.GetRuleCheck(req.Body))
		return events.APIGatewayProxyResponse{Body: apiResponse, StatusCode: http.StatusOK, IsBase64Encoded: false, Headers: nil}, nil
		//return events.APIGatewayV2HTTPResponse{Body: "{\"empname\":\"Rocky\",\"empid\":5454}", StatusCode: http.StatusOK, IsBase64Encoded: false, Headers: nil } , nil
	} else {
		err := errors.New("method not allowed NavyA" + req.Headers["Httpmethod"] + ":lambda ")
		return events.APIGatewayProxyResponse{Body: "Method not ok", StatusCode: http.StatusBadRequest, IsBase64Encoded: false, Headers: nil}, err
	}

}

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
