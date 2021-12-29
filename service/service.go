package service

import (
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"log"
	"ruleEngineProject/models"
	"ruleEngineProject/ruleEngine"
)

var FlightRuleServiceType FlightRuleService

type FlightRuleService interface {
	Search(ruleEngineInstance *engine.GruleEngine,
		knowledgeBase *ast.KnowledgeBase) (*models.SearchResponse, error)
}

type FlightCacheService struct {
	Request  *models.SearchRequest
	Response *models.SearchResponse
}

// Search Initiates the call to rule Engine service to check the Matching rules for the given Request
func (f FlightCacheService) Search(ruleEngineInstance *engine.GruleEngine,
	knowledgeBase *ast.KnowledgeBase) (*models.SearchResponse, error) {
	response, err := ruleEngine.Execute(f.Request, f.Response, ruleEngineInstance,
		knowledgeBase)

	log.Println(response.Cacheable)
	return response, err
}
