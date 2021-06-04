package baxpert

import (
	"testing"
	//"time"
)

func TestWOCreate(t *testing.T) {
	contributor1 := contributor{id: "jmb"}
	wo1 := WO{contributor: contributor1}
	t.Logf("Expected %+v", wo1.contributor)
}
