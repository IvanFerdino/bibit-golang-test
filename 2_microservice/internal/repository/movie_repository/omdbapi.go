package movie_repository

import (
	"IvanFerdino/bibit-golang-test/commons"
	"IvanFerdino/bibit-golang-test/internal/model"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type MovieSearchRepo struct {
	url string
	key string
}

func NewOmbd(url string,key string) *MovieSearchRepo{
	return &MovieSearchRepo{
		url: url,
		key: key,
	}
}

func (repo *MovieSearchRepo) GetSearchRequest(keyword string, page int) string {
	return fmt.Sprintf("%v/?apikey=%v&s=%v&page=%v",repo.url,repo.key,url.QueryEscape(keyword),page)
}
func (repo *MovieSearchRepo) GetDetailRequest(movieId string) string {
	return fmt.Sprintf("%v/?apikey=%v&i=%v",repo.url,repo.key,url.QueryEscape(movieId))
}

func (repo *MovieSearchRepo) Search(ctx context.Context,keyword string, page int) (*model.MovieSearchResponse,error) {
	ctx,cancel:=context.WithTimeout(ctx,3*time.Second)
	defer cancel()

	u:=fmt.Sprintf("%v/?apikey=%v&s=%v&page=%v",repo.url,repo.key,url.QueryEscape(keyword),page)
	client := &http.Client{}
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		commons.LogError(fmt.Sprintf("Error creating http request: %v\n",err))
		return nil,err
	}
	req = req.WithContext(ctx)


	resp, err := client.Do(req)
	if err != nil {
		commons.LogError(fmt.Sprintf("Error making http request: %v\n",err))
		return nil,err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		commons.LogError(fmt.Sprintf("Error making http response: response code %v\n",resp.StatusCode))
		return nil,err
	}

	body,err:=ioutil.ReadAll(resp.Body)
	if err!=nil{
		commons.LogError(fmt.Sprintf("Error read http response: %v\n",err))
		return nil,err
	}

	result:=model.MovieSearchResponse{}
	err=json.Unmarshal(body,&result)
	if err!=nil{
		commons.LogError(fmt.Sprintf("Error parse http response: %v\n",err))
		return nil,err
	}
	return &result,nil
}


func (repo *MovieSearchRepo) Detail(ctx context.Context,movieId string) (*model.MovieDetailResponse,error){
	ctx,cancel:=context.WithTimeout(ctx,3*time.Second)
	defer cancel()

	u:=fmt.Sprintf("%v/?apikey=%v&i=%v",repo.url,repo.key,url.QueryEscape(movieId))
	client := &http.Client{}
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		commons.LogError(fmt.Sprintf("Error creating http request: %v\n",err))
		return nil,err
	}
	req = req.WithContext(ctx)


	resp, err := client.Do(req)
	if err != nil {
		commons.LogError(fmt.Sprintf("Error making http request: %v\n",err))
		return nil,err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		commons.LogError(fmt.Sprintf("Error making http response: response code %v\n",resp.StatusCode))
		return nil,err
	}

	body,err:=ioutil.ReadAll(resp.Body)
	if err!=nil{
		commons.LogError(fmt.Sprintf("Error read http response: %v\n",err))
		return nil,err
	}

	result:=model.MovieDetailResponse{}
	err=json.Unmarshal(body,&result)
	if err!=nil{
		commons.LogError(fmt.Sprintf("Error parse http response: %v\n",err))
		return nil,err
	}
	return &result,nil

	return nil, nil
}