package handlers

import (
	"encoding/json"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/engine"
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

// GetRuleCheck Takes the request from Flight Cache service and checks if it needs to be cached or not
func (rl *Rules) GetRuleCheck(r string, ruleEngineInstance *engine.GruleEngine,
	knowledgeBase *ast.KnowledgeBase) ([]byte, error) {

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

	response, err := flightCacheService.Search(ruleEngineInstance, knowledgeBase)
	log.Println(searchRequest.DepartureDateTime, " console check ", response.Cacheable)
	responseData, err := json.Marshal(response) //(response.TfmRessponse)
	if err != nil {
		panic(err)
	}
	//_, err = rw.Write(responseData)
	//if err != nil {
	//	log.Println("Error in writing response: ", err.Error())
	//}
	return responseData, err
}

// Translates the given Flight Cache Search Query to Search Request
func translateRequest(query *models.FlightCacheSearchQuery) *models.SearchRequest {
	return &models.SearchRequest{
		Cacheable:            false,
		AirlineCode:          query.AirlineCode,
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
		log.Println(err)
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
