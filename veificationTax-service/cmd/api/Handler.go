package main

import (
	"context"
	"encoding/json"
	"github.com/D3xt3rrrr/verificationTax-service/data"
	"github.com/D3xt3rrrr/verificationTax-service/pb"
	_ "google.golang.org/genproto/googleapis/container/v1beta1"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func (app *Config) Tax(r http.ResponseWriter, w *http.Request) {
	bearerToken := w.Header.Get("Authorization")
	user := AuthGrpc(bearerToken)
	if !user.IsJwtValid {
		http.Redirect(r, w, "http://localhost:8080", 403)
		return
	}

	var body data.TaxBody

	err := json.NewDecoder(w.Body).Decode(&body)
	if err != nil {
		log.Panic(err)
	}

	tax := &data.Tax{
		Username:  user.Username,
		PstNumber: body.Pst,
		QstNumber: body.Qst,
	}

	qstVerification(tax)
	pstVerification(tax)

	log.Println("userna,e", user)
	err = app.Repo.Insert(tax)
	if err != nil {
		return
	}

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
	err = json.NewEncoder(r).Encode(tax)
	if err != nil {
		return
	}

	//connectGrpc(tax)
	//io.Copy(r, connectGrpc(tax))

	//_, err = r.Write(connectGrpc(tax))
	//if err != nil {
	//	return
	//}
	//xlsx.wri
}

func (app *Config) GetAllFromUser(r http.ResponseWriter, w *http.Request) {
	bearerToken := w.Header.Get("Authorization")
	user := AuthGrpc(bearerToken)
	if !user.IsJwtValid {
		http.Redirect(r, w, "http://localhost:8080", 403)
		return
	}
	taxList := app.Repo.GetAllFromUser(user.Username)
	err := json.NewEncoder(r).Encode(taxList)
	if err != nil {
		panic(err)
	}
}

func (app *Config) GetAll(r http.ResponseWriter, w *http.Request) []data.Tax {
	return nil
}

/*
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

}*/

func AuthGrpc(bearerToken string) data.User {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9009", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	c := pb.NewJwtTokenServiceClient(conn)

	resp, err := c.CheckToken(context.Background(), &pb.JwtRequest{Token: bearerToken})
	if err != nil {
		panic(err)
	}

	return data.User{
		Username:        resp.Username,
		IsJwtValid:      resp.IsJwtValid,
		PermissionLevel: resp.Permission,
	}

}
