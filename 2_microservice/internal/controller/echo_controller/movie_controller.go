package echo_controller

import (
	"IvanFerdino/bibit-golang-test/commons"
	"IvanFerdino/bibit-golang-test/internal/controller"
	"IvanFerdino/bibit-golang-test/internal/model"
	"IvanFerdino/bibit-golang-test/internal/service/movie_service"
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type movieController struct {
	movieService *movie_service.Service
}

func NewMovieController(movieService *movie_service.Service) ControllerEcho {
	return &movieController{
		movieService: movieService,
	}
}

func (ctrl *movieController) Setup(e *echo.Echo) {
	authPrefix := e.Group("/movie")
	authPrefix.GET("/search",ctrl.Search)
	authPrefix.GET("/detail",ctrl.Detail)
}

func (ctrl *movieController) Search(e echo.Context) error {
	ctx,cancel:=context.WithCancel(e.Request().Context())
	defer cancel()

	keyword := e.QueryParam("keyword")
	page,err:=strconv.Atoi(e.QueryParam("page"))
	if err!=nil{
		commons.LogError(fmt.Sprintf("request body not valid %v",err.Error()))
		return e.JSON(controller.HandleRestError(&commons.CustomApiError{Message:commons.INVALID_PARAMS, Code:http.StatusBadRequest}))
	}

	svc:=movie_service.NewSvc(ctx,ctrl.movieService)
	resp,err:=svc.SearchImpl(&model.MovieSearchRequest{
		Keyword: keyword,
		Page:    page,
	})
	if err!=nil {
		return e.JSON(controller.HandleRestError(err))
	}
	return e.JSON(http.StatusOK,controller.HandleRestSuccess(resp))
}

func (ctrl *movieController) Detail(e echo.Context) error {
	ctx,cancel:=context.WithCancel(e.Request().Context())
	defer cancel()

	movieId := e.QueryParam("movie_id")

	svc:=movie_service.NewSvc(ctx,ctrl.movieService)
	resp,err:=svc.DetailImpl(&model.MovieDetailRequest{
		MovieId: movieId,
	})
	if err!=nil {
		return e.JSON(controller.HandleRestError(err))
	}
	return e.JSON(http.StatusOK,controller.HandleRestSuccess(resp))}