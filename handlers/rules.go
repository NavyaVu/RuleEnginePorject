package handlers

import (
	//"github.com/gorilla/mux"
	"log"
	"net/http"
	"ruleEngineProject/data"
	//"strconv"
)

type Rules struct {
	l *log.Logger
}

func NewRules(l *log.Logger) *Rules {
	return &Rules{l}
}

func (rl *Rules) GetRules(rw http.ResponseWriter, r *http.Request) {
	rl.l.Println("Handle Get Rules")
	lr := data.GetRules()
	err := lr.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Error during JSON Marshal", http.StatusInternalServerError)
	}
}

//addition of new rules

func (rl *Rules) AddRule(rw http.ResponseWriter, r *http.Request) {
	rl.l.Println("Handles Add Rules")

	newRule := &data.Rule{}
	err := newRule.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to Unmarshal", http.StatusBadRequest)
	}
	rl.l.Printf("Rule %v", newRule)
}

func (rl *Rules) UpdateRule(rw http.ResponseWriter, r *http.Request) {

	//vars := mux.Vars(r)
	//id, _ := strconv.Atoi(vars["id"])
	//rl.l.Println("Updates the Rule: ", id)
}
