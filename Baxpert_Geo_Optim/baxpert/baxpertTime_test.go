// baxpertTime_test
package baxpert

import (
	"testing"

	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestBuildRegleContrainteOTJour(t *testing.T) {
	//on connait Activite / Produit et donc on peut récupérer le temps depuis la table
	dsn := "geooptim2020admin:optim20geo@tcp(chronograf.fluidicite.xyz:3307)/geooptim2020demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	jobWos := GetJobWOsAgence(db, "LYON Clients FTTH")

	for _, wo := range jobWos.Ots {
		BuildRegleContrainteOTJour(wo)
	}
	fmt.Printf("map %+v\n", ActiviteNbrOT)

}
