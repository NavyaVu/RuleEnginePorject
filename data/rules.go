package data

import (
	"encoding/json"
	"io"
)

// Same as JSON defined values
// Rule defines the structure for an API product
type Rule struct {
	Id               int      `json:"id"`
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	Salience         int      `json:"salience"`
	Action1          Action   `json:"action"`
	FlightConditions []Flight `json:"flight"`
}

type Flight struct {
	LeftOperand  LeftOpe  `json:"leftOperand"`
	Operator     []Ope    `json:"operator"`
	RightOperand RightOpe `json:"rightOperand"`
}

type LeftOpe struct {
	LeftOperand  string `json:"leftOperand"`
	Operator     string `json:"operator"`
	RightOperand string `json:"rightOperand"`
}

type Ope struct {
	Operator string
}

type RightOpe struct {
	LeftOperand  string `json:"leftOperand"`
	Operator     string `json:"operator"`
	RightOperand string `json:"rightOperand"`
}

type Action struct {
	Cacheable bool `json:"cacheable"`
}

//yet to add the details of flight conditions

//type Flight struct {
//	LeftOperand LO
//	Operators Op
//	RightOperand Rp
//}

// Rules is a collection of Product
type Rules []*Rule

// ToJSON serializes the contents of the collection to JSON
// NewEncoder provides better performance than json.Unmarshal as it does not
// have to buffer the output into an in memory slice of bytes
// this reduces allocations and the overheads of the service
//
// https://golang.org/pkg/encoding/json/#NewEncoder
func (r *Rules) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(r)
}

func (r *Rule) FromJSON(re io.Reader) error {
	d := json.NewDecoder(re)
	return d.Decode(r)
}

// GetRules returns a list of products
func GetRules() Rules {
	return ruleList
}

//post :create a new resource -- Use it for new
//put: creates or replaces a resources  -- Use it for update

//func AddProduct(r *Rule)

// productList is a hard coded list of products for this
// example data source
var ruleList = []*Rule{
	&Rule{
		Id:          1,
		Name:        "KLMFltCacheCheck",
		Description: "when",
		Salience:    10,
		Action1: Action{
			Cacheable: true,
		},
		FlightConditions: []Flight{{
			LeftOperand: LeftOpe{
				LeftOperand:  "",
				Operator:     "",
				RightOperand: "",
			},
			Operator: []Ope{
				{Operator: "&&"},
			},
			RightOperand: RightOpe{
				LeftOperand:  "",
				Operator:     "",
				RightOperand: "",
			},
		}},
	},
	&Rule{
		Id:          2,
		Name:        "AFFltCacheCheck",
		Description: "when",
		Salience:    10,
		Action1: Action{
			Cacheable: true,
		},
	},
}
