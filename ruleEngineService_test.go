package main

import (
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/stretchr/testify/assert"
	"ruleEngineProject/config"
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
