// baxpertAdatper.go
package baxpertAdapter

import (
	cgo "Outil_de_gestion_des_tournees/contrainteGeoOptim"
	ob "Outil_de_gestion_des_tournees/definitions"
	"Outil_de_gestion_des_tournees/missionGeoOptim"
	"database/sql"

	"errors"
	"fmt"
	"reflect"
	"time"
	//"fmt"
)

type OT_Tech map[JobWO]JobContributorUserAgent

type NumOT_NumTech map[string]int64

type Regle struct{
	Ui string
	Re string
	Depot string
}

type Lot map[Regle][]Solution

type Regle_Intervenants map[Regle]Intervenants

type Solution struct {
	OT    JobWO
	Tech  JobContributorUserAgent
	Score int64
}

type Solutions map[int64]Solution

type Disponibilites cgo.Disponibilite

type JobContributorUserAgent struct {
	missionGeoOptim.JobContributorUserAgent
}

type JobContributors struct {
	//JobContributorUserAgents []JobContributorUserAgent
	missionGeoOptim.JobContributors
	Dispos   map[int64]Disponibilites //map[Numero_utilisateur]Disponibilites
	Activite map[int64]string         //map[Numero_utilisateur]Activité
	Affecte    map[int64]bool           //map[Numero_utilisateur]libre ou pas
	//
}

type Intervenant struct {
	cgo.Intervenant
}

type Intervenants struct {
	Intervenants []cgo.Intervenant
}

type JobWOs struct {
	missionGeoOptim.JobWOs
}

type JobInfo struct {
	missionGeoOptim.JobInfo
}

type JobConstraint struct {
	missionGeoOptim.JobConstraint
}

type JobReissue struct {
	missionGeoOptim.JobReissue
}

type JobSafety struct {
	missionGeoOptim.JobSafety
}
type JobClient struct {
	missionGeoOptim.JobClient
}

type JobWO struct {
	//missionGeoOptim.JobWO
	JobInfo
	JobConstraint
	JobReissue
	JobSafety
	JobClient
	//SousTraitant //en appliquant RE+UI => SousTraitant
}

type ActiviteProduit struct {
	missionGeoOptim.ActiviteProduit
	/*Activite string `json:"activite"`
	Produit  string `json:"produit"`*/
}

type MapJob map[string]string

type ProduceMapJob interface {
	GetMap() MapJob
}

func PrettyPrint(at reflect.Type, a reflect.Value) (s string, err error) {
	switch at.Name() {
	case "NullString":
		s = ""
		if a.Interface().(sql.NullString).Valid {
			b := a.Interface().(sql.NullString).String
			s = fmt.Sprintf("%s", b)
		}
	case "NullTime":
		//b := sql.NullTime(a.Interface())
		s = ""
		if a.Interface().(sql.NullTime).Valid {
			b := a.Interface().(sql.NullTime).Time.Format(time.RFC3339Nano)
			s = fmt.Sprintf("%s", b)
		}
	case "NullBool":
		s = ""
		if a.Interface().(sql.NullBool).Valid {
			if a.Interface().(sql.NullBool).Bool {
				s = "Vrai"
			} else {
				s = "Faux"
			}
		}
	case "NullInt64":
		s = ""
		if a.Interface().(sql.NullInt64).Valid {
			s = fmt.Sprintf("%d", a.Interface().(sql.NullInt64).Int64)
		}
	default:
		err = errors.New("Format not supported")
	}
	return
}

func MergeMapJob(ms ...MapJob) map[string][]string {
	res := map[string][]string{}
	for _, m := range ms {
	srcMap:
		for k, v := range m {
			// Check if (k,v) was added before:
			for _, v2 := range res[k] {
				if v == v2 {
					continue srcMap
				}
			}
			res[k] = append(res[k], v)
		}
	}
	return res
}

func PrettyMap(val reflect.Value) (a MapJob) {
	a = make(MapJob)
	for i := 0; i < val.NumField(); i++ {
		if val.Field(i).Kind() == reflect.Struct {
			childVal := reflect.ValueOf(val.Field(i).Interface())
			for j := 0; j < childVal.NumField(); j++ {
				if childVal.IsValid() {
					z, e := PrettyPrint(childVal.Type().Field(j).Type, childVal.Field(j))
					if e == nil {
						a[childVal.Type().Field(j).Name] = z
					}
				}
			}
		} else {
			if val.IsValid() {
				z, e := PrettyPrint(val.Type().Field(i).Type, val.Field(i))
				if e == nil {
					a[val.Type().Field(i).Name] = z
				}
			}
		}
	}
	return
}

func (JW *JobInfo) GetMap() (a MapJob) {
	if JW != nil {
		val := reflect.ValueOf(JW).Elem()
		a = PrettyMap(val)
	}
	return
}

func (JW *JobConstraint) GetMap() (a MapJob) {
	if JW != nil {
		val := reflect.ValueOf(JW).Elem()
		a = PrettyMap(val)
	}
	return
}

func (JW *JobReissue) GetMap() (a MapJob) {
	if JW != nil {
		val := reflect.ValueOf(JW).Elem()
		a = PrettyMap(val)
	}
	return
}

func (JW *JobSafety) GetMap() (a MapJob) {
	if JW != nil {
		val := reflect.ValueOf(JW).Elem()
		a = PrettyMap(val)
	}
	return
}

func (JW *JobClient) GetMap() (a MapJob) {
	if JW != nil {
		val := reflect.ValueOf(JW).Elem()
		a = PrettyMap(val)
	}
	return
}

func (JW *JobContributorUserAgent) GetMap() (a MapJob) {
	if JW != nil {
		val := reflect.ValueOf(JW).Elem()
		a = PrettyMap(val)
	}
	return
}

func FindPath(idRegle int64) {
	var prmRes = ob.ParamResult{
		NumberOfRequetAPI: 0,
		Workers:           1,
		RayonClustring:    0.7,
		IndxOfEmplyee:     0,
		IndxOfAPIKey:      0,
		NumOfStopByDay:    2, //Ajouté
	}

	prmRes.Max_date_calcul_trajet = 5 //pour insérer un trajet s'il est daté de plus max_date_calcul_trajet jours
	prmRes.DateOfSchedualStr = time.Now().Add(time.Hour * 24).Format(ob.LayoutDate)
	prmRes.ConnString_parsTime = "geooptim2020admin:optim20geo@tcp(chronograf.fluidicite.xyz:3307)/geooptim2020demo?parseTime=true"
	var err error
	prmRes.DB, err = sql.Open("mysql", prmRes.ConnString_parsTime)
	if err == nil {
		//prmRes.DisposIntervenants = jobContributors.JobContributorUserAgents
		prmRes.DisposIntervenants = []cgo.Intervenant{
			{
				IdIntervenant:  int64(1),
				NomIntervenant: sql.NullString{String: "technicien1", Valid: true},
			},
			{
				IdIntervenant:  int64(2),
				NomIntervenant: sql.NullString{String: "technicien2", Valid: true},
			},
		}
		prmRes.NumberOfEmployee = len(prmRes.DisposIntervenants)
		ob.StartRoutingOptimization_WithoutSpinner(&prmRes, idRegle, "Bellegarde Clients Cuivre")
	}

	/* Modif: remplace DB par prmRes.DB
	dataBaseGeoOpt : ligne 246, 249, 257, 268.
	solver_maps : ligne 37, 43, 47, 87.
	solver : ligne 458.
	prendre en considération les modifs en definition : ligne 55 (//ajout JMB pour compilation 25042021)
	*/

}
