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
	assert.Equal(t, true, response.Cacheable, "Cache validation should pass but didn't for flt time")
}
