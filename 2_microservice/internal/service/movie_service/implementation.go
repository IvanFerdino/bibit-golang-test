package movie_service

import (
	"IvanFerdino/bibit-golang-test/commons"
	"IvanFerdino/bibit-golang-test/internal/model"
	"context"
	"fmt"
	"net/http"
	"sync"
)

type movieService struct {
	service *Service
	Ctx   context.Context
}


func NewSvc(ctx context.Context, service *Service) *movieService{
	return &movieService{
		service: service,
		Ctx:     ctx,
	}
}

func (s *movieService) doLog(req string) *sync.WaitGroup{
	wg:=new(sync.WaitGroup)
	wg.Add(1)
	go s.service.LogCall(s.Ctx,wg,&model.ApiCall{
		Request: req,
	})
	return wg
}

func (s *movieService) SearchImpl(param *model.MovieSearchRequest) (*model.MovieSearchResponse,error) {
	err := s.service.validator.Struct(param)
	if err != nil {
		commons.LogError(fmt.Sprintf("request body not valid: %v",err.Error()))
		return nil, &commons.CustomApiError{Message:commons.INVALID_PARAMS, Code:http.StatusBadRequest}
	}

	req:=s.service.MovieSearchRequest(param.Keyword,param.Page)
	wg:=s.doLog(req)

	res,err:=s.service.MovieSearch(s.Ctx,param.Keyword,param.Page)
	if err!=nil{
		return nil, &commons.CustomApiError{Message:commons.SERVER_ERROR, Code:http.StatusInternalServerError}
	}

	wg.Wait() //wait async process to finish. make sure when log is inserted to db, when insert to db take longer time than api call. else there's a probability api call not logged properly to db
	return res,nil
}

func (s *movieService) DetailImpl(param *model.MovieDetailRequest) (*model.MovieDetailResponse,error) {
	err := s.service.validator.Struct(param)
	if err != nil {
		commons.LogError(fmt.Sprintf("request body not valid: %v",err.Error()))
		return nil, &commons.CustomApiError{Message:commons.INVALID_PARAMS, Code:http.StatusBadRequest}
	}

	req:=s.service.MovieDetailRequest(param.MovieId)
	wg:=s.doLog(req)

	res,err:=s.service.MovieDetail(s.Ctx,param.MovieId)
	if err!=nil{
		return nil, &commons.CustomApiError{Message:commons.SERVER_ERROR, Code:http.StatusInternalServerError}
	}

	wg.Wait() //wait async process to finish. make sure when log is inserted to db, when insert to db take longer time than api call. else there's a probability api call not logged properly to db
	return res,nil
}



