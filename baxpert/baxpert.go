package baxpert

import (
	"time"
)

type globalSol map[int][]WO

var num_ot int  //numero des OT
var num_sol = 1 //numero des solutions

var solution = globalSol{}           //solution possible
var score = make(map[int]int)        //scores
var scoreEmiette = make(map[int]int) //ScoreEmiettement

//type des ordres de travailles possible, exemple: {OT1({T1,T2},{D1,D2}) ; OT2({T1,T2,T3},{D2,D3}) ; OT3({T3},{D1,D2,D3})}
type WOList struct {
	contributor []contributor //List des intervenents
	appointment []time.Time   //list des dates possibles
}

//generer une solution
func (gS globalSol) OrdreTravail(ots []WOList) {
	var sol_ot WO
	var sol []WO

	for _, ot := range ots {
		sol = nil
		//attribuer une date à chaque intervenant
		for _, inte := range ot.contributor {
			for _, d := range ot.appointment {
				//il faut tester si l'intervenants est disponible dans cette date
				if Has_D(inte.availability, d) {
					sol_ot.contributor = inte
					sol_ot.Time = d
					sol = append(sol, sol_ot)
				} else {
					continue
				}
			}
		}
		gS.SaveMap(sol)
	}
	return
}

func (gS globalSol) SaveMap(sol []WO) {
	num_ot += 1
	gS[num_ot] = sol
}

func (gS globalSol) MakeSolution() {
	n := 0

	for _, ots1 := range gS[1] {
		for _, ots2 := range gS[2] {
			for _, ots3 := range gS[3] {
				/*if !ots1.Has_OT(solution[n]) {
					solution[n] = append(solution[n], ots1)
				}
				if !ots2.Has_OT(solution[n]) {
					solution[n] = append(solution[n], ots2)
				}
				if !ots3.Has_OT(solution[n]) {
					solution[n] = append(solution[n], ots3)
				}*/

				solution[n] = append(solution[n], ots1)

				if !(ots1.contributor.id == ots2.contributor.id) {
					solution[n] = append(solution[n], ots2)
				}
				if !((ots1.contributor.id == ots3.contributor.id) || (ots1.contributor.id == ots3.contributor.id)) {
					solution[n] = append(solution[n], ots3)
				}
				n += 1
			}
		}
	}

}

//check si d de type time.Date in tmp
//cannot define new methods on non-local type time.Time
func Has_D(tmp []time.Time, d time.Time) bool {
	for _, b := range tmp {
		if b == d {
			return true
		}
	}
	return false
}

//niveau de couverture de la solution
func Score(num_ot int) {

	for i, sol := range solution {
		score[i] = int((len(sol) * 100 / num_ot))
	}
}

//Score d’émiettement
func ScoreEmiette() {
	var tmp1 []contributor
	var tmp2 []time.Time

	for i, sol := range solution {
		tmp1 = nil
		tmp2 = nil
		for _, ot := range sol {
			if !ot.contributor.Has_P(tmp1) {
				tmp1 = append(tmp1, ot.contributor)
			}
			if !Has_D(tmp2, ot.Time) {
				tmp2 = append(tmp2, ot.Time)
			}
		}
		scoreEmiette[i] = len(tmp1) + len(tmp2)
	}
}
