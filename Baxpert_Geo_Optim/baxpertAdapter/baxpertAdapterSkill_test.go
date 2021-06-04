package baxpertAdapter

import (
	//bxpt "Baxpert_Geo_Optim/baxpert"
	"database/sql"
	"fmt"
	"testing"
	/*"gorm.io/driver/mysql"
	"gorm.io/gorm"*/)

func TestGetJobContributorSkill(t *testing.T) {
	str := sql.NullString{String: `{"n":2933,"t":"ST","soc":"01COM","e":"TECH","typo":"SAV/PROD","cnt":0,"spe":null}`, Valid: true}
	jobContributorSkill := GetJobContributorSkill(str)
	fmt.Println(jobContributorSkill)
}

func TestGetStaticIntervenant(t *testing.T) {
	intervenants := GetStaticIntervenant()
	for _, i := range intervenants {
		fmt.Println("id Intervenant = ", i.IdIntervenant)
	}
}

func TestGetJobContributorsBySkill(t *testing.T) {

	/*dsn := "geooptim2020admin:optim20geo@tcp(chronograf.fluidicite.xyz:3307)/geooptim2020demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})*/
	/*
		jobContributors := JobContributors{}
		//jobContributors.JobContributorUserAgents = bxpt.GetJobContributors(db)

		intervenants := GetStaticIntervenant()

		jobcontributors := GetJobContributorsBySkill(Intervenants{intervenants}, jobContributors)
		for _, i := range jobcontributors.JobContributorUserAgents {
			fmt.Println("Nom Tech = ", i.Nom)
		}
	*/
}
