package baxpertRules

import (
	"fmt"
	"database/sql"
	"testing"
)

func TestDepot_Distance(t *testing.T) {

	r := NewRuleDepot()
	r.Depot_departure_latitude = sql.NullString{String :  "45.959887", Valid : true }
	r.Depot_departure_longitude = sql.NullString{String : "5.334645" , Valid : true }
	r.OT_departure_latitude = sql.NullString{String :  "45.7578008", Valid : true }
	r.OT_departure_longitude = sql.NullString{String :  "4.8349228", Valid : true }

	fmt.Println("Distance : ", r.Depot_Distance())
}

func TestRulesDepot(t *testing.T) {

	//
	c := &Constraint{}
	c.Title = "test"
	c.getList() // Empty List

	// instantiate regle
	r := NewRuleDepot()
	r.Depot_departure_latitude = sql.NullString{String :  "45.959887", Valid : true }
	r.Depot_departure_longitude = sql.NullString{String : "5.334645" , Valid : true }
	r.OT_departure_latitude = sql.NullString{String :  "45.7578008", Valid : true }
	r.OT_departure_longitude = sql.NullString{String :  "4.8349228", Valid : true }

	fmt.Println("Distance : ", r.Depot_Distance())

	//
	c.AssignRuleIntoConstraint(r)

	//c.getList()
	r.ProcessConstraint(c)
	c.ProcessList()
	c.getList()

	fmt.Printf("globalScore %v \n", c.GlobalScore)

}
