package cmd

import (
	"IvanFerdino/bibit-golang-test/internal/model"
	"IvanFerdino/bibit-golang-test/internal/service"
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


}