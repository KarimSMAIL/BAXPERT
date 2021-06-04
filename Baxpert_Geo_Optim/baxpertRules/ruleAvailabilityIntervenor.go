package baxpertRules

import (
	"database/sql"
	"time"
	"github.com/mitchellh/mapstructure"
	//"fmt"
)

type RuleAvailabilityIntervenor struct {
	IdRegle           int64 //KS
	Name              string
	Start_date        sql.NullTime
	End_date          sql.NullTime
	Intervention_date sql.NullTime
	Activite_OT       string
	Activite_Tech     string
	Score             int
}

//New Rule
func NewRuleAvailabilityIntervenor() *RuleAvailabilityIntervenor {
	return &RuleAvailabilityIntervenor{
		IdRegle:           120,
		Name:              "Intervenor availability",
		Score:             10,
	}
}
/*func NewRuleAvailabilityIntervenor(Start_date sql.NullTime, End_date sql.NullTime, Intervention_date sql.NullTime,
	Activite_OT string, Activite_Tech string) *RuleAvailabilityIntervenor {
	return &RuleAvailabilityIntervenor{
		IdRegle:           120,
		Name:              "Intervenor availability",
		Start_date:        Start_date,
		End_date:          End_date,
		Intervention_date: Intervention_date,
		Activite_OT:       Activite_OT,
		Activite_Tech:     Activite_Tech,
		Score:             10,
	}
}*/

//Check if check is betwen two dates
func InTimeSpan(start time.Time, end time.Time, check time.Time) bool {
	return check.After(start) && check.Before(end)
}

func (r *RuleAvailabilityIntervenor) ValidDispo() bool {
	return InTimeSpan(r.Start_date.Time, r.End_date.Time, r.Intervention_date.Time)
}

func (r *RuleAvailabilityIntervenor) ValidActivite() bool {
	return r.Activite_OT == r.Activite_Tech
}

func (ob *RuleAvailabilityIntervenor) Convert(event interface{})  {
	//ob := RuleWoNotAssigned{}
	mapstructure.Decode(event, &ob)

	//return ob
}

func (r *RuleAvailabilityIntervenor) ProcessConstraint(c *Constraint) {

	if r.ValidDispo(){
		//according to the state performs a calculation
		c.GlobalScore += r.Score //c.globalScore++
		//fmt.Printf("ruleAvailabilityIntervenor processConstraint, ValidDispo Ok => globalScore %v \n", c.GlobalScore)
	}
	if r.ValidActivite(){
		//according to the state performs a calculation
		c.GlobalScore += r.Score //c.globalScore++
		//fmt.Printf("ruleAvailabilityIntervenor processConstraint, ValidActivite Ok => globalScore %v \n", c.GlobalScore)
	}
}

type ruleAvailabilityIntervenorAdapter struct {
	ruleAvailabilityInter *RuleAvailabilityIntervenor
	function_title        string
}

func NewRuleAvailabilityIntervenorAdapter(ruleAvailabilityInter *RuleAvailabilityIntervenor, function_title string) *ruleAvailabilityIntervenorAdapter {
	return &ruleAvailabilityIntervenorAdapter{
		ruleAvailabilityInter: ruleAvailabilityInter,
		function_title:        function_title,
	}
}

func (r *RuleAvailabilityIntervenor) AssignRuleIntoConstraint(c *Constraint) {
	//fmt.Println("Rule adapter converts  rule1 to ruleAvailabilityIntervenor.")
	r.intoContrainte(c)
}

// assign ruleDepot to the constraint and add its score to the constraint score
func (ruleAvailabilityInter RuleAvailabilityIntervenor) intoContrainte(c *Constraint) {
	//fmt.Println("the ruleAvailabilityIntervenor is assigned to the constraint ", c.Title)
	c.ConstraintList = append(c.ConstraintList, ruleAvailabilityInter)
}


