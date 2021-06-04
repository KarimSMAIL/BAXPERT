package baxpertRules

import (
	//	"fmt"
	"testing"
)

func TestRules(t *testing.T) {

	//
	c := &Constraint{}
	c.Title = "test"
	c.getList() // Empty List

	// instantiate regles
	rule1 := newDummyRule()
	//rule2 := newRuleDepot()

	//
	c.AssignRuleIntoConstraint(rule1)
	//c.getList()
	c.ProcessList()
	c.getList()

	//	fmt.Printf("globalScore %v \n", c.globalScore)

	//
	/*regle2Adapter := &regle2Adapter{
		rule2:           rule2,
		function_title: "test f2",
	}

	//
	c.assignRuleIntoConstraint(regle2Adapter)
	//c.getList()
	//regle2Adapter.processConstraint(c)
	c.processList()
	c.getList()
	fmt.Printf("globalScore %v \n", c.globalScore)*/
}
