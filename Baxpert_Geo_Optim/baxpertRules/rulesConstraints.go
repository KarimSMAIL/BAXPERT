package baxpertRules

import (
	"fmt"

	bxprAdtr "Baxpert_Geo_Optim/baxpertAdapter"
)

type rule interface {
	AssignRuleIntoConstraint(c *Constraint)
	ProcessConstraint(c *Constraint)
}

type Constraint struct {
	ConstraintList []interface{}
	WO             bxprAdtr.JobWO
	Tech           bxprAdtr.JobContributorUserAgent
	Title          string
	GlobalScore    int
}

func (c *Constraint) AssignRuleIntoConstraint(r rule) {
	//here the append is done as r rule
	r.AssignRuleIntoConstraint(c) //ok we call the "parent" method
	//fmt.Printf("assign rule %+v to constraint\n\n", r)
	//c.ConstraintList = append(c.ConstraintList, r)
}

func (c *Constraint) getList() {
	if c.ConstraintList != nil {
		fmt.Printf("The rules of constraint: %+v\n\n", c.ConstraintList...)
	} else {
		fmt.Printf("Empty List!\n\n")
	}
}

func (c *Constraint) ProcessList() {
	fmt.Printf("processList \n")
	for _, a := range c.ConstraintList {
		if p, ok := a.(rule); ok {
			fmt.Printf("goodItem %+v !\n\n", p)
			p.ProcessConstraint(c)
		} else {
			//fmt.Printf("badItem %+v !\n\n", a) // the cast couldn't be performed
		}
	}
}
