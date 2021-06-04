package baxpertRules

import "fmt"

type dummyRule struct {
	Name  string
	Score int
}

func newDummyRule() *dummyRule {
	return &dummyRule{
		Name:  "ot unaffected",
		Score: 1,
	}
}

// assign dummyRule to the constraint and add its score to the constraint score
func (rgl1 *dummyRule) AssignRuleIntoConstraint(c *Constraint) {
	//c.constraintList = append(c.constraintList, rgl1) //here calls explicit redundant
	fmt.Println("dummyRule is assigned to the constraint c")
}

func (rgl1 *dummyRule) ProcessConstraint(c *Constraint) {
	c.GlobalScore += rgl1.Score // c.globalScore++
	fmt.Printf("dummyRule processConstraint, globalScore %v \n", c.GlobalScore)
}
