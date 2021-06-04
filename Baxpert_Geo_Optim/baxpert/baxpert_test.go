package baxpert

import (
	"fmt"
	bxprtAdptr "Baxpert_Geo_Optim/baxpertAdapter"
	/*lgo "Outil_de_gestion_des_tournees/localisationGeoOptim"
	cgo "Outil_de_gestion_des_tournees/contrainteGeoOptim"*/
	//"database/sql"
	/*"math/rand"
	"strconv"*/
	"testing"
	/*"gorm.io/gorm"
	"gorm.io/driver/mysql"*/
	//"time"
)

func TestRegle_Intervenats(t *testing.T) {
	m := Regle_Intervenats()
	r := bxprtAdptr.Regle{Ui : "PEM", Re : "HOU"}

	fmt.Println(m[r])
}

func TestUpdateAll(t *testing.T) {
	
}


/*func TestLen(t *testing.T) {
	d := ByDistance([]float64{1.0, 2.0, 3.0})
	if d.Len() == 3 {
		fmt.Println("Len is ok")
	}
}

func TestSwap(t *testing.T) {
	d := ByDistance([]float64{1.0, 2.0, 3.0})
	d.Swap(0, 1)
	if d[0] == 2.0 && d[1] == 1.0 {
		fmt.Println("Swap is ok")
	}
}

func TestLess(t *testing.T) {
	d := ByDistance([]float64{1.0, 2.0, 3.0})
	if d.Less(0, 1) {
		fmt.Println("Less is ok")
	}
}

func TestGetIntervenantsByRegle(t *testing.T) {
	jobWos := bxprtAdptr.JobWOs{}
	dsn := "geooptim2021admin:geo12;optim@tcp(chronograf.fluidicite.xyz:3307)/geooptim2021preprod?parseTime=true"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	jobWos = GetJobWOs(db, " ") //(db, apiKey)

	intervenants_regles := GetIntervenantsByRegle(jobWos)

	fmt.Println("liste des intervenats par regles")
	for key, value := range intervenants_regles {
		fmt.Println(key, ": ", len(value.Intervenants))
	}
}

func TestVerify_Activity(t *testing.T){

	dsn := "geooptim2021admin:geo12;optim@tcp(chronograf.fluidicite.xyz:3307)/geooptim2021preprod?parseTime=true"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//déinition des techniciens
	jobContributors := bxprtAdptr.JobContributors{}

	jobContributors.JobContributorUserAgents = GetJobContributors(db)

	db1, err := db.DB()

	if err != nil {
		fmt.Println(err)
	}

	RegleOrganisation := cgo.GetRegleDB(db1)

	idRegle := cgo.GetIdRegle(RegleOrganisation, "PEC", "ABO")
	tournees := cgo.GetTourneesByRegle(db1, idRegle)

	
	societe, err := cgo.GetSociete(RegleOrganisation, "PEC", "ABO", "P1")

	if err != nil {
		fmt.Println(err)
	}

	intervenants, err := cgo.GetIntervenantBySosciete(db1, societe)

	if err != nil {
		fmt.Println(err)
	}

	for _, id_tournee := range tournees {
		//recuperer la liste des OTs
		liste_ots := lgo.GetLocalisationsByIdTournee(db1, id_tournee)
		//liste_ots := []lgo.Localisation{}

		//verifier l'activite des ots avec specialités des intervenants
		intervenants_bis := Verify_Activity(liste_ots, bxprtAdptr.Intervenants{intervenants})

		//recuperer les jobContributors par regle
		jobcontributors := bxprtAdptr.GetJobContributorsBySkill(intervenants_bis, jobContributors)

		fmt.Println("id tournee :", id_tournee, "nombre des jobContributors :", len(jobcontributors.JobContributorUserAgents))
		//bxprtAdptr.UpdateDumbContributorsBySkill(id_tournee, jobcontributors)
	}
	defer db1.Close()
}


func TestSet_Activite(t *testing.T) {

}

func TestSet_Disponibilte(t *testing.T) {

}

func TestCalcul_Distance(t *testing.T) {
	OT := bxprtAdptr.JobWO{}
	Tech := bxprtAdptr.JobContributorUserAgent{}
	d := Calcul_Distance(OT, Tech)
	fmt.Println(d)
}

func TestSort_Solution(t *testing.T) {

}

func TestGet_BestSolutionByScore(t *testing.T) {

}

func TestMake_OT_Tech(t *testing.T) {

}

func TestMake_sol(t *testing.T) {

}
*/