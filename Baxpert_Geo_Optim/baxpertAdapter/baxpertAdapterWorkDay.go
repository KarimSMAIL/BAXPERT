package baxpertAdapter

import (
	"Outil_de_gestion_des_tournees/contrainteGeoOptim"
	"Outil_de_gestion_des_tournees/contrainteIntervention"
	lgo "Outil_de_gestion_des_tournees/localisationGeoOptim"

	//mgo "Outil_de_gestion_des_tournees/missionGeoOptim"
	"database/sql"
	"fmt"
)

//fait la mise à jour des intervenant - décale les dates
func UpdateDumbContributorsByDate(jobContributors JobContributors) error {
	var last_date = make(map[int64]int) // map[id_intervenant]nombre de jour à décaler
	//GetList
	//db, err := sql.Open("mysql", "snipeuser:snipeprod998A;@tcp(chronograf.fluidicite.xyz:3307)/snipe_geo?parseTime=true")
	db, err := sql.Open("mysql", "geooptim2020admin:optim20geo@tcp(chronograf.fluidicite.xyz:3307)/geooptim2020demo?parseTime=true")
	//db, err := sql.Open("mysql", "geooptim2021admin:geo12;optim@tcp(chronograf.fluidicite.xyz:3307)/geooptim2021preprod?parseTime=true")
	defer db.Close()
	//	var dispoIntervenants = contrainteGeoOptim.GetDisponibiliteByIntervenant(db)
	//recuperer la liste des tournees en OverBook
	tourneeOverBook, err := contrainteGeoOptim.GetListIntervenantOverbookTournee(db) //id_tournee
	if err != nil {
		fmt.Println(err)
	}
	for _, id_tournee := range tourneeOverBook.ListIdTournee { //on parcours les tournees
		//parcourir la liste des regles associé à cette tournee (à priori == 1)
		var trnPrevu = contrainteIntervention.TourneePrevu{}
		trnPrevu.GetTourneeByIdTournee(db, id_tournee)
		date := trnPrevu.DateIntervention.Time
		for _, rgl := range tourneeOverBook.RegleIntervenantListTournee[id_tournee] {
			//recuperer la liste des intervenants associé à la regle
			intervenants, err := contrainteGeoOptim.GetIntervenantByRegleDate(db, rgl, date)
			if err != nil {
				fmt.Println(err)
			}
			//recuperer la liste des intervenants not overbook
			intervenants_NotOverBook := IntervenantsNotOverBook(Intervenants{intervenants}, tourneeOverBook)
			intervenants_Dispo := contrainteGeoOptim.DispoIntervenant{intervenants_NotOverBook.Intervenants}

			//Test Activite/Produit
			liste_ots := lgo.GetLocalisationsByIdTournee(db, id_tournee)
			intervenants_bis := Verify_Activity_Produit(liste_ots, Intervenants{intervenants_Dispo.Intervenants})

			jobcontributors := GetJobContributorsBySkill(intervenants_bis, jobContributors)
			//fmt.Println("map 1 :jobcontributors : ", len(jobcontributors.JobContributorUserAgents))

			if len(jobcontributors.JobContributorUserAgents) > 0 {
				//Substitution
				//
				PutIntervenant(rgl.PrioriteRegle.Int64, rgl.IdRegle.Int64, id_tournee, rgl.IdIntervenant.Int64, jobcontributors.JobContributorUserAgents[len(jobcontributors.JobContributorUserAgents)-1].Numero_utilisateur)
				//mgo.UpdateIntervenantByIdTournee(db, id_tournee, jobcontributors.JobContributors, rgl.PrioriteRegle.Int64, rgl.IdRegle.Int64)
			} else {
				// //Etalement
				date = trnPrevu.DateIntervention.Time.AddDate(0, 0, last_date[rgl.IdIntervenant.Int64])
				trnPrevu.UpdateAllDateDayTourneePrevu(db, date)
				last_date[rgl.IdIntervenant.Int64] += 1
			}
		}
	}

	return err
}

func IntervenantsNotOverBook(intervenants Intervenants, tourneeOverBook *contrainteGeoOptim.ListIntervenantTourneeOverbook) Intervenants {
	intervenants_Bis := Intervenants{}
	for _, inter := range intervenants.Intervenants {
		tech_No_ok := false
		for _, inter_overbook := range tourneeOverBook.ListIdIntervenant {
			if inter.IdIntervenant == inter_overbook {
				tech_No_ok = true
				break
			}
		}
		if tech_No_ok == false {
			intervenants_Bis.Intervenants = append(intervenants_Bis.Intervenants, inter)
		}
	}
	return intervenants_Bis
}
func Verify_Activity_Produit(liste_ots []lgo.Localisation, liste_intervenants Intervenants) Intervenants {
	var intervenants Intervenants
	var tech_ok bool
	//fmt.Println("intervenants : ", len(liste_intervenants.Intervenants))
	for _, tech := range liste_intervenants.Intervenants {
		Qualite := GetJobContributorSkill(tech.QualiteIntervenantJson)
		//fmt.Println("ots : ", len(liste_ots))
		for _, ot := range liste_ots {
			tech_ok = false
			Activiteproduit := GetActiviteProduit(ot.TypeJson.String)
			for key, _ := range Qualite.Specialite {
				for key2, _ := range Qualite.Produit {
					//fmt.Println("activite ot :", Activiteproduit.Activite, "activite tech :", key)
					if Activiteproduit.Activite == key && Activiteproduit.Produit == key2 {
						//fmt.Println("activite true")
						tech_ok = true
						break
					}
				}
			}
			if tech_ok == false {
				break
			}
		}
		if tech_ok == true {
			intervenants.Intervenants = append(intervenants.Intervenants, tech)
		}
	}
	//fmt.Println("intervenants : ", len(intervenants.Intervenants))
	return intervenants
}

func Exist_Intervenant(idIntervenant int64, intervenants_Dispos contrainteGeoOptim.DispoIntervenant) bool {
	for _, intervenant := range intervenants_Dispos.Intervenants {
		if idIntervenant == intervenant.IdIntervenant {
			return true
		}
	}
	return false
}
