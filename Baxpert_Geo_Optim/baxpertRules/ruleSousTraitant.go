package baxpertRules

import (
	bxprAdtr "Baxpert_Geo_Optim/baxpertAdapter"
	//"Outil_de_gestion_des_tournees/missionGeoOptim"
	//	"database/sql"
	//"fmt"
	//"time"
	//"github.com/mitchellh/mapstructure"
	//	"gorm.io/gorm"
)

type RuleSousTraintant struct {
	IdRegle int64 //KS
	Name    string
	Techs   bxprAdtr.JobContributors
	Score   int
}

//New Rule
func NewRuleSousTraintant(IdRegle int64, Techs bxprAdtr.JobContributors) *RuleSousTraintant {
	return &RuleSousTraintant{
		IdRegle: IdRegle,
		Name:    "sous Traintant",
		Techs:   Techs,
		Score:   10,
	}
}

// func GetSousTraitant(db *gorm.DB, Code_UI sql.NullString, Code_RE sql.NullString) bxprAdtr.SousTraitant {

// 	return bxprAdtr.SousTraitant{}
// }

// func GetListTechnicians(db *gorm.DB, St bxprAdtr.SousTraitant){

// 	St.Techs = bxprAdtr.JobContributors{} //Requet sql
// }

func (r *RuleSousTraintant) ProcessConstraint(c *Constraint) {

}
