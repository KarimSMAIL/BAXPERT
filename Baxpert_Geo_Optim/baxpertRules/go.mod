module Baxpert_Geo_Optim/baxpertRules

go 1.16

require (
	Outil_de_gestion_des_tournees/contrainteGeoOptim v0.0.0
	Outil_de_gestion_des_tournees/contrainteIntervention v0.0.0 // indirect
	Outil_de_gestion_des_tournees/definitions v0.0.0
	Outil_de_gestion_des_tournees/localisationGeoOptim v0.0.0 // indirect
	Outil_de_gestion_des_tournees/missionGeoOptim v0.0.0
	github.com/mitchellh/mapstructure v1.4.1 // indirect
	gorm.io/gorm v1.21.9
)

replace (
	Outil_de_gestion_des_tournees/contrainteGeoOptim => ../../Outil_de_gestion_des_tournees/contrainteGeoOptim
	Outil_de_gestion_des_tournees/contrainteIntervention => ../../Outil_de_gestion_des_tournees/contrainteIntervention
	Outil_de_gestion_des_tournees/definitions => ../../Outil_de_gestion_des_tournees/definitions
	Outil_de_gestion_des_tournees/localisationGeoOptim => ../../Outil_de_gestion_des_tournees/localisationGeoOptim
	Outil_de_gestion_des_tournees/missionGeoOptim => ../../Outil_de_gestion_des_tournees/missionGeoOptim
)
