package baxpertRules

import (
	"database/sql"

	"fmt"
	"testing"
	"time"
)

//Test Dispo_client
func TestDispo_Site(t *testing.T) {
	layout := "2006-01-02"
	r := NewRuleAvailabilitySite()
	date1, _ := time.Parse(layout, "2021-06-01")
	date2, _ := time.Parse(layout, "2021-06-02")
	r.Intervention_date = sql.NullTime{Time: date1, Valid: true}
	r.Intervention_deadline = sql.NullTime{Time: date2, Valid: true}

	if r.Dispo_Site() {
		fmt.Println("Dispo_Site is ok")
	} else {
		fmt.Println("Dispo_Site is not ok")
	}
}

//Test RuleAvailabilityClient
func TestRuleAvailabilitySite(t *testing.T) {
	var rule RuleAvailabilitySite
	layout := "2006-01-02"

	// instantiate rule
	r := NewRuleAvailabilitySite()
	date1, _ := time.Parse(layout, "2021-06-01")
	date2, _ := time.Parse(layout, "2021-06-02")
	r.Intervention_date = sql.NullTime{Time: date1, Valid: true}
	r.Intervention_deadline = sql.NullTime{Time: date2, Valid: true}

	// instantiate new constraint
	c := &Constraint{}
	c.Title = "Availability Intervention Site"
	c.getList() // Empty List

	//Assign rule to the constraint
	c.AssignRuleIntoConstraint(r)

	rule.Convert(c.ConstraintList[0])
	rule.ProcessConstraint(c)

	c.ProcessList()
	c.getList()

	fmt.Printf("globalScore %v \n", c.GlobalScore)


}
