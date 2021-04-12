package baxpert

import (
	"time"
)

type globalSol map[int][]WO

var num_ot int  //number of OT
var num_sol = 1 //number of solutions

var solution = globalSol{}                 //possible solutions
var score = make(map[int]int)              //scores
var FragmentationScore = make(map[int]int) //Fragmentation scores

//generate a solution
func (gS globalSol) OrderWork(ots []WOList) {
	var sol_ot WO
	var sol []WO

	for _, ot := range ots {
		sol = nil
		//attribute an appointment of each contributor
		for _, contr := range ot.contributor {
			for _, appoit := range ot.appointment {
				//it's necessary to test if the contributor is available by this date
				if Has_D(contr.availability, appoit) {
					sol_ot.contributor = contr
					sol_ot.Time = appoit
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

//Make one solution
func (oti WO) Make_One_Solution(otj WO, n int) {
	if !oti.Has_OT(solution[n]) {
		solution[n] = append(solution[n], oti)
	}
	if !otj.Has_OT(solution[n]) {
		solution[n] = append(solution[n], otj)
	}
}

//Make all solutions
func (gS globalSol) MakeSolution() {
	n := 0

	for i := 0; i < len(gS); i++ {
		n = 0
		for _, oti := range gS[i] {
			for _, otj := range gS[i+1] {
				oti.Make_One_Solution(otj, n)
				n += 1
			}
		}

	}

	/*for _, ots1 := range gS[1] {
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

	/*solution[n] = append(solution[n], ots1)

				if !(ots1.contributor.id == ots2.contributor.id) {
					solution[n] = append(solution[n], ots2)
				}
				if !((ots1.contributor.id == ots3.contributor.id) || (ots1.contributor.id == ots3.contributor.id)) {
					solution[n] = append(solution[n], ots3)
				}
				n += 1
			}
		}
	}*/

}

//check if d of type time.Date in tmp
//cannot define new methods on non-local type time.Time
func Has_D(tmp []time.Time, d time.Time) bool {
	for _, b := range tmp {
		if b == d {
			return true
		}
	}
	return false
}

//solution coverage level
func Score(num_ot int) {

	for i, sol := range solution {
		score[i] = int((len(sol) * 100 / num_ot))
	}
}

//calculate the Fragmentation scores
func FragScore() {
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
		FragmentationScore[i] = len(tmp1) + len(tmp2)
	}
}
