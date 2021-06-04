package baxpertRules

import (
	"database/sql"

	"github.com/mitchellh/mapstructure"
	//"fmt"
)

type RuleAvailabilityClient struct {
	IdRegle     int64 //KS
	Name        string
	Date_RDV    sql.NullTime
	RdV_Fixe    sql.NullBool
	RDV_present sql.NullBool
	Score       int
}

//New Rule
func NewRuleAvailabilityClient() *RuleAvailabilityClient {
	return &RuleAvailabilityClient{
		Name:  "Client availability",
		Score: 12,
	}
}

func (ob *RuleAvailabilityClient) Convert(event interface{}) {
	//ob := RuleWoNotAssigned{}
	mapstructure.Decode(event, &ob)

	//return ob
}

func (r *RuleAvailabilityClient) Dispo_Client() bool {
	if (r.RdV_Fixe.Bool) && (r.RDV_present.Bool) {
		return true
	}
	return false
}

type ruleAvailabilityClientAdapter struct {
	ruleAvailabilityClient *RuleAvailabilityClient
	function_title         string
}

func NewRuleAvailabilityclientAdapter(ruleAvailabilityclient *RuleAvailabilityClient, function_title string) *ruleAvailabilityClientAdapter {
	return &ruleAvailabilityClientAdapter{
		ruleAvailabilityClient: ruleAvailabilityclient,
		function_title:         function_title,
	}
}

func (r *RuleAvailabilityClient) AssignRuleIntoConstraint(c *Constraint) {
	//fmt.Println("Rule adapter converts  rule1 to ruleAvailabilityIntervenor.")
	r.intoContrainte(c)
}

// assign ruleDepot to the constraint and add its score to the constraint score
func (r RuleAvailabilityClient) intoContrainte(c *Constraint) {
	//fmt.Println("the ruleAvailabilityIntervenor is assigned to the constraint ", c.Title)
	c.ConstraintList = append(c.ConstraintList, r)
}

func (r *RuleAvailabilityClient) ProcessConstraint(c *Constraint) {
	if r.Dispo_Client() {
		//according to the state performs a calculation
		c.GlobalScore += r.Score //c.GlobalScore++
		//fmt.Printf("RuleAvailabilityClient processConstraint, Dispo_Client Ok => globalScore %v \n", c.GlobalScore)
	}
}
