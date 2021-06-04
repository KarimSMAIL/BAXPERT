package baxpertRules

import (
	//bxprAdtr "Baxpert_Geo_Optim/baxpertAdapter"
	//"Outil_de_gestion_des_tournees/missionGeoOptim"
	"database/sql"

	"fmt"
	"testing"
	"time"
)

//Test InTimeSpan
func TestInTimeSpan(t *testing.T) {
	layout := "2006-01-02"
	start, _ := time.Parse(layout, "2021-05-01")
	end, _ := time.Parse(layout, "2021-05-30")
	check, _ := time.Parse(layout, "2021-05-07")

	if InTimeSpan(start, end, check) {
		fmt.Println("ok")
	} else {
		fmt.Println("InTimeSpan malfunction")
	}
}

//Test ValidDispo
func TestValidDispo(t *testing.T) {
	/*layout := "2006-01-02"
	Start_date, _ := time.Parse(layout, "2021-05-01")
	Start := sql.NullTime{Time: Start_date, Valid: true}
	End_date, _ := time.Parse(layout, "2021-05-31")
	End := sql.NullTime{Time: End_date, Valid: true}
	Intervention_date, _ := time.Parse(layout, "2021-05-10")
	Intervention := sql.NullTime{Time: Intervention_date, Valid: true}
	Activite_OT := "S02"
	Activite_Tech := "RIT"
	rule := NewRuleAvailabilityIntervenor(Start, End, Intervention, Activite_OT, Activite_Tech)

	if rule.ValidDispo() {
		fmt.Println("ok")
	} else {
		fmt.Println("ValidDsipo malfunction")
	}*/
}

//Test ValidActivite
func TestValidActivite(t *testing.T) {
	/*layout := "2006-01-02"
	Start_date, _ := time.Parse(layout, "2021-05-01")
	Start := sql.NullTime{Time: Start_date, Valid: true}
	End_date, _ := time.Parse(layout, "2021-05-31")
	End := sql.NullTime{Time: End_date, Valid: true}
	Intervention_date, _ := time.Parse(layout, "2021-06-10")
	Intervention := sql.NullTime{Time: Intervention_date, Valid: true}

	Activite_OT := "S02"
	Activite_Tech := "S02"
	rule := NewRuleAvailabilityIntervenor(Start, End, Intervention, Activite_OT, Activite_Tech)

	if rule.ValidActivite() {
		fmt.Println("ok")
	} else {
		fmt.Println("ValidActivite malfunction")
	}*/
}

//Test Convert
func TestConvert(t *testing.T) {
	var ob interface{}

	rule_base := NewRuleWoNotAssigned(sql.NullString{String: "Non affecté", Valid: true})
	ob = rule_base
	
	var rule RuleAvailabilityIntervenor
	rule.Convert(ob)

	if rule.Score > 0 {
		fmt.Println("ok")
	} else {
		fmt.Println("error")
	}
}

func TestRuleAvailabilityIntervenor(t *testing.T) {
	/*layout := "2006-01-02"
	var rule RuleAvailabilityIntervenor

	// instantiate new constraint
	c := &Constraint{}
	c.Title = "Availability"
	c.getList() // Empty List

	// instantiate rule
	Start_date, _ := time.Parse(layout, "2021-05-01")
	Start := sql.NullTime{Time: Start_date, Valid: true}
	End_date, _ := time.Parse(layout, "2021-05-31")
	End := sql.NullTime{Time: End_date, Valid: true}
	Intervention_date, _ := time.Parse(layout, "2021-05-10")
	Intervention := sql.NullTime{Time: Intervention_date, Valid: true}

	Activite_OT := "S02"
	Activite_Tech := "S02"
	ruleAvailabInterv := NewRuleAvailabilityIntervenor(Start, End, Intervention, Activite_OT, Activite_Tech)
	c.AssignRuleIntoConstraint(ruleAvailabInterv)

	rule.Convert(c.ConstraintList[0])
	rule.ProcessConstraint(c)
	c.ProcessList()
	c.getList()

	fmt.Printf("globalScore %v \n", c.GlobalScore)*/
}
