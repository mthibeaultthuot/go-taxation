package main

import (
	"context"
	data2 "github.com/D3xt3rrrr/excel-service/data"
	"github.com/D3xt3rrrr/excel-service/excels"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net"
)

type ExcelServer struct {
	excels.UnimplementedExcelServiceServer
	data []byte
}

func (e ExcelServer) Upload(ctx context.Context, req *excels.TaxRequest) (*excels.TaxResponse, error) {
	taxEntry := req.GetTaxEntry()

	tax := data2.Tax{
		IsPstValid: taxEntry.IsPstValid,
		IsQstValid: taxEntry.IsQstValid,
		PstNumber:  taxEntry.PstNumber,
		QstNumber:  taxEntry.QstNumber,
		Enterprise: taxEntry.Enterprise,
		Date:       taxEntry.Date,
	}

	newfile, err := CreateExcelFile(tax)
	if err != nil {
		log.Panic(err)
	}

	data, err := ioutil.ReadFile(newfile.Name())

	return &excels.TaxResponse{
		Result: data,
	}, nil
}

func grpcListen() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal("Failed to listen to port :9000 %v", err)
	}

	s := grpc.NewServer()
	excels.RegisterExcelServiceServer(s, &ExcelServer{data: []byte{}})

	err = s.Serve(lis)
	if err != nil {
		log.Panic(err)
	}
}
