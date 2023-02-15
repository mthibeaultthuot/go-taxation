package main

import (
	"context"
	"encoding/json"
	"github.com/D3xt3rrrr/verificationTax-service/data"
	"github.com/D3xt3rrrr/verificationTax-service/excels"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func (app *Config) Tax(r http.ResponseWriter, w *http.Request) {
	var body data.TaxBody

	err := json.NewDecoder(w.Body).Decode(&body)
	if err != nil {
		log.Panic(err)
	}

	tax := &data.Tax{
		false,
		false,
		body.Pst,
		body.Qst,
		"",
		"",
	}

	qstVerification(tax)
	pstVerification(tax)

	//file, err := CreateExcelFile(*tax)
	//if err != nil {
	//	return
	//}

	//r.Header().Set("Content-Disposition", "attachment; filename=search.xlsx")
	//r.Header().Set("Content-Type", w.Header.Get("Content-Type"))
	//r.Header().Set("Content-Length", w.Header.Get("Content-Length"))

	log.Println(tax)
	//io.Copy(r, file)
	//r.Write(&tax)
	r.Header().Set("Content-Type", "application/json")
	json.NewEncoder(r).Encode(tax)

	//connectGrpc(tax)
	//io.Copy(r, connectGrpc(tax))

	//_, err = r.Write(connectGrpc(tax))
	//if err != nil {
	//	return
	//}
	//xlsx.wri
}

func connectGrpc(tax *data.Tax) []byte {
	excelTax := &excels.Tax{
		IsPstValid: tax.IsPstValid,
		IsQstValid: tax.IsQstValid,
		PstNumber:  tax.PstNumber,
		QstNumber:  tax.QstNumber,
		Enterprise: tax.Enterprise,
		Date:       tax.Date,
	}

	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Panic(err)
	}

	defer conn.Close()

	c := excels.NewExcelServiceClient(conn)

	newTax := excels.TaxRequest{TaxEntry: excelTax}

	resp, err := c.Upload(context.Background(), &newTax)
	if err != nil {
		log.Panic(err)
	}

	return resp.GetResult()

}
