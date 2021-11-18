package main

import (
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/stretchr/testify/assert"
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
)

func setup() {
	knowledgeBase = ruleEngine.LoadRules(knowledgeBase.Name, knowledgeBase.Version)
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
		DepartureAirportCode: "AMS",
		ArrivalAirportCode:   "NYC",
		DepartureDateTime:    time.Date(2021, 11, 22, 0, 0, 0, 0, time.Local),
		ArrivalDateTime:      time.Date(2021, 11, 25, 0, 0, 0, 0, time.Local),
		RoundTrip:            true,
		BookingTime:          time.Now(),
	}

	response := &models.SearchResponse{}

	response, err = ruleEngine.Execute(request, response, ruleEngineInstance, knowledgeBase)
	assert.NoError(t, err)
	assert.Equal(t, true, response.Cacheable, "Cache validation should pass but didn't for flt time")
}
