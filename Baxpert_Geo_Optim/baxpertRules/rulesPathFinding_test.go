package baxpertRules

import (
	"fmt"
	"testing"
)

func TestRulePathFinding(t *testing.T) {

	// instantiate new constraint
	c := &Constraint{}
	c.Title = "Nature travaux"
	c.getList() // Empty List

	// instantiate rule
	rulePathFind := NewRulePathFinding()

	c.AssignRuleIntoConstraint(rulePathFind)
	c.getList()

	rulePathFind.ProcessConstraint(c)

	c.ProcessList()
	c.getList()

	fmt.Printf("globalScore %v \n", c.GlobalScore)

}
