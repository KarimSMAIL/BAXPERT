package baxpertRules

import (
	"Outil_de_gestion_des_tournees/missionGeoOptim"
	//bxprt "Baxpert_Geo_Optim/baxpert"
	"database/sql"
	"fmt"
	"testing"
)

//Test CheckAssignement
func TestCheckAssignement(t *testing.T) {
	jobWo := missionGeoOptim.JobWO{}
	jobInfo := missionGeoOptim.JobInfo{}
	jobInfo.Etat_interne = sql.NullString{String: "Non affecté", Valid: true}
	jobWo.JobInfo = jobInfo
	ruleWoNotassign := NewRuleWoNotAssigned(jobWo.JobInfo.Etat_interne)

	if ruleWoNotassign.CheckAssignement() {
		fmt.Println("ok")
	} else {
		fmt.Println("error")
	}
}

//Test Convert
func TestConvert_(t *testing.T) {
	var ob interface{}

	rule_base := NewRuleWoNotAssigned(sql.NullString{String: "Non affecté", Valid: true})
	ob = rule_base

	var rule RuleWoNotAssigned
	rule.Convert(ob)

	if rule.Score > 0 {
		fmt.Println("ok")
	} else {
		fmt.Println("error")
	}
}

//Test Fonctionnel
func TestRuleWoNotAssigned(t *testing.T) {

	var rule RuleWoNotAssigned

	// instantiate new constraint
	c := &Constraint{}
	c.Title = "OT"
	c.getList() // Empty List

	jobWo := missionGeoOptim.JobWO{}
	jobInfo := missionGeoOptim.JobInfo{}
	jobInfo.Etat_interne = sql.NullString{String: "Non affecté", Valid: true}
	jobWo.JobInfo = jobInfo

	// instantiate rule
	ruleWoNotassign := NewRuleWoNotAssigned(jobWo.JobInfo.Etat_interne)

	/*ruleWoNotAssignAdapter := &RuleWoNotAssignedAdapter{
		ruleWoNotAssign: ruleWoNotassign,
		function_title:  "Check Assignement",
	}*/

	c.AssignRuleIntoConstraint(ruleWoNotassign)

	rule.Convert(c.ConstraintList[0])

	rule.ProcessConstraint(c)

	c.ProcessList()
	c.getList()

	fmt.Printf("globalScore %v \n", c.GlobalScore)

}
