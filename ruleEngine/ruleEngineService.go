package ruleEngine

import (
	"fmt"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"io/ioutil"
	"log"
	"path/filepath"
	"ruleEngineProject/models"
)

func LoadRules(name, version string) *ast.KnowledgeBase {
	lib := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(lib)

	rulesFile, err := filepath.Abs("./resources/rules.json")

	jsonData, err := ioutil.ReadFile(rulesFile)
	if err != nil {
		log.Panicln("Rules file not found ", err.Error())
	}
	ruleset, err := pkg.ParseJSONRuleset(jsonData)
	if err != nil {
		panic(err)
	}

	fmt.Println("Parsed ruleset: ")
	fmt.Println(ruleset)
	err = ruleBuilder.BuildRuleFromResource(name, version, pkg.NewBytesResource([]byte(ruleset)))
	kb := lib.NewKnowledgeBaseInstance(name, version)
	return kb
}

func LoadDataContextToKnowledgeBase(searchRequest *models.SearchRequest,
	searchResponse *models.SearchResponse) *ast.IDataContext {
	dataContext := ast.NewDataContext()
	err := dataContext.Add("FltSearchRequest", searchRequest)
	if err != nil {
		fmt.Println("Error while loading search request (fact): ", err)
	}
	err = dataContext.Add("RuleInfo", searchResponse)
	if err != nil {
		fmt.Println(err)
	}
	return &dataContext
}

func Execute(searchRequest *models.SearchRequest, searchResponse *models.SearchResponse,
	ruleEngineInstance *engine.GruleEngine,
	knowledgeBase *ast.KnowledgeBase) (*models.SearchResponse, error) {
	dataContext := LoadDataContextToKnowledgeBase(searchRequest, searchResponse)

	//fetch matching rules TODO remove after analysis
	ruleEntries, err := ruleEngineInstance.FetchMatchingRules(*dataContext, knowledgeBase)
	if err != nil {
		fmt.Println("Unable to fetch all matching rules")
		panic(err)
	} else {
		fmt.Println("Matching Rule(s):")
		for i := 0; i < len(ruleEntries); i++ {
			fmt.Println(ruleEntries[i].RuleName + " : " + ruleEntries[i].RuleDescription)
		}
	}

	err = ruleEngineInstance.Execute(*dataContext, knowledgeBase)
	if err != nil {
		fmt.Println(err)
	}
	return searchResponse, err
}
