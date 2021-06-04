package baxpertAdapter

import (
	"database/sql"
	//"fmt"
)

type JobContributorCompany struct {
	//ici les informations dérivées de
	SocieteAsso sql.NullString
	EffectifTotal sql.NullInt64
	Specialite specialite //avec Effectif MAP ? Array ? (sur JobContributorSkill)
}

func NewJobContributorCompany() *JobContributorCompany{
	return &JobContributorCompany{}
}