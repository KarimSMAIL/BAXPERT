package baxpert

import (
	"fmt"
	"testing"
	"time"
)

func TestGlobal(t *testing.T) {
	var ot_possible = globalSol{} //enregistrer tous les ot aves leurs numeros
	p := fmt.Println
	LayoutDate := "02-01-2006"
	//(int, time.Month, int, int, int, int, int, *time.Location)
	D1 := time.Date(2021, 03, 29, 20, 34, 58, 651387237, time.UTC)
	D2 := time.Date(2021, 03, 30, 20, 34, 58, 651387237, time.UTC)
	D3 := time.Date(2021, 03, 31, 20, 34, 58, 651387237, time.UTC)
	D4 := time.Date(2021, 04, 01, 20, 34, 58, 651387237, time.UTC)

	T1 := contributor{"t1", []time.Time{D1, D2, D4}}
	T2 := contributor{"t2", []time.Time{D1, D3}}
	T3 := contributor{"t3", []time.Time{D2, D4}}

	ots1 := WOList{[]contributor{T1, T2, T3}, []time.Time{D1, D2, D4}}
	ots2 := WOList{[]contributor{T1, T2}, []time.Time{D1, D2, D3}}
	ots3 := WOList{[]contributor{T3}, []time.Time{D1, D3, D4}}

	ots := []WOList{ots1, ots2, ots3}

	ot_possible.OrdreTravail(ots)
	ot_possible.MakeSolution()

	for key, value := range ot_possible {
		p("OT", key, ":")
		for _, s := range value {
			p("(", s.contributor.id, ",", s.Time.Format(LayoutDate), ") \n")
		}
	}

	Score(3)
	ScoreEmiette()

	for key, value := range solution {
		p("solution", key, ":")
		for _, s := range value {
			p("(", s.contributor.id, ",", s.Time.Format("02-01-2006"), ")")
		}
		p("score =", score[key], "%")
		p("score Emiettement =", scoreEmiette[key], "\n")

	}
}