package repository

import (
	"IvanFerdino/bibit-golang-test/internal/model"
	"context"
	"sync"
)

type MovieRepository interface {
	Search(ctx context.Context,keyword string, page int) (*model.MovieSearchResponse,error)
	GetSearchRequest(keyword string, page int) string
	Detail(ctx context.Context,movieId string) (*model.MovieDetailResponse,error)
	GetDetailRequest(movieId string) string
}

type LogRepository interface {
	LogApiCall(ctx context.Context, wg *sync.WaitGroup,data *model.ApiCall)
}