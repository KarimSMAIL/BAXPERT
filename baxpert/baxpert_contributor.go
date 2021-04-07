package baxpert

import (
	"time"
)

//type des intervenants
type contributor struct {
	id           string
	availability []time.Time
}

//check si p de type intervenant in tmp ==> peut-on faire mieux?
func (p *contributor) Has_P(tmp []contributor) bool {
	for _, b := range tmp {
		if b.id == p.id {
			return true
		}
	}
	return false
}
