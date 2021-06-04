// baxperExporter_test.go
package baxpertExporter

import (
	bax "Baxpert_Geo_Optim/baxpertAdapter"
	"Outil_de_gestion_des_tournees/missionGeoOptim"
	"database/sql"

	"testing"

	"github.com/tealeg/xlsx-1.0.5"
)

func TestExporterMap(t *testing.T) {
	aJobContributorUserAgent := bax.JobContributorUserAgent{missionGeoOptim.JobContributorUserAgent{Nom: sql.NullString{"Nom_123", true}}}
	aJobClient := bax.JobClient{missionGeoOptim.JobClient{Nom_client: sql.NullString{"Nom_client_123", true}}}
	aJobReissue := bax.JobReissue{missionGeoOptim.JobReissue{ND: sql.NullString{"ND_123", true}}}
	aJobConstraint := bax.JobConstraint{missionGeoOptim.JobConstraint{Code_UI: sql.NullString{"code_UI_123", true}}}
	aJobInfo := bax.JobInfo{missionGeoOptim.JobInfo{Code_base: sql.NullString{"code_base_123", true}}}
	aJobWO := bax.JobWO{JobInfo: aJobInfo, JobConstraint: aJobConstraint, JobReissue: aJobReissue, JobClient: aJobClient}

	t1 := aJobInfo.GetMap()
	t.Logf("map %+v\n", t1)
	t2 := aJobWO.JobClient.GetMap()
	t.Logf("map %+v\n", t2)
	t3 := aJobContributorUserAgent.GetMap()
	t.Logf("map %+v\n", t3)
}

func TestExporterHeaderXLS(t *testing.T) {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("testHeader")
	//aJobContributorUserAgent := bax.JobContributorUserAgent{missionGeoOptim.JobContributorUserAgent{Nom: sql.NullString{"Nom_123", true}}}
	aJobClient := bax.JobClient{missionGeoOptim.JobClient{Nom_client: sql.NullString{"Nom_client_123", true}}}
	aJobReissue := bax.JobReissue{missionGeoOptim.JobReissue{ND: sql.NullString{"ND_123", true}}}
	aJobConstraint := bax.JobConstraint{missionGeoOptim.JobConstraint{Code_UI: sql.NullString{"code_UI_123", true}}}
	aJobInfo := bax.JobInfo{missionGeoOptim.JobInfo{Code_base: sql.NullString{"code_base_123", true}}}
	aJobWO := bax.JobWO{JobInfo: aJobInfo, JobConstraint: aJobConstraint, JobReissue: aJobReissue, JobClient: aJobClient}
	row := sheet.AddRow()
	addHeader(row, aJobWO)
	err = file.Save("./test_baxpertExporterHeader.xlsx")
	t.Logf("err in TestExporterHeaderXLS %+v\n", err)
}

//type OT_Tech map[JobWO]JobContributorUserAgent
func TestExporterXLS(t *testing.T) {
	aJobClient := bax.JobClient{missionGeoOptim.JobClient{Nom_client: sql.NullString{"Nom_client_123", true}}}
	aJobReissue := bax.JobReissue{missionGeoOptim.JobReissue{ND: sql.NullString{"ND_123", true}}}
	aJobConstraint := bax.JobConstraint{missionGeoOptim.JobConstraint{Code_UI: sql.NullString{"code_UI_123", true}}}
	aJobInfo := bax.JobInfo{missionGeoOptim.JobInfo{Code_base: sql.NullString{"code_base_123", true}}}
	aJobWO := bax.JobWO{JobInfo: aJobInfo, JobConstraint: aJobConstraint, JobReissue: aJobReissue, JobClient: aJobClient}
	aJobContributorUserAgent := bax.JobContributorUserAgent{missionGeoOptim.JobContributorUserAgent{Nom: sql.NullString{"Nom_123", true}}}
	OT_Tech_test := make(bax.OT_Tech, 0)
	OT_Tech_test[aJobWO] = aJobContributorUserAgent
	err := WriteExcelFile("./test_baxpertExporter.xlsx", OT_Tech_test)
	t.Logf("err in TestExporterXLS %+v\n", err)
}
