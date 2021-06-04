// rulesPathFinding
//ici gestion des paths - utilisation de FindPath()
package baxpertRules

import (
	//bxprAdtr "Baxpert_Geo_Optim/baxpertAdapter"
	cgo "Outil_de_gestion_des_tournees/contrainteGeoOptim"
	ob "Outil_de_gestion_des_tournees/definitions"
	"database/sql"

	//"fmt"
	"time"
)

type RulePathFinding struct {
	IdRegle int64 //KS
	Name    string
	Score   int
}

func NewRulePathFinding() *RulePathFinding {
	return &RulePathFinding{
		IdRegle: 100,
		Name:    "the fastest route between the technician and the intervention site",
		Score:   20,
	}
}

// assign ruleDepot to the constraint and add its score to the constraint score
func (rulePathFind RulePathFinding) intoContrainte(c *Constraint) {
	c.ConstraintList = append(c.ConstraintList, rulePathFind)
	//fmt.Println("the rulePathFinding is assigned to the constraint ", c.Title)
}

//this function calculates a trip for each technician given in input, it is necessary for calculate Trajet ???
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

type RulePathFindingAdapter struct {
	rulePathFind   *RulePathFinding
	function_title string
}

func NewRulePathFindingAdapter(rulePathFind *RulePathFinding, function_title string) *RulePathFindingAdapter {
	return &RulePathFindingAdapter{
		rulePathFind:   rulePathFind,
		function_title: function_title,
	}
}

func (r *RulePathFinding) AssignRuleIntoConstraint(c *Constraint) {
	//fmt.Println("Rule adapter converts  rule1 to rulePathFinding.")
	r.intoContrainte(c)
}

func (r *RulePathFinding) ProcessConstraint(c *Constraint) {

	//according to the state performs a calculation
	c.GlobalScore += r.Score //c.globalScore++
	//fmt.Printf("rulePathFinding processConstraint, globalScore %v \n", c.GlobalScore)

}

/*func (rPF *rulePathFindingAdapter) InvokeFindPath() {
	bxprAdtr.FindPath(rPF.rulePathFind.IdRegle)
	fmt.Println(rPF.function_title, " has done its calculation, the shortest path is calculate => the rule is respected")
}*/

/*func ProcessConstraint(c *Constraint, r rulePathFinding) {

	//according to the state performs a calculation
	c.globalScore += r.score //c.globalScore++
	fmt.Printf("rulePathFinding processConstraint, globalScore %v \n", c.globalScore)

}


*/
