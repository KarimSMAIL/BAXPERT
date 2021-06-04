module Baxpert_Geo_Optim/baxpertExporter

go 1.16

require (
github.com/tealeg/xlsx v1.0.5
Baxpert_Geo_Optim/baxpertAdapter v0.0.0
Outil_de_gestion_des_tournees/contrainteGeoOptim v0.0.0 // indirect
Outil_de_gestion_des_tournees/missionGeoOptim v0.0.0 // indirect
Outil_de_gestion_des_tournees/contrainteIntervention v0.0.0 // indirect
Outil_de_gestion_des_tournees/localisationGeoOptim v0.0.0 // indirect
Outil_de_gestion_des_tournees/definitions v0.0.0 // indirect
)

replace (
Outil_de_gestion_des_tournees/contrainteGeoOptim => ../../Outil_de_gestion_des_tournees/contrainteGeoOptim
Outil_de_gestion_des_tournees/missionGeoOptim => ../../Outil_de_gestion_des_tournees/missionGeoOptim
Baxpert_Geo_Optim/baxpertAdapter => ../baxpertAdapter
Outil_de_gestion_des_tournees/contrainteIntervention => ../../Outil_de_gestion_des_tournees/contrainteIntervention
Outil_de_gestion_des_tournees/localisationGeoOptim => ../../Outil_de_gestion_des_tournees/localisationGeoOptim
Outil_de_gestion_des_tournees/definitions => ../../Outil_de_gestion_des_tournees/definitions
)