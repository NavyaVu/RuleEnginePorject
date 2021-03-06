package main

import (
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/stretchr/testify/assert"
	"ruleEngineProject/config"
	"ruleEngineProject/handlers"
	"ruleEngineProject/models"
	"ruleEngineProject/ruleEngine"
	"testing"
	"time"
)

var (
	ruleEngineInstance *engine.GruleEngine
	knowledgeBase      = &ast.KnowledgeBase{
		Name:          "testRuleEngine",
		Version:       "1.0",
		DataContext:   nil,
		WorkingMemory: nil,
		RuleEntries:   nil,
	}
	ruleEngineProperties = config.LoadProperties()
	//se *models.SearchResponse
)

func setup() {
	filePath, _ := ruleEngineProperties.Get("rule-file-path")
	knowledgeBase = ruleEngine.LoadRules(knowledgeBase.Name, knowledgeBase.Version, filePath)
	ruleEngineInstance = engine.NewGruleEngine()
}

func Test_CheckForRoute(t *testing.T) {
	var (
		err error
	)
	setup()
	unCacheableRequest := &models.SearchRequest{
		Cacheable:            false,
		AirlineCode:          "AF",
		DepartureAirportCode: "NYC",
		ArrivalAirportCode:   "HAJ",
		DepartureDateTime:    time.Now().Add(time.Hour * 24),
		ArrivalDateTime:      time.Now().Add(time.Hour * 72),
		RoundTrip:            true,
		BookingTime:          time.Now(),
	}

	cacheableRequest := &models.SearchRequest{
		Cacheable:            false,
		AirlineCode:          "AF",
		DepartureAirportCode: "AMS",
		ArrivalAirportCode:   "HAJ",
		DepartureDateTime:    time.Now().Add(time.Hour * 24),
		ArrivalDateTime:      time.Now().Add(time.Hour * 72),
		RoundTrip:            true,
		BookingTime:          time.Now(),
	}

	//response.AddDays(time.Now(), 8)
	response := &models.SearchResponse{}
	response, err = ruleEngine.Execute(unCacheableRequest, response, ruleEngineInstance, knowledgeBase)
	//log.Println(response.AddDays(request.DepartureDateTime, 3), ": Days")
	assert.NoError(t, err)
	assert.Equal(t, false, response.Cacheable, "Cache validation should pass but didn't for flt route")

	response, err = ruleEngine.Execute(cacheableRequest, response, ruleEngineInstance, knowledgeBase)
	//log.Println(response.AddDays(request.DepartureDateTime, 3), ": Days")
	assert.NoError(t, err)
	assert.Equal(t, true, response.Cacheable, "Cache validation should pass but didn't for flt route")
}

func Test_CheckForPastDate(t *testing.T) {
	var (
		err error
	)
	setup()
	request := &models.SearchRequest{
		Cacheable:            false,
		AirlineCode:          "AF",
		DepartureAirportCode: "NYC",
		ArrivalAirportCode:   "HAJ",
		DepartureDateTime:    time.Date(2021, 11, 23, 0, 0, 0, 0, time.Local),
		ArrivalDateTime:      time.Date(2021, 11, 25, 0, 0, 0, 0, time.Local),
		RoundTrip:            true,
		BookingTime:          time.Now(),
	}

	response := &models.SearchResponse{}

	//response.AddDays(time.Now(), 8)

	response, err = ruleEngine.Execute(request, response, ruleEngineInstance, knowledgeBase)
	//log.Println(response.AddDays(request.DepartureDateTime, 3), ": Days")
	assert.NoError(t, err)
	assert.Equal(t, false, response.Cacheable, "Should return false for past dates")
}

func Test_CheckForPastDateWithFutureDeptDate(t *testing.T) {
	var (
		err error
	)
	setup()
	request := &models.SearchRequest{
		Cacheable:            false,
		AirlineCode:          "AF",
		DepartureAirportCode: "AMS",
		ArrivalAirportCode:   "JFK",
		DepartureDateTime:    time.Date(2022, 11, 23, 0, 0, 0, 0, time.Local),
		ArrivalDateTime:      time.Date(2022, 11, 25, 0, 0, 0, 0, time.Local),
		RoundTrip:            true,
		BookingTime:          time.Now(),
	}
	response := &models.SearchResponse{}
	response, err = ruleEngine.Execute(request, response, ruleEngineInstance, knowledgeBase)
	assert.NoError(t, err)
	assert.Equal(t, true, response.Cacheable, "Should return true for this route and future dates")
}

func Test_CheckForPeakworkRule(t *testing.T) {
	var (
		err error
	)
	setup()
	request := &models.SearchRequest{
		Cacheable:            false,
		AirlineCode:          "AF",
		DepartureAirportCode: "NYC",
		ArrivalAirportCode:   "HAJ",
		DepartureDateTime:    time.Date(2021, 11, 23, 0, 0, 0, 0, time.Local),
		ArrivalDateTime:      time.Date(2021, 11, 25, 0, 0, 0, 0, time.Local),
		RoundTrip:            true,
		BookingTime:          time.Now(),
		RequestType:          "GetSearchScenarios",
		RuleGroup:            "Peakwork",
	}

	response := &models.SearchResponse{
		Cacheable:    false,
		AirlineCode:  "AF",
		Destinations: "",
		Origins:      "",
	}

	//response.AddDays(time.Now(), 8)

	response, err = ruleEngine.Execute(request, response, ruleEngineInstance, knowledgeBase)
	//log.Println(response.AddDays(request.DepartureDateTime, 3), ": Days")
	assert.NoError(t, err)
	assert.NotEmpty(t, response.Origins, "Should have some origin airport codes")
}

func Test_RetrievePeakWorkConfig(t *testing.T) {
	/*request := &models.FlightCacheSearchQuery{
		DepartureDateTimeInUtc: "",
		AirlineCode:            "",
		BookingTimeInUtc:       "",
		Origin:                 "",
		Destination:            "",
		JourneyType:            "",
		RequestType:            "GetSearchScenarios",
		RequestGroup:           "Peakwork",
	}*/

	request := &models.FlightCacheSearchQuery{
		DepartureDateTimeInUtc: "2022-12-01",
		AirlineCode:            "EK",
		BookingTimeInUtc:       "2022-12-01",
		Origin:                 "AMS",
		Destination:            "JFK",
		JourneyType:            "ONEWAY",
		RequestType:            "GetSearchScenarios",
		RequestGroup:           "Peakwork",
	}

	var (
		err error
	)
	setup()
	searchRequest := handlers.TranslateRequest(request)

	response := &models.SearchResponse{}
	response, err = ruleEngine.Execute(searchRequest, response, ruleEngineInstance, knowledgeBase)
	assert.NoError(t, err)
	assert.Equal(t, true, response.Cacheable, "Should return true for this route and future dates")
	assert.Equal(t, "2022-02-01", response.PeakworkEarliestDepDate, "Should return true for this route and future dates")
}
