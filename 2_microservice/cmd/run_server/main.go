package main

import (
	"IvanFerdino/bibit-golang-test/commons"
	"IvanFerdino/bibit-golang-test/internal/model"
	"IvanFerdino/bibit-golang-test/internal/server/grpc_server"
	"IvanFerdino/bibit-golang-test/internal/server/rest_server"
	"IvanFerdino/bibit-golang-test/internal/service"
	"context"
	"fmt"
	"log"
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
	log.SetOutput(os.Stdout)
	commons.LogInfo(fmt.Sprintf("in this case i will run both server from single binary/main.go.\n but it can be splitted into two main.go. and run both binary (rest server and grpc server)"))
	commons.LogInfo(fmt.Sprintf("Service starting"))

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
	wg.Add(2)

	restsrvr:= rest_server.New(movieSvc,"8320")
	go func() {
		err := restsrvr.Run()
		if err != nil {
			commons.LogError(fmt.Sprintf("%v",err))
			wg.Done()
			os.Exit(1) //exit this service / all goroutine will be killed too. to make sure both rest and grpc server are running.
		}
		commons.LogError(fmt.Sprintf("Server Echo Shutdown"))
		wg.Done()
	}()

	grpcsrvr:= grpc_server.New(movieSvc,"8321")
	go func() {
		err := grpcsrvr.Run(context.Background())
		if err != nil {
			commons.LogError(fmt.Sprintf("%v",err))
			wg.Done()
			os.Exit(1) //exit this service / all goroutine will be killed too. to make sure both rest and grpc server are running.
		}
		commons.LogError(fmt.Sprintf("Server GRPC Shutdown"))
		wg.Done()
	}()
	wg.Wait()
	commons.LogInfo(fmt.Sprintf("Service shutdown"))
}