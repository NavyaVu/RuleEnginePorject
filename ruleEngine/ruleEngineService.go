package ruleEngine

import (
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"io/ioutil"
	"log"
	"path/filepath"
	"ruleEngineProject/models"
)

// LoadRules This is where the knowledge base is loaded
func LoadRules(name, version string, rulesFilePath string) *ast.KnowledgeBase {
	lib := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(lib)

	rulesFile, err := filepath.Abs(rulesFilePath)
	if err != nil {
		log.Panicln("Rules file not found at ", rulesFilePath, " error: ", err.Error())
	}

	jsonData, err := ioutil.ReadFile(rulesFile)
	if err != nil {
		log.Panicln("Rules file could not be parsed as json ", err.Error())
	}
	ruleset, err := pkg.ParseJSONRuleset(jsonData)
	if err != nil {
		panic(err)
	}

	log.Println("Parsed ruleset: ")
	log.Println(ruleset)
	err = ruleBuilder.BuildRuleFromResource(name, version, pkg.NewBytesResource([]byte(ruleset)))
	kb := lib.NewKnowledgeBaseInstance(name, version)
	return kb
}

// LoadDataContextToKnowledgeBase Adds the objects to data Context for comparing
func LoadDataContextToKnowledgeBase(searchRequest *models.SearchRequest,
	searchResponse *models.SearchResponse) *ast.IDataContext {
	dataContext := ast.NewDataContext()
	err := dataContext.Add("FltSearchRequest", searchRequest)
	if err != nil {
		log.Println("Error while loading search request (fact): ", err)
	}
	err = dataContext.Add("RuleInfo", searchResponse)
	if err != nil {
		log.Println(err)
	}
	return &dataContext
}

// Execute Checks for the matching Rules and returns True in Response if found and false for none matched
func Execute(searchRequest *models.SearchRequest, searchResponse *models.SearchResponse,
	ruleEngineInstance *engine.GruleEngine,
	knowledgeBase *ast.KnowledgeBase) (*models.SearchResponse, error) {
	dataContext := LoadDataContextToKnowledgeBase(searchRequest, searchResponse)

	//fetch matching rules TODO remove after analysis
	ruleEntries, err := ruleEngineInstance.FetchMatchingRules(*dataContext, knowledgeBase)
	if err != nil {
		log.Println("Unable to fetch all matching rules")
		panic(err)
	} else {
		log.Println("Matching Rule(s):")
		for i := 0; i < len(ruleEntries); i++ {
			log.Println(ruleEntries[i].RuleName + " : " + ruleEntries[i].RuleDescription)
		}
	}

	err = ruleEngineInstance.Execute(*dataContext, knowledgeBase)
	if err != nil {
		log.Println(err)
	}
	return searchResponse, err
}
