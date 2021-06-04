module Baxpert_Geo_Optim/baxpert

go 1.16

require (
	Baxpert_Geo_Optim/baxpertAdapter v0.0.0
	Baxpert_Geo_Optim/baxpertRules v0.0.0
	Outil_de_gestion_des_tournees/contrainteGeoOptim v0.0.0
	Outil_de_gestion_des_tournees/localisationGeoOptim v0.0.0 // indirect
	Outil_de_gestion_des_tournees/missionGeoOptim v0.0.0
	gorm.io/driver/mysql v1.1.0 // indirect
	gorm.io/gorm v1.21.9
)

replace (
	Baxpert_Geo_Optim/baxpertAdapter => ../baxpertAdapter
	Baxpert_Geo_Optim/baxpertRules => ../baxpertRules
	Outil_de_gestion_des_tournees/contrainteGeoOptim => ../../Outil_de_gestion_des_tournees/contrainteGeoOptim
	Outil_de_gestion_des_tournees/contrainteIntervention => ../../Outil_de_gestion_des_tournees/contrainteIntervention
	Outil_de_gestion_des_tournees/definitions => ../../Outil_de_gestion_des_tournees/definitions
	Outil_de_gestion_des_tournees/localisationGeoOptim => ../../Outil_de_gestion_des_tournees/localisationGeoOptim
	Outil_de_gestion_des_tournees/missionGeoOptim => ../../Outil_de_gestion_des_tournees/missionGeoOptim
)
