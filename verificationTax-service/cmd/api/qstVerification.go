package main

import (
	"encoding/json"
	"github.com/D3xt3rrrr/verificationTax-service/data"
	"io"
	"log"
	"net/http"
)

var (
	qstToken    = "https://www.businessregistration-inscriptionentreprise.gc.ca/ebci/brom/registry/pub/reg_01_Ld.action"
	qstRequest  = "https://www.businessregistration-inscriptionentreprise.gc.ca/ebci/brom/registry/pub/reg_01_Sbmt.action"
	qstResponse = "https://www.businessregistration-inscriptionentreprise.gc.ca/ebci/brom/registry/pub/reg_02_Ld.action"
)

func qstVerification(tax *data.Tax) jsonResponse {
	res, err := http.Get("https://svcnab2b.revenuquebec.ca/2019/02/ValidationTVQ/" + tax.QstNumber)
	if err != nil {
		log.Panic(err)
	}
	response := res.Body
	bodyBytes, err := io.ReadAll(response)

	data := jsonResponse{}
	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		log.Panic(err)
	}

	tax.Enterprise = data.Resultat.NomEntreprise
	tax.IsPstValid = data.OperationReussie

	return data
}
