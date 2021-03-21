package movie_service

import (
	"IvanFerdino/bibit-golang-test/internal/model"
	"IvanFerdino/bibit-golang-test/internal/repository/mockgen_generated/mock_repository"
	v1 "IvanFerdino/bibit-golang-test/pkg/grpc/v1"
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/url"
	"sync"
	"testing"
)

//cannot run, because of waitgroup issue in unit test
func Test_Search_Movie(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock:=mock_repository.NewMockMovieRepository(ctrl)
	mock2:=mock_repository.NewMockLogRepository(ctrl)
	vldtr:=validator.New()
	t.Run("SERACH MOVIE SUCCESS", func(t *testing.T) {
		request:=&v1.MovieSearchRequest{
			Keyword: "Batman",
			Page:    1,
		}
		expectedRepoResponse:=&model.MovieSearchResponse{
			Movies:      nil,//let say movie response return nil movie
			TotalResult: "413",
			Response:    "TRUE",
		}
		expectedResponse:=&v1.MovieSearchResponse{
			Movies:      nil,
			TotalResult: "413",
			Response:    "TRUE",
		}
		apiRequest:=fmt.Sprintf("http://www.omdbapi.com/?apikey=faf7e5bb&s=%v&page=%v",url.QueryEscape(request.Keyword),request.Page)

		//mock repo call
			mock.EXPECT().GetSearchRequest("Batman",1).Return(apiRequest).Times(1)

			//ISSUE HERE
			wg:=new(sync.WaitGroup)
			wg.Add(1)
			mock2.EXPECT().LogApiCall(context.Background(),wg,&model.ApiCall{
				Request: apiRequest,
			})
			wg.Wait()

			mock.EXPECT().Search(context.Background(),"Batman",1).Return(expectedRepoResponse,nil).Times(1)

		//create service
			svc:=New(mock,mock2,vldtr)
		//execute method
			res,err:=svc.Search(context.Background(),request)
			assert.Nil(t, err)
			assert.Equal(t, expectedResponse,res)

	})
}