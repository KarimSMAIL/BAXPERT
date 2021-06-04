package baxpertRules

import (
	"database/sql"

	"fmt"
	"testing"
	//"time"
)

//Test Dispo_client
func TestDispo_Client(t *testing.T) {
	r := NewRuleAvailabilityClient()
	r.RdV_Fixe = sql.NullBool{Bool: true, Valid: true}
	r.RDV_present = sql.NullBool{Bool: true, Valid: true}

	if r.Dispo_Client() {
		fmt.Println("Dispo_Client is ok")
	} else {
		fmt.Println("Dispo_Client is not ok")
	}
}

//Test RuleAvailabilityClient
func TestRuleAvailabilityClient(t *testing.T) {
	var rule RuleAvailabilityClient

	// instantiate rule
	r := NewRuleAvailabilityClient()
	r.RdV_Fixe = sql.NullBool{Bool: true, Valid: true}
	r.RDV_present = sql.NullBool{Bool: true, Valid: true}

	// instantiate new constraint
	c := &Constraint{}
	c.Title = "Availability Client"
	c.getList() // Empty List

	//Assign rule to the constraint
	c.AssignRuleIntoConstraint(r)

	rule.Convert(c.ConstraintList[0])
	rule.ProcessConstraint(c)

	c.ProcessList()
	c.getList()

	fmt.Printf("globalScore %v \n", c.GlobalScore)


}
