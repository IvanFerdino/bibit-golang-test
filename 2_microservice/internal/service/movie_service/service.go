package movie_service

import (
	"IvanFerdino/bibit-golang-test/internal/model"
	"IvanFerdino/bibit-golang-test/internal/repository"
	v1 "IvanFerdino/bibit-golang-test/pkg/grpc/v1"
	"context"
	"github.com/go-playground/validator/v10"
	"sync"
)

type Service struct {
	movieRepo repository.MovieRepository
	logRepo repository.LogRepository
	validator *validator.Validate
	v1.UnimplementedMovieServiceServer
}

func New(movieRepo repository.MovieRepository, logRepo repository.LogRepository, vldtr *validator.Validate) *Service {
	return &Service{
		movieRepo:                       movieRepo,
		logRepo:                         logRepo,
		validator:                       vldtr,
	}
}

func (s *Service) Search(ctx context.Context, request *v1.MovieSearchRequest) (*v1.MovieSearchResponse, error){
	svc := NewSvc(ctx, s)
	req:=&model.MovieSearchRequest{
		Keyword: request.Keyword,
		Page:    int(request.Page),
	}

	res,err:=svc.SearchImpl(req)
	if err!=nil{
		return nil,err
	}

	movies:=make([]*v1.Movie,0)
	for _,d:=range res.Movies {
		movies=append(movies,&v1.Movie{
			Title:  d.Title,
			Year:   d.Year,
			ImdbId: d.ImdbId,
			Type:   d.Type,
			Poster: d.Poster,
		})
	}

	return &v1.MovieSearchResponse{
		Movies:      movies,
		TotalResult: res.TotalResult,
		Response:    res.Response,
	}, nil
}

func (s *Service) Detail(ctx context.Context, request *v1.MovieDetailRequest) (*v1.MovieDetailResponse, error){
	svc := NewSvc(ctx, s)
	req:=&model.MovieDetailRequest{
		MovieId: request.MovieId,
	}

	res,err:=svc.DetailImpl(req)
	if err!=nil{
		return nil,err
	}

	ratings:=make([]*v1.MovieRating,0)
	for _,d:=range res.Ratings{
		ratings=append(ratings,&v1.MovieRating{
			Source: d.Source,
			Value:  d.Value,
		})
	}

	return &v1.MovieDetailResponse{
		Title:      res.Title,
		Year:       res.Year,
		Rated:      res.Rated,
		Released:   res.Released,
		Runtime:    res.Runtime,
		Genre:      res.Genre,
		Director:   res.Director,
		Writer:     res.Writer,
		Actors:     res.Actors,
		Plot:       res.Plot,
		Language:   res.Language,
		Country:    res.Country,
		Awards:     res.Awards,
		Poster:     res.Poster,
		Ratings:    ratings,
		Metascore:  res.Metascore,
		ImdbRating: res.ImdbRating,
		ImdbVotes:  res.ImdbVotes,
		ImdbID:     res.ImdbID,
		Type:       res.Type,
		DVD:        res.DVD,
		BoxOffice:  res.BoxOffice,
		Production: res.Production,
		Website:    res.Website,
		Response:   res.Response,
	}, nil
}


//helper method to access repo(s)
func (s *Service) MovieSearch(ctx context.Context,keyword string, page int) (*model.MovieSearchResponse,error){
	return s.movieRepo.Search(ctx,keyword,page)
}

func (s *Service) MovieSearchRequest(keyword string, page int) string{
	return s.movieRepo.GetSearchRequest(keyword,page)
}

func (s *Service) MovieDetail(ctx context.Context,movieId string) (*model.MovieDetailResponse,error) {
	return s.movieRepo.Detail(ctx,movieId)
}

func (s *Service) MovieDetailRequest(movieId string) string{
	return s.movieRepo.GetDetailRequest(movieId)
}

func (s *Service) LogCall(ctx context.Context, wg *sync.WaitGroup, data *model.ApiCall) {
	s.logRepo.LogApiCall(ctx, wg, data)
}

