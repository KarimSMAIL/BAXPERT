module Example_GeoOptim

go 1.16

require (
	Baxpert_Geo_Optim/baxpert v0.0.0 // indirect
	Baxpert_Geo_Optim/baxpertAdapter v0.0.0 // indirect
	Baxpert_Geo_Optim/baxpertExporter v0.0.0 // indirect
	Baxpert_Geo_Optim/baxpertRules v0.0.0 // indirect
	Baxpert_Geo_Optim/baxpertAgenda v0.0.0 // indirect
	Outil_de_gestion_des_tournees/contrainteGeoOptim v0.0.0 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	gorm.io/driver/mysql v1.1.0
	gorm.io/gorm v1.21.9
)

replace (
	Baxpert_Geo_Optim/baxpert => ./baxpert
	Baxpert_Geo_Optim/baxpertAdapter => ./baxpertAdapter
	Baxpert_Geo_Optim/baxpertExporter => ./baxpertExporter
	Baxpert_Geo_Optim/baxpertRules => ./baxpertRules
	Baxpert_Geo_Optim/baxpertAgenda => ./baxpertAgenda
	Outil_de_gestion_des_tournees/contrainteGeoOptim => ../Outil_de_gestion_des_tournees/contrainteGeoOptim
	Outil_de_gestion_des_tournees/contrainteIntervention => ../Outil_de_gestion_des_tournees/contrainteIntervention
	Outil_de_gestion_des_tournees/definitions => ../Outil_de_gestion_des_tournees/definitions
	Outil_de_gestion_des_tournees/localisationGeoOptim => ../Outil_de_gestion_des_tournees/localisationGeoOptim
	Outil_de_gestion_des_tournees/missionGeoOptim => ../Outil_de_gestion_des_tournees/missionGeoOptim
	gogs.tetricite.com/yaroune/Outil_de_gestion_des_tournees => ../Outil_de_gestion_des_tournees
)
