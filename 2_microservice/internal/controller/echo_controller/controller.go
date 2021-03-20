package echo_controller

import (
	"IvanFerdino/bibit-golang-test/internal/service/movie_service"
	"github.com/labstack/echo/v4"
)

type ControllerEcho interface {
	Setup(e *echo.Echo)
}

func GetControllers(movieService *movie_service.Service) []ControllerEcho {
	movieController:=NewMovieController(movieService)
	//init more controller if needed ...
	return []ControllerEcho{movieController}
}
