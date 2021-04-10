package baxpert

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestGlobal(t *testing.T) {

	var ot_possible = globalSol{} //save all ot with their numbers
	var dates = make([]time.Time, 50)
	var contributors = make([]contributor, 50)
	var ots = make([]WOList, 50)
	var id string

	p := fmt.Println
	LayoutDate := "02-01-2006"

	//(int, time.Month, int, int, int, int, int, *time.Location)
	/*D1 := time.Date(2021, 03, 29, 20, 34, 58, 651387237, time.UTC)
	D2 := time.Date(2021, 03, 30, 20, 34, 58, 651387237, time.UTC)
	D3 := time.Date(2021, 03, 31, 20, 34, 58, 651387237, time.UTC)
	D4 := time.Date(2021, 04, 01, 20, 34, 58, 651387237, time.UTC)*/

	//dates creation (n = 50)
	for i := 0; i < 50; i++ {
		dates[i] = Create_Date(2021, 04, i+1)
	}

	/*T1 := contributor{"t1", []time.Time{D1, D2}}
	T2 := contributor{"t2", []time.Time{D1, D3}}
	T3 := contributor{"t3", []time.Time{D2, D4}}*/

	//contributors creation (n = 50)
	for i := 1; i < 50; i++ {
		id = "T" + strconv.Itoa(i)
		contributors[i] = Init_contributor(id, dates[(i%47):(i%47)+3])
	}

	/*ots1 := WOList{[]contributor{T1, T2, T3}, []time.Time{D1, D2, D4}}
	ots2 := WOList{[]contributor{T1, T2}, []time.Time{D1, D2, D3}}
	ots3 := WOList{[]contributor{T3}, []time.Time{D1, D3, D4}}*/

	//creation of the WOlists
	for i := 0; i < 50; i++ {
		t := rand.Intn(50)
		ots[i] = Init_WOList(contributors[t:50], dates[t:50])
	}

	//ots = []WOList{ots1, ots2, ots3}

	ot_possible.OrderWork(ots)
	ot_possible.MakeSolution()

	for key, value := range ot_possible {
		p("OT", key, ":")
		for _, s := range value {
			p("(", s.contributor.id, ",", s.Time.Format(LayoutDate), ") \n")
		}
	}

	Score(50)
	FragScore()

	for key, value := range solution {
		p("solution", key, ":")
		for _, s := range value {
			p("(", s.contributor.id, ",", s.Time.Format(LayoutDate), ")")
		}
		p("Score =", score[key], "%")
		p("Fragmentation Score =", FragmentationScore[key], "\n")

	}
}
