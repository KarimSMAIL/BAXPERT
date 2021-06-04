package baxpertAdapter

import (
	"Outil_de_gestion_des_tournees/missionGeoOptim"
	"database/sql"
	"fmt"
	"testing"
)

func TestFindPath(t *testing.T) {
	/*idRegle := int64(100)
	FindPath(idRegle)*/
}

func TestJobInfoGetMap(t *testing.T) {
	a := &JobInfo{missionGeoOptim.JobInfo{Code_base: sql.NullString{"code_base_123", true}}}
	t1 := a.GetMap()
	fmt.Printf("map %+v\n", t1)
}

func TestJobConstraintGetMap(t *testing.T) {
	a := &JobConstraint{missionGeoOptim.JobConstraint{Code_UI: sql.NullString{"code_UI_123", true}}}
	t1 := a.GetMap()
	fmt.Printf("map %+v\n", t1)
}

func TestJobReissueGetMap(t *testing.T) {
	a := &JobReissue{missionGeoOptim.JobReissue{ND: sql.NullString{"ND_123", true}}}
	t1 := a.GetMap()
	fmt.Printf("map %+v\n", t1)
}

func TestJobClientGetMap(t *testing.T) {
	a := &JobClient{missionGeoOptim.JobClient{Nom_client: sql.NullString{"Nom_client_123", true}}}
	t1 := a.GetMap()
	fmt.Printf("map %+v\n", t1)
}

func TestJobContributorUserAgentGetMap(t *testing.T) {
	a := &JobContributorUserAgent{missionGeoOptim.JobContributorUserAgent{Nom: sql.NullString{"Nom_123", true}}}
	t1 := a.GetMap()
	fmt.Printf("map %+v\n", t1)
}
