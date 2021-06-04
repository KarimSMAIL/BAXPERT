package baxpert

import (
	cgo "Outil_de_gestion_des_tournees/contrainteGeoOptim"
	lgo "Outil_de_gestion_des_tournees/localisationGeoOptim"
	"Outil_de_gestion_des_tournees/missionGeoOptim"
	"database/sql"
	"math"
	"strconv"

	bxprtAdptr "Baxpert_Geo_Optim/baxpertAdapter"
	bxpertRules "Baxpert_Geo_Optim/baxpertRules"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"sort"
	"time"

	"fmt"
)

//Utiliser pour trier les solutions par distance entre OT et Tech
type ByDistance []float64
type UIREs map[string][]string

func (d ByDistance) Len() int {
	return len(d)
}
func (d ByDistance) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}
func (d ByDistance) Less(i, j int) bool {
	return d[i] < d[j]
}

func UpdateAll(list_of_rules []int64, jobContributors bxprtAdptr.JobContributors) {
	var priorite_regle int64
	var intervenants []cgo.Intervenant
	//db, err := sql.Open("mysql", "geooptim2021admin:geo12;optim@tcp(chronograf.fluidicite.xyz:3307)/geooptim2021preprod?parseTime=true")
	db, err := sql.Open("mysql", "geooptim2020admin:optim20geo@tcp(chronograf.fluidicite.xyz:3307)/geooptim2020demo?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err)
	}
	RegleOrganisation := cgo.GetRegleDB(db)

	priorites := cgo.GetPrioriteRegle(db)

	for _, idRegle := range list_of_rules {
		tournees := cgo.GetTourneesByRegle(db, idRegle)
		ui, re := cgo.GetRegleByIdRegle(db, idRegle)

		for _, str := range priorites {
			societe, err := cgo.GetSociete(RegleOrganisation, ui, re, str.TitrePriorite.String)
			if err != nil {
				fmt.Println(err)
			}

			if societe == "" {
				continue
			}
			matricule, err := strconv.Atoi(societe)
			if err == nil {
				intervenants = make([]cgo.Intervenant, 1)
				var interv = cgo.NewIntervenant()
				err = interv.GetIntervenant(db, int64(matricule))

				if err != nil {
					fmt.Println(err)
				}
				intervenants = append(intervenants, interv)
			} else { //avant appel non bufferisé GetIntervenantBySosciete
				intervenants, err = cgo.GetIntervenantBySociete(db, societe)

				if err != nil {
					fmt.Println(err)
				}
			}

			if len(intervenants) > 0 {
				//
				priorite_regle = str.PrioriteRegle
				for _, intervenant := range intervenants {
					var regleIntervenant = cgo.RegleIntervenant{
						IdRegle:       sql.NullInt64{Int64: idRegle, Valid: true},
						PrioriteRegle: sql.NullInt64{Int64: str.PrioriteRegle, Valid: true},
						IdIntervenant: sql.NullInt64{Int64: intervenant.IdIntervenant, Valid: true},
					}
					err = regleIntervenant.InsertRegleIntervenant(db)
					if err != nil {
						fmt.Println(err)
					}
				}
				break
			}
		}

		if err != nil {
			fmt.Println(err)
		}

		for _, id_tournee := range tournees {
			//update Regle Localisation
			err := cgo.UpdatePrioriteRegleRegleLocalisation(db, id_tournee, priorite_regle)
			if err != nil {
				fmt.Println(err)
			}
			//recuperer la liste des OTs
			liste_ots := lgo.GetLocalisationsByIdTournee(db, id_tournee)
			//liste_ots := []lgo.Localisation{}

			//verifier l'activite des ots avec specialités des intervenants
			intervenants_bis := Verify_Activity(liste_ots, bxprtAdptr.Intervenants{intervenants})
			//fmt.Println("intervenants_bis : ", len(intervenants_bis.Intervenants))

			//recuperer les jobContributors par regle
			jobcontributors := bxprtAdptr.GetJobContributorsBySkill(intervenants_bis, jobContributors)
			//fmt.Println("jobcontributors : ", len(jobcontributors.JobContributorUserAgents))

			//fmt.Println("id regle :", idRegle, "nombre des jobContributors :", len(jobcontributors.JobContributorUserAgents))
			bxprtAdptr.UpdateDumbContributorsBySkill(id_tournee, jobcontributors, priorite_regle)
			//bxprtAdptr.UpdateDumbContributorsByDate()
		}
	} // fin des mise à jour des intervenants des tournée.
}

func Verify_Activity(liste_ots []lgo.Localisation, liste_intervenants bxprtAdptr.Intervenants) bxprtAdptr.Intervenants {
	var intervenants bxprtAdptr.Intervenants
	var tech_ok bool
	//fmt.Println("intervenants : ", len(liste_intervenants.Intervenants))
	for _, tech := range liste_intervenants.Intervenants {
		Qualite := bxprtAdptr.GetJobContributorSkill(tech.QualiteIntervenantJson)
		//fmt.Println("ots : ", len(liste_ots))
		for _, ot := range liste_ots {
			tech_ok = false
			Activiteproduit := bxprtAdptr.GetActiviteProduit(ot.TypeJson.String)
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

type ContrainteOTJour struct {
	EffectifIntervenant int
	NombreIntervention  int
}

//activite effectifOT
var ActiviteNbrOT = make(map[string]int)

func BuildRegleContrainteOTJour(wo missionGeoOptim.JobWO) {
	fmt.Printf("%+v,%+v,%+v,%+v,%+v \n", wo.Duree_probable, wo.Activite, wo.Code_UI, wo.Code_centre, wo.Produit)
	if wo.Duree_probable.Valid {
		if wo.Activite.Valid {
			a, err := time.Parse("15:04:05", wo.Duree_probable.String)
			i := a.Hour()*60 + a.Minute()
			fmt.Printf("%d, %v\n", i, err)
			if err == nil {
				ActiviteNbrOT[wo.Activite.String] = int(math.Round(5 * 60 / float64(i)))
			}
		}
	}
}

func Regle_Intervenats() map[bxprtAdptr.Regle]ContrainteOTJour {

	var regles_existe = make(map[bxprtAdptr.Regle]ContrainteOTJour)
	var intervenants []cgo.Intervenant

	jobWos := bxprtAdptr.JobWOs{}
	//dsn := "geooptim2021admin:geo12;optim@tcp(chronograf.fluidicite.xyz:3307)/geooptim2021preprod?parseTime=true"
	dsn := "geooptim2020admin:optim20geo@tcp(chronograf.fluidicite.xyz:3307)/geooptim2020demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//jobWos = GetJobWOs(db, " ")
	//On doit mettre le nom de l'agence en parametre
	//jobWos = GetJobWOsAgence(db, "LYON Clients FTTH")
	//jobWos = GetJobWOsAgence(db, "ALPES Clients Cuivre")
	jobWos = GetJobWOsAgence(db, "LYON Clients Cuivre")

	jobContributors := bxprtAdptr.JobContributors{}

	jobContributors.JobContributorUserAgents = GetJobContributors(db)

	//defer db.Close()

	//db2, err := sql.Open("mysql", "geooptim2021admin:geo12;optim@tcp(chronograf.fluidicite.xyz:3307)/geooptim2021preprod?parseTime=true")
	db2, err := sql.Open("mysql", "geooptim2020admin:optim20geo@tcp(chronograf.fluidicite.xyz:3307)/geooptim2020demo?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	RegleOrganisation := cgo.GetRegleDB(db2)

	priorites := cgo.GetPrioriteRegle(db2)
	for _, wo := range jobWos.Ots {
		BuildRegleContrainteOTJour(wo)
	}
	for _, wo := range jobWos.Ots { // Pour chaque OT
		regle := bxprtAdptr.Regle{Ui: wo.JobConstraint.Code_UI.String, Re: wo.JobConstraint.Code_centre.String, Depot: wo.JobConstraint.Depot.String}
		//if wo.JobConstraint.Agence.String == "ALPES Clients Cuivre" {
		for _, str := range priorites {
			societe, err := cgo.GetSociete(RegleOrganisation, wo.JobConstraint.Code_UI.String, wo.JobConstraint.Code_centre.String, str.TitrePriorite.String)

			if err != nil {
				fmt.Println(err)
			}

			intervenants, err = cgo.GetIntervenantBySosciete(db2, societe)

			if err != nil {
				fmt.Println(err)
			}

			if len(intervenants) > 0 {
				//
				//priorite_regle = str.PrioriteRegle
				break
			}
		}

		jobcontributors := bxprtAdptr.GetJobContributorsBySkill(bxprtAdptr.Intervenants{intervenants}, jobContributors)
		if len(jobcontributors.JobContributorUserAgents) > 0 && regles_existe[regle].EffectifIntervenant == 0 {
			contrainteOTJour := regles_existe[regle]
			contrainteOTJour.EffectifIntervenant += len(jobcontributors.JobContributorUserAgents)
			contrainteJ, existJ := ActiviteNbrOT[wo.Activite.String]
			if existJ {
				contrainteOTJour.NombreIntervention = contrainteJ
			} else {
				contrainteOTJour.NombreIntervention = 5
			}
			regles_existe[regle] = contrainteOTJour
		}
		//}
	}
	//defer db.Close()

	return regles_existe
}

//Get Wos knowing that Wo.Etat_interne == "non affecte"
func GetJobWOs(db *gorm.DB, apiKey string) bxprtAdptr.JobWOs {
	jobWOsBis := missionGeoOptim.JobWOs{}
	_ = jobWOsBis.GetJobWOs(db)

	return bxprtAdptr.JobWOs{jobWOsBis}
}

//Get Wos knowing that Wo.Etat_interne == "non affecte"
func GetJobWOsAgence(db *gorm.DB, name_agence string) bxprtAdptr.JobWOs {
	jobWOsBis := missionGeoOptim.JobWOs{}
	_ = jobWOsBis.GetJobWOsAgence(db, name_agence)

	return bxprtAdptr.JobWOs{jobWOsBis}
}

//create JobWOs statics for the test
func GetStaticJobWOs() bxprtAdptr.JobWOs {
	jobWOs := bxprtAdptr.JobWOs{}
	for i := 0; i < 20; i++ {
		code := "code" + strconv.Itoa(i+1)
		jobInfo := missionGeoOptim.JobInfo{}
		jobInfo.Code_base.String = code
		jobInfo.Numero_intervention.String = fmt.Sprint(i)
		jobInfo.Latitude.String = fmt.Sprint(5678 + float64(i))
		jobInfo.Longitude.String = fmt.Sprint(345678 + float64(i))
		jobInfo.Etat_interne.String = "non affecte"

		adr := "adresse" + strconv.Itoa(i+1)
		activite := "activite" + strconv.Itoa(i%3+1)
		layout := "2006-01-02"
		str := string("2021-04-" + strconv.Itoa(1+i))
		jobConstraint := missionGeoOptim.JobConstraint{}
		jobConstraint.Adresse.String = adr
		jobConstraint.Activite.String = activite
		jobConstraint.Date_intervention.Time, _ = time.Parse(layout, str)
		jobConstraint.Duree_probable.String = "2h"

		jobReissue := missionGeoOptim.JobReissue{}
		jobSafety := missionGeoOptim.JobSafety{}
		jobClient := missionGeoOptim.JobClient{}

		jobWo := missionGeoOptim.JobWO{jobInfo, jobConstraint, jobReissue, jobSafety, jobClient}
		jobWOs.Ots = append(jobWOs.Ots, jobWo)
	}
	return jobWOs
}

//Get the JobContributors from the db
func GetJobContributors(db *gorm.DB) []missionGeoOptim.JobContributorUserAgent {
	jobContributorsBis := missionGeoOptim.JobContributors{}
	_ = jobContributorsBis.GetJobContributors(db)

	return jobContributorsBis.JobContributorUserAgents
}

//create JobContributors statics for the test
func GetStaticJobConributors() bxprtAdptr.JobContributors {
	jobContributors := bxprtAdptr.JobContributors{}
	jobContributors.JobContributorUserAgents = make([]missionGeoOptim.JobContributorUserAgent, 20)
	for i := 0; i < len(jobContributors.JobContributorUserAgents); i++ {
		jobContributors.JobContributorUserAgents[i].Numero_utilisateur = int64(i + 1)
		jobContributors.JobContributorUserAgents[i].Latitude_depart.String = fmt.Sprint(78 + float64(i))
		jobContributors.JobContributorUserAgents[i].Longitude_depart.String = fmt.Sprint(67 + float64(i))
	}
	return jobContributors
}

//Définir manuellement les activités des technicien
func Set_Activite(jobContributors bxprtAdptr.JobContributors) {
	mapActivite := map[int]string{0: "PLI", 1: "S02", 2: "R02", 3: "EIF", 4: "PRE", 5: "CLG", 6: "IQL", 7: "IHL", 8: "IQP", 9: "PFC",
		10: "PMS", 11: "PFO", 12: "PCC", 13: "LML", 14: "LMR", 15: "F12", 16: "R10", 17: "IQC", 18: "IHD",
		19: "LMD", 20: "LMP", 21: "DET", 22: "FO3", 23: "R03", 24: "MSH", 25: "ITL", 26: "MIT"}

	//Définition les activités
	for i, tech := range jobContributors.JobContributorUserAgents {
		jobContributors.Activite[tech.Numero_utilisateur] = mapActivite[i%27]
	}
}

//Définir manuellement les disponibilités des technicien
func Set_Disponibilte(jobContributors bxprtAdptr.JobContributors) {
	for i, tech := range jobContributors.JobContributorUserAgents {
		layout := "2006-01-02"
		str := string("2021-01-" + strconv.Itoa(1+i))
		tDeb, _ := time.Parse(layout, str)
		tFin, _ := time.Parse(layout, "2021-12-30")
		//Création des disponibilités
		jobContributors.Dispos[tech.Numero_utilisateur] = bxprtAdptr.Disponibilites{IdDisponibilite: int64(i + 1), DateDisponibiliteDebut: sql.NullTime{Time: tDeb, Valid: true}, DateDisponibiliteFin: sql.NullTime{Time: tFin, Valid: true}}
	}
}

func Calcul_Distance(OT bxprtAdptr.JobWO, Tech bxprtAdptr.JobContributorUserAgent) float64 {
	//convertir les coordonnées geo de string en float
	latDebut, _ := strconv.ParseFloat(Tech.Latitude_depart.String, 64)
	longDebut, _ := strconv.ParseFloat(Tech.Longitude_depart.String, 64)
	lat, _ := strconv.ParseFloat(OT.Latitude.String, 64)
	long, _ := strconv.ParseFloat(OT.Longitude.String, 64)

	//distance(A,B)=√((x2−x1)2+(y2−y1)2)
	return math.Sqrt(math.Pow(latDebut-lat, 2) + math.Pow(longDebut-long, 2))
}

//Trier les solution par distance
func Sort_Solution(solutions_Bis bxprtAdptr.Solutions, distances []float64) bxprtAdptr.Solutions {
	solutions := make(bxprtAdptr.Solutions)
	//trier les solutions par distance
	sort.Sort(ByDistance(distances))
	for i, dist := range distances {
		for _, sol := range solutions_Bis {
			if Calcul_Distance(sol.OT, sol.Tech) == dist {
				solutions[int64(i)] = sol
				break
			}
		}
	}
	return solutions
}

//Le meilleur score
func Get_BestSolutionByScore(solutions bxprtAdptr.Solutions) bxprtAdptr.Solution {
	best_solution := solutions[0]

	for _, value := range solutions {

		if value.Score > best_solution.Score {
			best_solution = value
		}
	}
	return best_solution
}

//Le 2eme meilleur score
func Get_BestSolutionByScore_Bis(solutions bxprtAdptr.Solutions) bxprtAdptr.Solution {
	best_solution := Get_BestSolutionByScore(solutions)
	if len(solutions) > 1 {
		for key, value := range solutions {
			if value.Score == best_solution.Score {
				solutions[key] = bxprtAdptr.Solution{Score: int64(0)}
				break
			}
		}
	}
	return Get_BestSolutionByScore(solutions)
}

func Make_OT_Tech(jobwo_tech bxprtAdptr.NumOT_NumTech, jobWos bxprtAdptr.JobWOs, jobContributors bxprtAdptr.JobContributors) bxprtAdptr.OT_Tech {
	OTs_Techs := make(bxprtAdptr.OT_Tech)
	for num_intervention, num_tech := range jobwo_tech {
		for _, jobwo := range jobWos.Ots {
			if jobwo.JobInfo.Numero_intervention.String == num_intervention {
				for _, tech := range jobContributors.JobContributorUserAgents {
					if tech.Numero_utilisateur == num_tech {
						jobInfo := bxprtAdptr.JobInfo{jobwo.JobInfo}
						jobConstraint := bxprtAdptr.JobConstraint{jobwo.JobConstraint}
						jobReissue := bxprtAdptr.JobReissue{jobwo.JobReissue}
						jobSafety := bxprtAdptr.JobSafety{jobwo.JobSafety}
						jobClient := bxprtAdptr.JobClient{jobwo.JobClient}
						Ot := bxprtAdptr.JobWO{jobInfo, jobConstraint, jobReissue, jobSafety, jobClient}
						OTs_Techs[Ot] = bxprtAdptr.JobContributorUserAgent{tech}
					}
				}
			}
		}
	}
	return OTs_Techs
}

func GetIntervenantsByRegle(jobWos bxprtAdptr.JobWOs) bxprtAdptr.Regle_Intervenants {
	var regle_intervenants = make(bxprtAdptr.Regle_Intervenants)

	//db, err := sql.Open("mysql", "geooptim2021admin:geo12;optim@tcp(chronograf.fluidicite.xyz:3307)/geooptim2021preprod?parseTime=true")
	db, err := sql.Open("mysql", "geooptim2020admin:optim20geo@tcp(chronograf.fluidicite.xyz:3307)/geooptim2020demo?charset=utf8mb4&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
	RegleOrganisation := cgo.GetRegleDB(db)
	for _, wo := range jobWos.Ots { // Pour chaque OT

		societe, err := cgo.GetSociete(RegleOrganisation, wo.JobConstraint.Code_UI.String, wo.JobConstraint.Code_centre.String, "P1")

		if err != nil {
			fmt.Println(err)
		}

		intervenants, err := cgo.GetIntervenantBySosciete(db, societe)

		if err != nil {
			fmt.Println(err)
		}

		regle := bxprtAdptr.Regle{Ui: wo.JobConstraint.Code_UI.String, Re: wo.JobConstraint.Code_centre.String}
		regle_intervenants[regle] = bxprtAdptr.Intervenants{intervenants}
	}
	return regle_intervenants
}

func Max_Regle(m map[bxprtAdptr.Regle]int) (bxprtAdptr.Regle, int) {
	var max_regle bxprtAdptr.Regle
	max := 0
	for key, value := range m {
		if value > max {
			max = value
			max_regle = key
		}
	}

	return max_regle, max
}

//Etablir la liste des intervenants possibles pour chaque JobWO (vérification de : dispo, activité)
func Make_sol(jobwo_tech bxprtAdptr.NumOT_NumTech, jobWos bxprtAdptr.JobWOs,
	jobContributors bxprtAdptr.JobContributors, constraints *bxpertRules.Constraint, Kind_of_solution int) bxprtAdptr.OT_Tech {
	fmt.Println(" ")
	//var uires = make(UIREs)
	var Best_sol bxprtAdptr.Solution
	var regles = make(map[bxprtAdptr.Regle]int)
	fmt.Println(regles)
	//constraintsList[0] : RuleWoNotAssigned
	//constraintsList[1] : RulePathFinding
	//constraintsList[2] : RuleAvailabilityIntervenor
	//constraintsList[3] : RuleAvailabilityClient
	//constraintsList[4] : RuleAvailabilitySite

	var ruleWoNotAssigned bxpertRules.RuleWoNotAssigned
	//var rulePathFinding bxpertRules.RulePathFinding
	var ruleAvailabilityIntervenor bxpertRules.RuleAvailabilityIntervenor
	var ruleAvailabilityClient bxpertRules.RuleAvailabilityClient
	//var ruleAvailabilitySite bxpertRules.RuleAvailabilitySite

	var solutions bxprtAdptr.Solutions
	var distances []float64
	i := int64(0)

	//db, err := sql.Open("mysql", "geooptim2021admin:geo12;optim@tcp(chronograf.fluidicite.xyz:3307)/geooptim2021preprod?parseTime=true")
	db, err := sql.Open("mysql", "geooptim2020admin:optim20geo@tcp(chronograf.fluidicite.xyz:3307)/geooptim2020demo?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	RegleOrganisation := cgo.GetRegleDB(db)
	fmt.Println("lenth =", len(jobWos.Ots))

	for _, wo := range jobWos.Ots { // Pour chaque OT
		regle := bxprtAdptr.Regle{Ui: wo.JobConstraint.Code_UI.String, Re: wo.JobConstraint.Code_centre.String}
		regles[regle] += 1
		if wo.JobConstraint.Agence.String == "ALPES Clients Cuivre" {
			//recuperer le idRegle ?? UI/RE de l'OT
			//recuperer les intervenants pour l'OT => List = intervenants

			societe, err := cgo.GetSociete(RegleOrganisation, wo.JobConstraint.Code_UI.String, wo.JobConstraint.Code_centre.String, "P1")

			if err != nil {
				fmt.Println(err)
			}

			/*if len(uire) > 0 {
				//recupérer les regles qui n'existe pas
				uires[uire[0]] = append(uires[uire[0]], uire[1])
			}*/

			intervenants, err := cgo.GetIntervenantBySosciete(db, societe)

			if err != nil {
				fmt.Println(err)
			}
			//on récupere la liste des jobContributors pour faire marcher ProcessConstraint de chaque regle
			jobcontributors := bxprtAdptr.GetJobContributorsBySkill(bxprtAdptr.Intervenants{intervenants}, jobContributors)
			//fmt.Println("lenth jobContributors =", len(jobcontributors.JobContributorUserAgents))
			i = 0
			solutions = make(bxprtAdptr.Solutions)
			distances = make([]float64, len(jobcontributors.JobContributorUserAgents))
			for _, tech := range jobcontributors.JobContributorUserAgents {
				constraints.GlobalScore = 0
				//if jobContributors.Libre[tech.Numero_utilisateur] {
				w_o := bxprtAdptr.JobWO{bxprtAdptr.JobInfo{wo.JobInfo}, bxprtAdptr.JobConstraint{wo.JobConstraint},
					bxprtAdptr.JobReissue{wo.JobReissue}, bxprtAdptr.JobSafety{wo.JobSafety},
					bxprtAdptr.JobClient{wo.JobClient}}

				constraints.WO = w_o
				constraints.Tech = bxprtAdptr.JobContributorUserAgent{tech}

				//Constraint Dispo/Activite
				ruleAvailabilityIntervenor.Convert(constraints.ConstraintList[2])
				ruleAvailabilityIntervenor.Start_date = jobContributors.Dispos[tech.Numero_utilisateur].DateDisponibiliteDebut
				ruleAvailabilityIntervenor.End_date = jobContributors.Dispos[tech.Numero_utilisateur].DateDisponibiliteFin
				ruleAvailabilityIntervenor.Intervention_date = wo.JobConstraint.Date_intervention
				ruleAvailabilityIntervenor.Activite_OT = wo.JobConstraint.Activite.String
				ruleAvailabilityIntervenor.Activite_Tech = jobContributors.Activite[tech.Numero_utilisateur]
				ruleAvailabilityIntervenor.ProcessConstraint(constraints)

				//Constraint OT Non affecte
				ruleWoNotAssigned.Convert(constraints.ConstraintList[0])
				ruleWoNotAssigned.Internal_state = wo.JobInfo.Etat_interne
				ruleWoNotAssigned.ProcessConstraint(constraints)

				//Constraint Availability Client
				ruleAvailabilityClient.Convert(constraints.ConstraintList[3])
				ruleAvailabilityClient.RdV_Fixe = wo.JobConstraint.RdV_Fixe
				ruleAvailabilityClient.RDV_present = wo.JobClient.RDV_present
				ruleAvailabilityClient.ProcessConstraint(constraints)

				/*//Constraint Availability Site
				ruleAvailabilitySite.Convert(constraints.ConstraintList[3])
				ruleAvailabilitySite.ProcessConstraint(constraints)*/

				solutions[i] = bxprtAdptr.Solution{w_o, bxprtAdptr.JobContributorUserAgent{tech}, int64(constraints.GlobalScore)}
				distances[i] = Calcul_Distance(w_o, bxprtAdptr.JobContributorUserAgent{tech})
				i += 1
			}

			//générer les solutions possible trié par distance
			solutions = Sort_Solution(solutions, distances)

			if Kind_of_solution == 1 { // 1 => la bonne solution
				//retourner la meilleur solution selon les scores
				Best_sol = Get_BestSolutionByScore(solutions) //{wo, tech, score}

			} else {
				//retourner la 2éme meilleur solution selon les scores
				Best_sol = Get_BestSolutionByScore_Bis(solutions) //{wo, tech, score}
			}

			jobwo_tech[wo.JobInfo.Numero_intervention.String] = Best_sol.Tech.Numero_utilisateur

		}
	}

	//Ecrire les regle qui n'existe pas
	/*fmt.Println("liste des regles qui n'existe pas")
	for key, value := range uires {
		fmt.Println(key, ": ", len(value))
	}*/

	fmt.Println(Max_Regle(regles))

	return Make_OT_Tech(jobwo_tech, jobWos, jobContributors)
}

func Make_Lot(OT_Tech bxprtAdptr.OT_Tech) bxprtAdptr.Lot {
	var Lots = make(bxprtAdptr.Lot)
	var regle bxprtAdptr.Regle
	for key, value := range OT_Tech {
		regle = bxprtAdptr.Regle{Ui: key.JobConstraint.Code_UI.String, Re: key.JobConstraint.Code_centre.String}

		Lots[regle] = append(Lots[regle], bxprtAdptr.Solution{key, value, 0})
	}

	return Lots
}
