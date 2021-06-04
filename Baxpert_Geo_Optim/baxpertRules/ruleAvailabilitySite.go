package baxpertRules

import (
	"database/sql"

	"github.com/mitchellh/mapstructure"
	//"fmt"
)

type RuleAvailabilitySite struct {
	IdRegle     			int64 //KS
	Name       				string
	Intervention_date 		sql.NullTime
	Probable_duration       sql.NullString
	Intervention_deadline 	sql.NullTime
	Score       			int
}

//New Rule
func NewRuleAvailabilitySite() *RuleAvailabilitySite {
	return &RuleAvailabilitySite{
		Name:  "Intervention site availability",
		Score: 7,
	}
}

func (ob *RuleAvailabilitySite) Convert(event interface{}) {
	//ob := RuleWoNotAssigned{}
	mapstructure.Decode(event, &ob)

	//return ob
}

func (r *RuleAvailabilitySite) Dispo_Site() bool {
	return true
}

type ruleAvailabilitySiteAdapter struct {
	ruleAvailabilitySite *RuleAvailabilitySite
	function_title         string
}

func NewRuleAvailabilitysiteAdapter(ruleAvailabilitysite *RuleAvailabilitySite, function_title string) *ruleAvailabilitySiteAdapter {
	return &ruleAvailabilitySiteAdapter{
		ruleAvailabilitySite: ruleAvailabilitysite,
		function_title:         function_title,
	}
}

func (r *RuleAvailabilitySite) AssignRuleIntoConstraint(c *Constraint) {
	//fmt.Println("Rule adapter converts  rule1 to ruleAvailabilityIntervenor.")
	r.intoContrainte(c)
}

// assign ruleDepot to the constraint and add its score to the constraint score
func (r RuleAvailabilitySite) intoContrainte(c *Constraint) {
	//fmt.Println("the ruleAvailabilityIntervenor is assigned to the constraint ", c.Title)
	c.ConstraintList = append(c.ConstraintList, r)
}

func (r *RuleAvailabilitySite) ProcessConstraint(c *Constraint) {
	if r.Dispo_Site() {
		//according to the state performs a calculation
		c.GlobalScore += r.Score //c.GlobalScore++
		//fmt.Printf("RuleAvailabilitySite processConstraint, Dispo_Client Ok => globalScore %v \n", c.GlobalScore)
	}
}
