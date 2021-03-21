package grpcclient

import (
	"IvanFerdino/bibit-golang-test/commons"
	v1 "IvanFerdino/bibit-golang-test/pkg/grpc/v1"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"os"
	"time"
)

type ClientConnection struct {
	Conn   *grpc.ClientConn
	ctx    context.Context
	movieServiceCl v1.MovieServiceClient
	//client 2 etc...
	cancel context.CancelFunc
}

func New(grpchost string, grpcport string) *ClientConnection {
	conn, err := grpc.Dial(fmt.Sprintf("%v:%v", grpchost, grpcport), grpc.WithInsecure())
	if err != nil {
		commons.LogError(fmt.Sprintf("Failed to connect: %v",err.Error()))
		os.Exit(1)
	}
	commons.LogInfo(fmt.Sprintf("connected to grpc server %v:%v\n",grpchost,grpcport))
	movieServiceCl := v1.NewMovieServiceClient(conn)
	//cient2 etc...

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	cl := ClientConnection{
		ctx:    ctx,
		Conn:   conn,
		movieServiceCl: movieServiceCl,
		cancel: cancel,
	}

	return &cl
}

func Search(cl *ClientConnection,param *v1.MovieSearchRequest)(*v1.MovieSearchResponse,error) {
	response, err := cl.movieServiceCl.Search(cl.ctx, param)
	cl.cancel() //close cl
	return response, err
}

func Detail(cl *ClientConnection,param *v1.MovieDetailRequest)(*v1.MovieDetailResponse,error) {
	response, err := cl.movieServiceCl.Detail(cl.ctx, param)
	cl.cancel()//close cl
	return response, err
}
//etc, implement others
