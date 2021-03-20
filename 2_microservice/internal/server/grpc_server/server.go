package grpc_server

import (
	"IvanFerdino/bibit-golang-test/commons"
	"IvanFerdino/bibit-golang-test/pkg/grpc/v1"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
)

const (
	banner = `
	 ██████╗ ██████╗ ██████╗  ██████╗    ███████╗███████╗██████╗ ██╗   ██╗███████╗██████╗ 
	██╔════╝ ██╔══██╗██╔══██╗██╔════╝    ██╔════╝██╔════╝██╔══██╗██║   ██║██╔════╝██╔══██╗
	██║  ███╗██████╔╝██████╔╝██║         ███████╗█████╗  ██████╔╝██║   ██║█████╗  ██████╔╝
	██║   ██║██╔══██╗██╔═══╝ ██║         ╚════██║██╔══╝  ██╔══██╗╚██╗ ██╔╝██╔══╝  ██╔══██╗
	╚██████╔╝██║  ██║██║     ╚██████╗    ███████║███████╗██║  ██║ ╚████╔╝ ███████╗██║  ██║
	 ╚═════╝ ╚═╝  ╚═╝╚═╝      ╚═════╝    ╚══════╝╚══════╝╚═╝  ╚═╝  ╚═══╝  ╚══════╝╚═╝  ╚═╝
	_____________________________________________
`
)

type Server struct {
	movieServiceServer v1.MovieServiceServer
	port string
}

func New(movieServiceServer v1.MovieServiceServer,port string) *Server{
	commons.LogInfo(fmt.Sprintf("Initialize grpc server...,port %v",port))
	return &Server{
		movieServiceServer:   movieServiceServer,
		port: port,
	}
}

func (s *Server) Run(ctx context.Context) error {
	listen, err := net.Listen("tcp", ":"+s.port)
	if err != nil {
		return err
	}

	srv:=grpc.NewServer()
	v1.RegisterMovieServiceServer(srv,s.movieServiceServer)
	//initialize other service server if needed...

	c:=make(chan os.Signal,1)
	signal.Notify(c,os.Interrupt)
	go func() {
		for range c {
			commons.LogInfo(fmt.Sprintf("Shutting down GRPC server..."))
			srv.GracefulStop()
			<-ctx.Done()
		}
	}()

	commons.LogInfo(fmt.Sprintf("Initialization done ..."))
	commons.LogInfo(fmt.Sprintf("%v",banner))
	commons.LogInfo(fmt.Sprintf("GRPC Server up and listening port %v",s.port))
	return srv.Serve(listen)
}