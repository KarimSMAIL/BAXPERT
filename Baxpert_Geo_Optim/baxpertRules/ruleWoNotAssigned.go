package baxpertRules

import (
	"database/sql"
	//"fmt"
	"github.com/mitchellh/mapstructure"

)

type RuleWoNotAssigned struct {
	IdRegle        int64 //KS
	Name           string
	Internal_state sql.NullString
	Score          int
}

func NewRuleWoNotAssigned(Internal_state sql.NullString) *RuleWoNotAssigned {
	return &RuleWoNotAssigned{
		IdRegle:        110,
		Name:           "WO not assigned",
		Internal_state: Internal_state,
		Score:          15,
	}
}

func (r RuleWoNotAssigned) CheckAssignement() bool {
	return r.Internal_state.String == "Non affecté" //|| r.Internal_state == "A terminer"
}

func (ob *RuleWoNotAssigned) Convert(event interface{})  {
	//ob := RuleWoNotAssigned{}
	mapstructure.Decode(event, &ob)

	//return ob
}

func (r *RuleWoNotAssigned)ProcessConstraint(c *Constraint) {
	if r.CheckAssignement() {
		c.GlobalScore += r.Score //c.globalScore++
		//fmt.Printf("RuleWoNotAssigned processConstraint, CheckAssignement Ok => globalScore %v \n", c.GlobalScore)
	}
	//according to the state performs a calculation
}

type RuleWoNotAssignedAdapter struct {
	ruleWoNotAssign *RuleWoNotAssigned
	function_title  string
}

func NewRuleWoNotAssignedAdapter(ruleWoNotAssign *RuleWoNotAssigned, function_title string) *RuleWoNotAssignedAdapter {
	return &RuleWoNotAssignedAdapter{
		ruleWoNotAssign: ruleWoNotAssign,
		function_title:  function_title,
	}
}

func (r *RuleWoNotAssigned) AssignRuleIntoConstraint(c *Constraint) {
	//fmt.Println("Rule adapter converts  rule1 to RuleWoNotAssigned.")
	r.intoContrainte(c)
}

// assign RuleWoNotAssigned to the constraint and add its score to the constraint score
func (ruleWoNotAssign RuleWoNotAssigned) intoContrainte(c *Constraint) {
	c.ConstraintList = append(c.ConstraintList, ruleWoNotAssign)
	//fmt.Println("the RuleWoNotAssigned is assigned to the constraint ", c.Title)
}


/*func NewRuleWoNotAssigned() *RuleWoNotAssigned {
	return &RuleWoNotAssigned{
		IdRegle: 110,
		Name:    "WO not assigned",
		Score:   15,
	}
}

*/