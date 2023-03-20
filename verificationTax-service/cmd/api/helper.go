package main

type jsonResponse struct {
	Resultat struct {
		StatutSousDossierUsager string      `json:"StatutSousDossierUsager"`
		DescriptionStatut       string      `json:"DescriptionStatut"`
		DateStatut              string      `json:"DateStatut"`
		NomEntreprise           string      `json:"NomEntreprise"`
		RaisonSociale           interface{} `json:"RaisonSociale"`
	} `json:"Resultat"`
	OperationReussie     bool          `json:"OperationReussie"`
	MessagesFonctionnels []interface{} `json:"MessagesFonctionnels"`
}
