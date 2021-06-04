// main
package main

import (
	bxpt "Baxpert_Geo_Optim/baxpert"
	bxpAdtr "Baxpert_Geo_Optim/baxpertAdapter"

	bxpertAgenda "Baxpert_Geo_Optim/baxpertAgenda"
	cgo "Outil_de_gestion_des_tournees/contrainteGeoOptim"
	"fmt"

	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s durée %v\n", what, time.Since(start))
	}
}
func main() {
	defer elapsed("baxpert")()
	//dsn := "geooptim2021admin:geo12;optim@tcp(chronograf.fluidicite.xyz:3307)/geooptim2021preprod?parseTime=true"
	dsn := "geooptim2020admin:optim20geo@tcp(chronograf.fluidicite.xyz:3307)/geooptim2020demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//déinition des techniciens
	jobContributors := bxpAdtr.JobContributors{}

	jobContributors.JobContributorUserAgents = bxpt.GetJobContributors(db)

	db2, err := db.DB() //sql.Open("mysql", "geooptim2021admin:geo12;optim@tcp(chronograf.fluidicite.xyz:3307)/geooptim2021preprod?parseTime=true")
	if err != nil {
		fmt.Println(err)
	}
	RegleOrganisation := cgo.GetRegleDB(db2)

	Regles := bxpt.Regle_Intervenats()
	var list_of_rules = make([]int64, len(Regles))
	var i = 0

	for key, _ := range Regles {
		list_of_rules[i] = cgo.GetIdRegle(RegleOrganisation, key.Ui, key.Re)
		i += 1
	}

	bxpt.UpdateAll(list_of_rules, jobContributors)
	// Substitution And Etalement
	err = bxpAdtr.UpdateDumbContributorsByDate(jobContributors)
	if err != nil {
		fmt.Println(err)
	}
	//	Planing
	agenda := bxpertAgenda.Agenda{}
	agenda.InitAgenda(db2)
	agenda.Planing(db2)
	agenda.PutRegleIntervenantTournee(db2)
	agenda.UpdateLastDate(db2)
}
