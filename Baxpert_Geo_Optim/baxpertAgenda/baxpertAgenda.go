// baxpertAgenda project baxpertAgenda.go
package baxpertAgenda

// gestion des Agendas globaux
// décrit l'objet Agenda et ses méthodes
// dans les méthodes : écrétage d'agenda
// mise-à-jour de l'agenda via baxpertAdapter

import (
	bxprtAdptr "Baxpert_Geo_Optim/baxpertAdapter"
	cgo "Outil_de_gestion_des_tournees/contrainteGeoOptim"
	"Outil_de_gestion_des_tournees/contrainteIntervention"

	//	cinterv "Outil_de_gestion_des_tournees/contrainteIntervention"

	lgo "Outil_de_gestion_des_tournees/localisationGeoOptim"
	"database/sql"
	"fmt"
	"time"
)

var p = fmt.Println

type Agenda struct {
	ListIntervenantAgenda                   []int64
	ListTourneeByIntervenantAll             map[int64][]cgo.RegleIntervenantTournee
	ListTourneeByIntervenantAgenda          map[int64][]cgo.RegleIntervenantTournee
	ListTourneeByIntervenantNotAgenda       map[int64][]cgo.RegleIntervenantTournee
	ResultListIntervenantTournee            map[int64][]cgo.RegleIntervenantTournee
	Moyenne                                 int
	LastDateResultListIntervenantTournee    map[int64]time.Time // derniere date pour un intervenant
	ListTourneeByIntervenantNotAgendaSupMoy map[int64][]int64   // les tournees restantes aprés l'equilibrage aprés la moyenne
	LastDateByIntervenantAgenda             map[int64]time.Time // derniere date avant la moyenne

}

func (agenda *Agenda) InitAgenda(db *sql.DB) {
	agenda.ListTourneeByIntervenantAll = make(map[int64][]cgo.RegleIntervenantTournee)
	agenda.ListTourneeByIntervenantAgenda = make(map[int64][]cgo.RegleIntervenantTournee)
	agenda.ListTourneeByIntervenantNotAgenda = make(map[int64][]cgo.RegleIntervenantTournee)
	agenda.ResultListIntervenantTournee = make(map[int64][]cgo.RegleIntervenantTournee)
	agenda.LastDateResultListIntervenantTournee = make(map[int64]time.Time)
	agenda.ListTourneeByIntervenantNotAgendaSupMoy = make(map[int64][]int64)
	agenda.LastDateByIntervenantAgenda = make(map[int64]time.Time)
	var intervTournee cgo.IntervenantTournee

	agenda.ListIntervenantAgenda, _ = intervTournee.GetDumbIntervenantAgendaTournee(db)
	//	p(agenda.ListIntervenantAgenda)
	for _, id := range agenda.ListIntervenantAgenda {
		agenda.ListTourneeByIntervenantAll[id], _ = cgo.GetTotalisationTourneeIntervenantAgenda(db, id)
		//	p(agenda.ListTourneeByIntervenantAll[id])
	}

	//init Moyenne
	agenda.Average()

	//Diviser les intervenants par moyenne
	for id, value := range agenda.ListTourneeByIntervenantAll {
		if len(value) > agenda.Moyenne {
			agenda.ListTourneeByIntervenantAgenda[id] = value
		} else {
			agenda.ListTourneeByIntervenantNotAgenda[id] = value
			agenda.ResultListIntervenantTournee[id] = nil
		}
	}
}

func (agenda *Agenda) Average() {
	//Calcule Moyenne
	var sum int
	for _, value := range agenda.ListTourneeByIntervenantAll {
		sum += len(value)
	}
	if sum != 0 {
		agenda.Moyenne = sum/len(agenda.ListTourneeByIntervenantAll) + 5
	}
}

func VerifyActivityProduit(liste_ots []lgo.Localisation, intervenant bxprtAdptr.Intervenant) bool {
	var tech_ok bool
	//fmt.Println("intervenants : ", len(liste_intervenants.Intervenants))
	Qualite := bxprtAdptr.GetJobContributorSkill(intervenant.QualiteIntervenantJson)
	//fmt.Println("ots : ", len(liste_ots))
	for _, ot := range liste_ots {
		tech_ok = false
		Activiteproduit := bxprtAdptr.GetActiviteProduit(ot.TypeJson.String)
		for key, _ := range Qualite.Specialite {
			for key2, _ := range Qualite.Produit {
				if Activiteproduit.Activite == key && Activiteproduit.Produit == key2 {
					tech_ok = true
					break
				}
			}
		}
		if tech_ok == false {
			break
		}
	}
	return tech_ok
}

func Exist(idIntervenant int64, intervenants bxprtAdptr.Intervenants) bool {
	for _, inter := range intervenants.Intervenants {
		if inter.IdIntervenant == idIntervenant {
			return true
		}
	}
	return false
}

func (agenda *Agenda) PutRegleIntervenantTournee(db *sql.DB) {
	//	var last_date = make(map[int64]int)
	agenda.LastDateResultListIntervenantTournee = cgo.GetLastDateListIntervenantTournee(db, agenda.ResultListIntervenantTournee)
	for idIntervenant, values := range agenda.ResultListIntervenantTournee {
		for _, regleIntervTournee := range values {
			var intervTournee = cgo.IntervenantTournee{
				PrioriteRegle: regleIntervTournee.RegleIntervenant.PrioriteRegle,
				IdRegle:       regleIntervTournee.RegleIntervenant.IdRegle,
				IdTournee:     sql.NullInt64{Int64: regleIntervTournee.IdTournee, Valid: true},
				IdIntervenant: regleIntervTournee.RegleIntervenant.IdIntervenant,
			}
			//last_date[idIntervenant] += 1
			intervTournee.UpdateIntervenantTournee(db, idIntervenant)
			var trnPrevu = contrainteIntervention.TourneePrevu{}
			trnPrevu.GetTourneeByIdTournee(db, regleIntervTournee.IdTournee)

			date := agenda.LastDateResultListIntervenantTournee[idIntervenant]
			date = date.AddDate(0, 0, 1)
			trnPrevu.UpdateAllDateDayTourneePrevu(db, date)
			agenda.LastDateResultListIntervenantTournee[idIntervenant] = date
		}
	}
}

func (agenda *Agenda) Planing(db *sql.DB) {
	for _, tournees := range agenda.ListTourneeByIntervenantAgenda { //[]RegleIntervenantTournee
		for i, tournee := range tournees {

			if i > agenda.Moyenne {

				intervenants, err := cgo.GetIntervenantByRegle(db, tournee.RegleIntervenant)
				if err != nil {
					fmt.Println(err)
				}
				for id2, tournees2 := range agenda.ResultListIntervenantTournee {
					if len(tournees2) < agenda.Moyenne && Exist(id2, bxprtAdptr.Intervenants{intervenants}) {
						fmt.Println("Exist")
						liste_ots := lgo.GetLocalisationsByIdTournee(db, tournee.IdTournee)
						var tech = cgo.Intervenant{}
						err = tech.GetIntervenant(db, id2)
						if err != nil {
							fmt.Println(err)
						}
						if VerifyActivityProduit(liste_ots, bxprtAdptr.Intervenant{tech}) {
							fmt.Println("c'est verifié")
							agenda.ResultListIntervenantTournee[id2] = append(agenda.ResultListIntervenantTournee[id2], tournee)
							break
						}
					}
				}

			} else {
				var trnPrevu = contrainteIntervention.TourneePrevu{}
				trnPrevu.GetTourneeByIdTournee(db, tournee.IdTournee)
				if agenda.LastDateByIntervenantAgenda[tournee.IdIntervenant.Int64].Before(trnPrevu.DateIntervention.Time) {
					agenda.LastDateByIntervenantAgenda[tournee.IdIntervenant.Int64] = trnPrevu.DateIntervention.Time
				}
			}
		}
	}

}

func (agenda *Agenda) UpdateLastDate(db *sql.DB) {
	agenda.ListTourneeByIntervenantNotAgendaSupMoy = contrainteIntervention.GetLastTourneeAgenda(db, agenda.LastDateByIntervenantAgenda)
	for id, tournees := range agenda.ListTourneeByIntervenantNotAgendaSupMoy {
		for i, tournee := range tournees {
			last_date := agenda.LastDateByIntervenantAgenda[id]
			New_date := last_date.AddDate(0, 0, i+1)
			fmt.Println(last_date, New_date)
			var trnPrevu = contrainteIntervention.TourneePrevu{}
			trnPrevu.GetTourneeByIdTournee(db, tournee)
			trnPrevu.UpdateAllDateDayTourneePrevu(db, New_date)
		}
	}
}
