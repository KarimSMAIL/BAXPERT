package main

import (
	bxpAdtr "Baxpert_Geo_Optim/baxpertAdapter"
	cgo "Outil_de_gestion_des_tournees/contrainteGeoOptim"

	bxpt "Baxpert_Geo_Optim/baxpert"

	"database/sql"
	"fmt"

	//bxpExprt "Baxpert_Geo_Optim/baxpertExporter"

	bxpertRules "Baxpert_Geo_Optim/baxpertRules"

	//"time"

	//"math"
	//"strconv"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestGetOptim(t *testing.T) {

	p := fmt.Println
	p("")
	//var OTs_Techs bxpAdtr.OT_Tech

	rule0 := bxpertRules.NewRuleWoNotAssigned(sql.NullString{String: " ", Valid: false})
	rule1 := bxpertRules.NewRulePathFinding()
	rule2 := bxpertRules.NewRuleAvailabilityIntervenor()
	rule3 := bxpertRules.NewRuleAvailabilityClient()
	rule4 := bxpertRules.NewRuleAvailabilitySite()

	var constraints bxpertRules.Constraint

	//constraints.Title = "Constraints"
	constraints.AssignRuleIntoConstraint(rule0)
	constraints.AssignRuleIntoConstraint(rule1)
	constraints.AssignRuleIntoConstraint(rule2)
	constraints.AssignRuleIntoConstraint(rule3)
	constraints.AssignRuleIntoConstraint(rule4)

	//constraintss :=
	/*p(constraints.ConstraintList[0])
	p(constraints.ConstraintList[2])*/

	//définition des JobWOs
	//jobWos := bxpAdtr.JobWOs{}
	dsn := "geooptim2021admin:geo12;optim@tcp(chronograf.fluidicite.xyz:3307)/geooptim2021preprod?parseTime=true"
	//dsn := "geooptim2020admin:optim20geo@tcp(chronograf.fluidicite.xyz:3307)/geooptim2020demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//jobWos = bxpt.GetJobWOs(db) //(db, apiKey)
	//jobWos = bxpt.GetStaticJobWOs()


	//déinition des techniciens
	jobContributors := bxpAdtr.JobContributors{}

	jobContributors.JobContributorUserAgents = bxpt.GetJobContributors(db)

	//jobContributors = bxpt.GetStaticJobConributors()
	jobContributors.Dispos = make(map[int64]bxpAdtr.Disponibilites)
	jobContributors.Activite = make(map[int64]string)

	//Définition des disponibilités et les activités
	bxpt.Set_Activite(jobContributors)
	bxpt.Set_Disponibilte(jobContributors)

	/*for key, value := range bxpt.Regle_Intervenats() {
		p("regle : ", key, "=>Nombre de techniciens :", value)
	}*/

	db2, err := db.DB() //sql.Open("mysql", "geooptim2021admin:geo12;optim@tcp(chronograf.fluidicite.xyz:3307)/geooptim2021preprod?parseTime=true")
	if err != nil {
		fmt.Println(err)
	}
	RegleOrganisation := cgo.GetRegleDB(db2)

	//defer db2.Close()

	//var list_of_rules = []int64{100, 433, 505}

	Regles := bxpt.Regle_Intervenats()
	var list_of_rules = make([]int64, len(Regles))
	var i = 0

	for key, _ := range Regles{
		list_of_rules[i] = cgo.GetIdRegle(RegleOrganisation, key.Ui, key.Re)
		i += 1
	}
	
	bxpt.UpdateAll(list_of_rules, jobContributors)

	//la liste des jobwo avec le technicien associé
	/*jobwo_tech := make(bxpAdtr.NumOT_NumTech)
	Kind_of_solution := 1 //1 = La bonne solution
	OTs_Techs = bxpt.Make_sol(jobwo_tech, jobWos, jobContributors, &constraints, Kind_of_solution)
	//Lots := bxpt.Make_Lot(OTs_Techs)
	//p(Lots)

	//WriteFile xls
	bxpExprt.WriteExcelFile("./Resultat.xlsx", OTs_Techs)
	p("Succes generate of Xlsx file")*/

	//affichage
	/*for key, value := range OTs_Techs {
		p("Numéro OT : ", key.JobInfo.Numero_intervention.String, "=>Nom technicien :", value.Nom.String)
	}*/

}
