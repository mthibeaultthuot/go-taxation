package main

import (
	"context"
	"github.com/D3xt3rrrr/go-cloud/authentication-server/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GrpcServer struct {
	pb.UnimplementedJwtTokenServiceServer
}

func grpcListen() {
	log.Println("Auth Grpc start on http://localhost:9009")
	lis, err := net.Listen("tcp", ":9009")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterJwtTokenServiceServer(s, &GrpcServer{})

	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}

func (s GrpcServer) CheckToken(ctx context.Context, request *pb.JwtRequest) (*pb.JwtResponse, error) {
	jwtToken := getTokenFromBearer(request.Token)
	isValid, user := VerifyJwt(jwtToken)
	return &pb.JwtResponse{Username: user.Username, Permission: int32(user.Permission), IsJwtValid: isValid}, nil
}
