package baxpertRules

import (
	"fmt"
	"database/sql"
	"math"
	"strconv"
	"github.com/mitchellh/mapstructure"
)

type RuleDepot struct {
	Name  string
	Depot_departure_latitude  sql.NullString
	Depot_departure_longitude sql.NullString
	OT_departure_latitude  sql.NullString
	OT_departure_longitude sql.NullString
	Score int
}

func NewRuleDepot() *RuleDepot {
	return &RuleDepot{
		Name:  "nearby deposit available",
		Score: 20,
	}
}

func (r *RuleDepot)Depot_Distance() float64{
	//convertir les coordonnées geo de string en float
	latDebut, _ := strconv.ParseFloat(r.Depot_departure_latitude.String, 64)
	longDebut, _ := strconv.ParseFloat(r.Depot_departure_longitude.String, 64)
	lat, _ := strconv.ParseFloat(r.OT_departure_latitude.String, 64)
	long, _ := strconv.ParseFloat(r.OT_departure_longitude.String, 64)

	//distance(A,B)=√((x2−x1)2+(y2−y1)2)
	return math.Sqrt(math.Pow(latDebut-lat, 2) + math.Pow(longDebut-long, 2))
}

func (ob *RuleDepot) Convert(event interface{})  {
	//ob := RuleWoNotAssigned{}
	mapstructure.Decode(event, &ob)

	//return ob
}

// assign ruleDepot to the constraint and add its score to the constraint score
func (r *RuleDepot) intoContrainte(c *Constraint) {

	c.ConstraintList = append(c.ConstraintList, r)
	fmt.Println("the ruleDepot is assigned to the constraint ", c.Title)
}

type Rule2Adapter struct {
	rule2          *RuleDepot
	function_title string
}

func (r *RuleDepot) AssignRuleIntoConstraint(c *Constraint) {
	fmt.Println("Assign Rule to the constraint.")
	r.intoContrainte(c)
}

func (r *RuleDepot) ProcessConstraint(c *Constraint) {
	//r.Depot_Distance()
	//according to the state performs a calculation
	c.GlobalScore += r.Score //c.globalScore++
	fmt.Printf("ruleDepot ProcessConstraint, globalScore %v \n", c.GlobalScore)

}
