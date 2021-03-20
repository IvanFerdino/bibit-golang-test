package main

import (
	"IvanFerdino/bibit-golang-test/commons"
	"IvanFerdino/bibit-golang-test/internal/model"
	"IvanFerdino/bibit-golang-test/internal/server/rest_server"
	"IvanFerdino/bibit-golang-test/internal/service"
	"fmt"
	"os"
	"sync"
)

const (
	APIURL = "http://www.omdbapi.com"
	APIKEY = "faf7e5bb"
	DBHOST = ""
	DBUSERNAME = ""
	DBPASSWORD = ""
	DBPORT = ""
	DBSCHEMA = ""
)

func main() {
	commons.LogInfo(fmt.Sprintf("in this case i will run both server from single binary/main.go.\n but it can be splitted into two main.go. and run both binary (rest server and grpc server)"))
	movieSvc,db:=service.InitService(&model.InitModel{
		ApiUrl:     APIURL,
		ApiKey:     APIKEY,
		DbHost:     DBHOST,
		DbUsername: DBUSERNAME,
		DbPassword: DBPASSWORD,
		DbPort:     DBPORT,
		DbSchema:   DBSCHEMA,
	})
	defer db.Close()

	wg := new(sync.WaitGroup)
	wg.Add(1)


	server:= rest_server.New(movieSvc,"8320")
	go func() {
		err := server.Run()
		if err != nil {
			commons.LogError(fmt.Sprintf("%v",err))
			wg.Done()
			os.Exit(1) //exit this service / all goroutine will be killed too. to make sure both rest and grpc server are running.
		}
		commons.LogError(fmt.Sprintf("Server Echo Shutdown"))
		wg.Done()
	}()

	wg.Wait()
}