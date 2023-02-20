package main

import (
	"context"
	"github.com/D3xt3rrrr/go-cloud/authentication-server/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GrpcServer struct {
	pb.UnimplementedExcelServiceServer
	data []byte
}

func (grpc *GrpcServer) CheckToken(ctx context.Context, request *pb.JwtRequest) (*pb.JwtResponse, error) {
	jwtResp := &pb.JwtResponse{
		IsValid:       false,
		Authorization: "",
	}

	return jwtResp, nil
}

func grpcListen() {
	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatal("GRPC ==> Failed to listen to port :9001 %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterExcelServiceServer(s, &GrpcServer{data: []byte{}})

	err = s.Serve(lis)
	if err != nil {
		log.Panic(err)
	}
}
