package main

import (
	"IvanFerdino/bibit-golang-test/commons"
	v1 "IvanFerdino/bibit-golang-test/pkg/grpc/v1"
	"IvanFerdino/bibit-golang-test/pkg/grpcclient"
	"fmt"
	"google.golang.org/grpc/status"
	"log"
	"os"
)

const (
	GRPCSERVERHOST = "localhost"
	GRPCSERVERPORT = "8321"
)

//GRPC CLIENT USAGE EXAMPLE
func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	commons.LogInfo(fmt.Sprintf("sample rpc client"))


	search,err:=search(grpcclient.New(GRPCSERVERHOST,GRPCSERVERPORT))
	if err!=nil{
		commons.LogError(fmt.Sprintf("Error %v\n",err))
		os.Exit(1)
	}
	commons.LogInfo(fmt.Sprintf("response: %v", search))



	det,err:=detail(grpcclient.New(GRPCSERVERHOST,GRPCSERVERPORT))
	if err!=nil{
		commons.LogError(fmt.Sprintf("Error %v\n",err))
		os.Exit(1)
	}
	commons.LogInfo(fmt.Sprintf("response: %v", det))
}
func detail(cl *grpcclient.ClientConnection) (*v1.MovieDetailResponse,error){
	param:=&v1.MovieDetailRequest{
		MovieId: "tt0372784",
	}

	//from pkg/grpcclient package
	resp,err:=grpcclient.Detail(cl,param)
	if err!=nil{
		commons.LogError(fmt.Sprintf("response: %v, error: %v", resp, status.Convert(err)))
		return nil, status.Convert(err).Err()
	}
	return resp,nil
}
func search(cl *grpcclient.ClientConnection) (*v1.MovieSearchResponse,error){
	param:=&v1.MovieSearchRequest{
		Keyword: "Batman",
		Page:    1,
	}

	//from pkg/grpcclient package
	resp,err:=grpcclient.Search(cl,param)
	if err!=nil{
		commons.LogError(fmt.Sprintf("response: %v, error: %v", resp, status.Convert(err)))
		return nil, status.Convert(err).Err()
	}
	return resp,nil
}
