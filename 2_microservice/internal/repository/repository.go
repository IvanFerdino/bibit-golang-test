package repository

import (
	"IvanFerdino/bibit-golang-test/internal/model"
	"context"
)

type MovieSearchRepository interface {
	Search(ctx context.Context,keyword string, page int) (*model.MovieSearchResponse,*model.ApiCall,error)
}

type LogRepository interface {
	LogApiCall(ctx context.Context,data *model.ApiCall)
}