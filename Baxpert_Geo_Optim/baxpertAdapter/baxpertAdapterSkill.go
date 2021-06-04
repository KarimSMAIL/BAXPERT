package baxpertAdapter

import (
	"Outil_de_gestion_des_tournees/contrainteGeoOptim"
	"database/sql"
	"encoding/json"

	"fmt"
	//	"time"
)

var p = fmt.Println

type specialite map[string]int

type JobContributorSkill struct {
	Numero        int64          `json:"numero"` //numero_utilisateur
	SType         string         `json:"sType "` // ST, Societe, Eiffage
	SocieteAsso   string         `json:"societeAsso"`
	Emploi        string         `json:"emploi"`
	Typo          string         `json:"typo"`
	EffectifTotal int            `json:"effectifTotal"`
	Produit       map[string]int `json:"produit"`
	Specialite    map[string]int `json:"specialite"` //effectif par specialite
}

func NewJobContributorSkill() *JobContributorSkill {
	return &JobContributorSkill{}
}

func GetJobContributorSkill(qualiteIntervenantJson sql.NullString) JobContributorSkill {
	Skills := JobContributorSkill{}
	_ = json.Unmarshal([]byte(qualiteIntervenantJson.String), &Skills)
	return Skills
}

//recuperer la liste des techniciens par regle (par OT)
func GetJobContributorsBySkill(intervenants Intervenants, jobContributors JobContributors) JobContributors {
	jobcontributors := JobContributors{}
	for _, inter := range intervenants.Intervenants {
		//fmt.Println("Numero intervenant")
		qI := GetJobContributorSkill(inter.QualiteIntervenantJson)
		//fmt.Println("Numero intervenant", qI.Numero)
		for _, tech := range jobContributors.JobContributorUserAgents {
			if qI.Numero == tech.Numero_utilisateur {
				jobcontributors.JobContributorUserAgents = append(jobcontributors.JobContributorUserAgents, tech)
			}
		}
	}
	return jobcontributors
}

//fait la mise à jour avec une condition sur les techniciens en utilisant la table geo_regle_intervenant
// regarde : GetListIntervenantOverbookTournee
func UpdateDumbContributorsBySkill(idTournee int64, jobcontributors JobContributors, priorite_regle int64) error {
	//GetList
	//db, err := sql.Open("mysql", "snipeuser:snipeprod998A;@tcp(chronograf.fluidicite.xyz:3307)/snipe_geo?parseTime=true")
	//db, err := sql.Open("mysql", "geooptim2021admin:geo12;optim@tcp(chronograf.fluidicite.xyz:3307)/geooptim2021preprod?parseTime=true")
	db, err := sql.Open("mysql", "geooptim2020admin:optim20geo@tcp(chronograf.fluidicite.xyz:3307)/geooptim2020demo?charset=utf8mb4&parseTime=True&loc=Local")
	defer db.Close()
	for _, tech := range jobcontributors.JobContributorUserAgents {
		var intervTournee = contrainteGeoOptim.IntervenantTournee{}
		err := intervTournee.GetDumbIntervenantTournee(db, idTournee)
		if err != nil {
			continue
		}
		intervTournee.PrioriteRegle = sql.NullInt64{Int64: priorite_regle, Valid: true}
		intervTournee.UpdateIntervenantTournee(db, tech.Numero_utilisateur)
	}
	return err
}

func PutIntervenant(priorite_regle int64, id_regle int64, id_tournee int64, old_numtech int64, new_numtech int64) (err error) {
	var intervTournee = contrainteGeoOptim.IntervenantTournee{
		PrioriteRegle: sql.NullInt64{Int64: priorite_regle, Valid: true},
		IdRegle:       sql.NullInt64{Int64: id_regle, Valid: true},
		IdTournee:     sql.NullInt64{Int64: id_tournee, Valid: true},
		IdIntervenant: sql.NullInt64{Int64: old_numtech, Valid: true},
		//DatePlanification: sql.NullTime{Time: t1, Valid: true},
	}
	//db, err := sql.Open("mysql", "snipeuser:snipeprod998A;@tcp(chronograf.fluidicite.xyz:3307)/snipe_geo?parseTime=true")
	//db, err := sql.Open("mysql", "geooptim2021admin:geo12;optim@tcp(chronograf.fluidicite.xyz:3307)/geooptim2021preprod?parseTime=true")
	db, err := sql.Open("mysql", "geooptim2020admin:optim20geo@tcp(chronograf.fluidicite.xyz:3307)/geooptim2020demo?charset=utf8mb4&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		return
	}
	err = intervTournee.UpdateIntervenantTournee(db, new_numtech)
	return
}

func GetStaticIntervenant() []Intervenant {
	intervenants := make([]Intervenant, 2)
	intervenants[0].IdIntervenant = int64(1)
	intervenants[0].IdLocalisation = sql.NullInt64{Int64: 1, Valid: true}
	intervenants[0].NomIntervenant = sql.NullString{String: " ", Valid: true}
	intervenants[0].QualiteIntervenantJson = sql.NullString{String: `{"n":2933,"t":"Societe","soc":"01COM","e":"","typo":"","cnt":15,"spe":{"PLI":10,"S02":2}}`, Valid: true}

	intervenants[1].IdIntervenant = int64(2)
	intervenants[1].IdLocalisation = sql.NullInt64{Int64: 2, Valid: true}
	intervenants[1].NomIntervenant = sql.NullString{String: " ", Valid: true}
	intervenants[1].QualiteIntervenantJson = sql.NullString{String: `{"n":1996,"t":"ST","soc":"01COM","e":"TECH","typo":"SAV/PROD","cnt":0,"spe":null}`, Valid: true}

	return intervenants
}

func GetActiviteProduit(TypeJson string) ActiviteProduit {
	Activiteproduit := ActiviteProduit{}
	_ = json.Unmarshal([]byte(TypeJson), &Activiteproduit)
	return Activiteproduit

}
