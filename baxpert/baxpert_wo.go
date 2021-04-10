package baxpert

import (
	"time"
)

//Type of Work Order (solution) example : { OT1(T1,D1) OT2(T2,D3) OT3(T3,D2) }
type WO struct {
	contributor
	time.Time
}

//type of Work Orders possible, example: {OT1({T1,T2},{D1,D2}) ; OT2({T1,T2,T3},{D2,D3}) ; OT3({T3},{D1,D2,D3})}
type WOList struct {
	contributor []contributor //List of contributors
	appointment []time.Time   //list of possible dates
}

//Create WO List
func Init_WOList(cs []contributor, ds []time.Time) WOList {
	return WOList{cs, ds}
}

//check if ot in solution ==> can we do better?
func (ot *WO) Has_OT(tmp []WO) bool {
	for _, b := range tmp {
		if b.contributor.id == ot.contributor.id && b.Time == ot.Time {
			return true
		}
	}
	return false
}
