package baxpert

import (
	"time"
)

//Ordre de travail (solution) exemple : { OT1(T1,D1) OT2(T2,D3) OT3(T3,D2) }
type WO struct {
	contributor
	time.Time
}

//check si ot in solution ==> peut-on faire mieux?
/*func (ot *WO) Has_OT(tmp []WO) bool {
	for _, b := range tmp {
		if b.contributor.id == ot.contributor.id {
			return true
		}
	}
	return false
}*/
