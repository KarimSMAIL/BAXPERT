// baxpertExporter
package baxpertExporter

import (
	"database/sql"

	bax "Baxpert_Geo_Optim/baxpertAdapter"

	"sort"

	"github.com/tealeg/xlsx"
)

func CheckErr(err error) {
}

func addHeader(row *xlsx.Row, aJobWO bax.JobWO) (headerSlice []string) {
	//mergeMapJob
	if row != nil {
		var styl *xlsx.Style
		styl = xlsx.NewStyle()
		styl.Font.Size = 14
		styl.Font.Name = "Calibri"
		styl.Alignment.Horizontal = "center"
		var cell *xlsx.Cell
		a := []bax.MapJob{}
		a = append(a, aJobWO.JobClient.GetMap())
		a = append(a, aJobWO.JobConstraint.GetMap())
		a = append(a, aJobWO.JobInfo.GetMap())
		a = append(a, aJobWO.JobReissue.GetMap())
		a = append(a, aJobWO.JobSafety.GetMap())
		pheader := bax.MergeMapJob(a...)
		headerSlice = make([]string, 0)
		for textH, _ := range pheader {
			headerSlice = append(headerSlice, textH)
		}
		// Now sort the slice
		sort.Strings(headerSlice)
		cell = row.AddCell()
		cell.Value = "Nom Intervenant"
		cell.SetStyle(styl)
		cell = row.AddCell()
		cell.Value = "Prénom Intervenant"
		cell.SetStyle(styl)
		// Iterate over all keys in a sorted order
		for _, key := range headerSlice {
			//fmt.Printf("Key: %d, Value: %s\n", key, pheader[key])
			//for textH, _ := range pheader {
			cell = row.AddCell()
			cell.Value = key
			cell.SetStyle(styl)
		}
	}
	return
}

func addRow(row *xlsx.Row, headerSlice []string, aJobWO bax.JobWO, aAgent bax.JobContributorUserAgent) (err error) {

	var cell *xlsx.Cell
	a := []bax.MapJob{}
	a = append(a, aJobWO.JobClient.GetMap())
	a = append(a, aJobWO.JobConstraint.GetMap())
	a = append(a, aJobWO.JobInfo.GetMap())
	a = append(a, aJobWO.JobReissue.GetMap())
	a = append(a, aJobWO.JobSafety.GetMap())
	pheader := bax.MergeMapJob(a...)
	styl := xlsx.NewStyle()
	styl.Font.Size = 11
	styl.Font.Name = "Calibri"
	styl.Alignment.Horizontal = "left"

	b := aAgent.GetMap()

	cell = row.AddCell()
	cell.Value = b["Nom"] //"Nom Intervenant"
	cell.SetStyle(styl)
	cell = row.AddCell()
	cell.Value = b["Prenom"] //"Prénom Intervenant"
	cell.SetStyle(styl)
	for _, key := range headerSlice {
		//for textH, _ := range pheader {
		cell = row.AddCell()
		cell.Value = pheader[key][0]
		cell.SetStyle(styl)
	}
	return
}

func WriteExcelFile(pathOfsave string, aJobWOMap bax.OT_Tech) (err error) {
	maxNbOfDates := 1
	var ot bax.JobWO

	var file *xlsx.File
	var sheet *xlsx.Sheet

	//	var cell *xlsx.Cell
	var styl *xlsx.Style

	styl = xlsx.NewStyle()
	styl.Font.Size = 11
	styl.Font.Name = "Calibri"
	styl.Alignment.Horizontal = "center"

	file = xlsx.NewFile()
	//Onglet1
	sheet, err = file.AddSheet("Affectations")
	CheckErr(err)
	row := sheet.AddRow()
	//row.AddCell()
	ot.Activite = sql.NullString{"", false}
	//way to get an item of the Map
	var k bax.JobWO
	for k, _ = range aJobWOMap {
		break
	}

	headerSlice := addHeader(row, k)

	for solWO, AgentWO := range aJobWOMap {

		row = sheet.AddRow()
		addRow(row, headerSlice, solWO, AgentWO)

	}

	sheet.Cols[0].Width = 20
	sheet.Cols[1].Width = 20
	sheet.Cols[2].Width = 20
	sheet.Cols[3].Width = 15
	sheet.Cols[4].Width = 15
	sheet.Cols[5].Width = 20
	sheet.Cols[6].Width = 20
	sheet.Cols[7].Width = float64(maxNbOfDates * 15)

	err = file.Save(pathOfsave)
	CheckErr(err)
	return
}
