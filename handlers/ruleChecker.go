package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"ruleEngineProject/models"
	"ruleEngineProject/service"
	"time"
)

type RuleChecker struct {
	l *log.Logger
}

func NewRuleChecker(l *log.Logger) *Rules {
	return &Rules{l}
}

//func (rl *Rules) GetRuleCheck(rw http.ResponseWriter, r *http.Request){
func (rl *Rules) GetRuleCheck(r string) ([]byte, error) {
	//body, err := ioutil.ReadAll(r.Body)
	//if err != nil {
	//	panic(err)
	//}
	fcsQuery := models.FlightCacheSearchQuery{}
	err := json.Unmarshal([]byte(r), &fcsQuery)
	if err != nil {
		panic(err)
	}

	searchRequest := translateRequest(&fcsQuery)
	flightCacheService := &service.FlightCacheService{
		Request: searchRequest,
		Response: &models.SearchResponse{
			Cacheable:   false,
			AirlineCode: searchRequest.AirlineCode,
		},
	}
	kbDetails := &models.KnowledgeBaseForCacheRule{
		Name:    "Test",
		Version: "0.0.1",
	}
	response := flightCacheService.Search(kbDetails)
	fmt.Println(searchRequest.DepartureDateTime, " console check ", response.Cacheable)
	responseData, err := json.Marshal(response) //(response.TfmRessponse)
	if err != nil {
		panic(err)
	}
	//_, err = rw.Write(responseData)
	//if err != nil {
	//	fmt.Println("Error in writing response: ", err.Error())
	//}
	return responseData, err
}

func translateRequest(query *models.FlightCacheSearchQuery) *models.SearchRequest {
	return &models.SearchRequest{
		Cacheable:            false,
		AirlineCode:          "KL", //TODO limited to AF for the poc
		DepartureAirportCode: query.Origin,
		ArrivalAirportCode:   query.Destination,
		DepartureDateTime:    convertDate(query.DepartureDateTimeInUtc),
		ArrivalDateTime:      time.Time{},
		RoundTrip:            isRoundTripJourney(query.JourneyType),
		BookingTime:          time.Now(),
	}
}

func convertDate(date string) time.Time {
	layout := "2006-01-02"
	t, err := time.Parse(layout, date)

	if err != nil {
		fmt.Println(err)
	}

	return t
}

func isRoundTripJourney(journey string) bool {
	isRoundTripJourney := false
	if journey != "ONEWAY" {
		isRoundTripJourney = true
	}

	return isRoundTripJourney
}
