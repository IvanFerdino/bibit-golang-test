package movie_search_repository

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

func NewOmbd(url string) *MovieSearchRepo{
	return &MovieSearchRepo{
		url: url,
	}
}

//TODO http get with timeout
//TODO log to postgre

func (repo *MovieSearchRepo) Search(ctx context.Context,keyword string, page int) (*model.MovieSearchResponse,*model.ApiCall,error) {
	ctx,cancel:=context.WithTimeout(ctx,3*time.Second)
	defer cancel()

	u:=fmt.Sprintf("%v/?apikey=%v&s=%v&page=%v",repo.url,repo.key,url.QueryEscape(keyword),page)
	client := &http.Client{}
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		commons.LogError(fmt.Sprintf("Error creating http request: %v\n",err))
		return nil,nil,err
	}
	req = req.WithContext(ctx)


	resp, err := client.Do(req)
	if err != nil {
		commons.LogError(fmt.Sprintf("Error making http request: %v\n",err))
		return nil,nil,err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		commons.LogError(fmt.Sprintf("Error making http response: response code %v\n",resp.StatusCode))
		return nil,nil,err
	}

	body,err:=ioutil.ReadAll(resp.Body)
	if err!=nil{
		commons.LogError(fmt.Sprintf("Error read http response: %v\n",err))
		return nil,nil,err
	}

	result:=model.MovieSearchResponse{}
	err=json.Unmarshal(body,&result)
	if err!=nil{
		commons.LogError(fmt.Sprintf("Error parse http response: %v\n",err))
		return nil,&model.ApiCall{
			Request: u,
			Type:    "GET",
			Code:    resp.StatusCode,
		},err
	}
	return &result,&model.ApiCall{
		Request: u,
		Type:    "GET",
		Code:    resp.StatusCode,
		ErrorMsg: nil,
	},nil
}
