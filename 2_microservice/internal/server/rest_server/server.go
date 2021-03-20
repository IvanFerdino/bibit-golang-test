package rest_server

import (
	"IvanFerdino/bibit-golang-test/commons"
	"IvanFerdino/bibit-golang-test/internal/controller/echo_controller"
	"IvanFerdino/bibit-golang-test/internal/service/movie_service"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tylerb/graceful"
	"time"
)

const (
	banner = `
	   ____    __
	  / __/___/ /  ___
	 / _// __/ _ \/ _ \
	/___/\__/_//_/\___/
	High performance, minimalist Go web framework
	_____________________________________________
	`
)

type Server struct {
	movieService *movie_service.Service
	port string
}

func New(movieService *movie_service.Service,port string) *Server {
	commons.LogInfo(fmt.Sprintf("Initialize echoserver server...,port %v",port))
	return &Server{
		movieService: movieService,
		port:    port,
	}
}

func (s *Server) Run() error {
	e := echo.New()
	e.Server.Addr = ":"+s.port
	e.Pre(middleware.RemoveTrailingSlash())
	e.Pre(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper:                    nil,
		ErrorMessage: "request timeout",
		Timeout:                    time.Second * 2,
	}))

	e.Use(middleware.Recover())

	//Get array of ControllerEcho
	controllers := echo_controller.GetControllers(s.movieService)
	for _, c := range controllers{
		c.Setup(e) //setup each of the controller
	}

	commons.LogInfo(fmt.Sprintf("Initialization done ..."))
	commons.LogInfo(fmt.Sprintf("%v",banner))
	commons.LogInfo(fmt.Sprintf("Echo Server up and listening port %v",s.port))
	return 	graceful.ListenAndServe(e.Server, 10*time.Second) //wait up to 10 second before shutting down
}


