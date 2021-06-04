package baxpertAgenda

import (
	//"time"
	bxprtAdptr "Baxpert_Geo_Optim/baxpertAdapter"
	cgo "Outil_de_gestion_des_tournees/contrainteGeoOptim"
	lgo "Outil_de_gestion_des_tournees/localisationGeoOptim"
	"database/sql"
	"fmt"
	"testing"
)

func TestAverage(t *testing.T) {
	//db, _ := sql.Open("mysql", "geooptim2021admin:geo12;optim@tcp(chronograf.fluidicite.xyz:3307)/geooptim2021preprod?parseTime=true")
	db, _ := sql.Open("mysql", "geooptim2020admin:optim20geo@tcp(chronograf.fluidicite.xyz:3307)/geooptim2020demo?parseTime=true")
	agenda := Agenda{}
	agenda.InitAgenda(db)

	agenda.Average()
	fmt.Println("la moyenne est :", agenda.Moyenne)
}

func TestUpdateLastDate(t *testing.T) {
	//db, _ := sql.Open("mysql", "geooptim2021admin:geo12;optim@tcp(chronograf.fluidicite.xyz:3307)/geooptim2021preprod?parseTime=true")
	db, _ := sql.Open("mysql", "geooptim2020admin:optim20geo@tcp(chronograf.fluidicite.xyz:3307)/geooptim2020demo?parseTime=true")
	agenda := Agenda{}
	agenda.InitAgenda(db)

	agenda.Planing(db)

	agenda.UpdateLastDate(db)
}

func TestPlaning(t *testing.T) {
	//db, _ := sql.Open("mysql", "geooptim2021admin:geo12;optim@tcp(chronograf.fluidicite.xyz:3307)/geooptim2021preprod?parseTime=true")
	db, _ := sql.Open("mysql", "geooptim2020admin:optim20geo@tcp(chronograf.fluidicite.xyz:3307)/geooptim2020demo?parseTime=true")
	defer db.Close()
	agenda := Agenda{}
	agenda.InitAgenda(db)

	agenda.Planing(db)
	fmt.Println("len Result: ", len(agenda.ResultListIntervenantTournee))
	// for id, value := range agenda.ResultListIntervenantTournee {
	// 	fmt.Println("nombre de tournee en Result pour id intervenants : ", id, ", ", len(value))
	// }
}

func TestInitAgenda(t *testing.T) {
	//db, _ := sql.Open("mysql", "geooptim2021admin:geo12;optim@tcp(chronograf.fluidicite.xyz:3307)/geooptim2021preprod?parseTime=true")
	db, _ := sql.Open("mysql", "geooptim2020admin:optim20geo@tcp(chronograf.fluidicite.xyz:3307)/geooptim2020demo?parseTime=true")
	defer db.Close()
	agenda := Agenda{}
	agenda.InitAgenda(db)
	t.Log("nombre des intervenants à equilibrer", len(agenda.ListIntervenantAgenda))
	t.Log("nombre des tournee à equilibrer", len(agenda.ListTourneeByIntervenantAll))

}

func TestExist(t *testing.T) {
	db, _ := sql.Open("mysql", "geooptim2021admin:geo12;optim@tcp(chronograf.fluidicite.xyz:3307)/geooptim2021preprod?parseTime=true")
	//db, _ := sql.Open("mysql", "geooptim2020admin:optim20geo@tcp(chronograf.fluidicite.xyz:3307)/geooptim2020demo?parseTime=true")
	defer db.Close()
	var tech = cgo.Intervenant{}
	err := tech.GetIntervenant(db, 1988)
	if err != nil {
		fmt.Println(err)
	}
	rgl := cgo.RegleIntervenant{
		IdRegle:       sql.NullInt64{Int64: 3, Valid: true},
		PrioriteRegle: sql.NullInt64{Int64: 2, Valid: true},
		IdIntervenant: sql.NullInt64{Int64: 1988, Valid: true},
	}
	intervenants, err := cgo.GetIntervenantByRegle(db, rgl)

	fmt.Println("Exist of IdIntervenant : 1988 is :", Exist(1988, bxprtAdptr.Intervenants{intervenants}))
}

func TestVerifyActivityProduit(t *testing.T) {
	db, _ := sql.Open("mysql", "geooptim2021admin:geo12;optim@tcp(chronograf.fluidicite.xyz:3307)/geooptim2021preprod?parseTime=true")
	liste_ots := lgo.GetLocalisationsByIdTournee(db, 123)
	var tech = cgo.Intervenant{}
	err := tech.GetIntervenant(db, 2130)

	if err != nil {
		fmt.Println(err)
	}
	if VerifyActivityProduit(liste_ots, bxprtAdptr.Intervenant{tech}) {
		fmt.Println("c'est verifié, Activité et Produit ok")
	} else {
		fmt.Println("Activité et Produit non ok")
	}
}

/*func TestPutRegleIntervenantTournee(t *testing.T) {
	//db, _ := sql.Open("mysql", "geooptim2021admin:geo12;optim@tcp(chronograf.fluidicite.xyz:3307)/geooptim2021preprod?parseTime=true")
	db, _ := sql.Open("mysql", "geooptim2020admin:optim20geo@tcp(chronograf.fluidicite.xyz:3307)/geooptim2020demo?parseTime=true")
	agenda := Agenda{}
	agenda.InitAgenda(db)

	agenda.PutRegleIntervenantTournee(db)
}

*/
