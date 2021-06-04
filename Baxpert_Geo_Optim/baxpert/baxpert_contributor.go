package baxpert

import (
	"time"
)

//contributor type
type contributor struct {
	id           string
	availability []time.Time
}

//constructure
func Init_contributor(id string, availability []time.Time) contributor {
	return contributor{id, availability}
}

//make Date (int, time.Month, int, int, int, int, int, *time.Location)
func Create_Date(year int, month time.Month, day int) time.Time {
	//more tests to add ???
	if day > 30 {
		day = 1
		month += 1
	}
	if month > 12 {
		month = 1
		year += 1
	}
	return time.Date(year, month, day, 20, 34, 58, 651387237, time.UTC)

}

//check if p of type contributor in tmp ==> can we do better?
func (p *contributor) Has_P(tmp []contributor) bool {
	for _, b := range tmp {
		if b.id == p.id {
			return true
		}
	}
	return false
}
