package service

import (
	"fmt"
	"ruleEngineProject/models"
	"ruleEngineProject/ruleEngine"
)

type FlightCacheService struct {
	Request  *models.SearchRequest
	Response *models.SearchResponse
}

func (f FlightCacheService) Search(knowledgeBaseDetails *models.KnowledgeBaseForCacheRule) *models.SearchResponse {
	response := ruleEngine.Execute(f.Request, f.Response,
		knowledgeBaseDetails.Name, knowledgeBaseDetails.Version)

	fmt.Println(response.Cacheable)
	return response
}
