package service

import (
	"IvanFerdino/bibit-golang-test/commons"
	"IvanFerdino/bibit-golang-test/internal/model"
	"IvanFerdino/bibit-golang-test/internal/repository/log_repository"
	"IvanFerdino/bibit-golang-test/internal/repository/movie_repository"
	"IvanFerdino/bibit-golang-test/internal/service/movie_service"
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"
)

func LoadDb(i *model.InitModel) *sql.DB {
	connectionString := fmt.Sprintf(
		"postgres://%v:%v@%v:%v/%v?sslmode=disable",
		i.DbUsername, i.DbPassword, i.DbHost, i.DbPort, i.DbSchema,
	)
	db,err:=sql.Open("postgres",connectionString)
	if err != nil {
		commons.LogError(fmt.Sprintf("Cant open db connection %v",err))
		os.Exit(1) //kill service if cant connect ot db
	}

	db.SetMaxIdleConns(3) //should be < maxopenconns
	db.SetMaxOpenConns(50) //max connection for this db conn
	db.SetConnMaxIdleTime(5*time.Minute) //max time connection bisa idle
	db.SetConnMaxLifetime(10*time.Minute)//connection age. if > age, killed

	pingDb(db)
	return db
}

func pingDb(db *sql.DB){
	ctx,cancel:=context.WithTimeout(context.Background(),time.Second * 15)
	defer cancel()
	query := "SELECT 1"

	rows, err := db.QueryContext(ctx,query)
	defer rows.Close()
	if err!=nil{
		db.Close()
		commons.LogError(fmt.Sprintf("Cant ping db connection %v",err))
		os.Exit(1)
	}

	t:=30*time.Second
	commons.LogInfo(fmt.Sprintf("KEEP ALIVE DB CONNECTION EVERY %v second\n",t))
	ticker := time.NewTicker(t)
	go func() {
		for _ = range ticker.C {
			ctx,cancel:=context.WithTimeout(context.Background(),15*time.Second)
			defer cancel()

			rows, err := db.QueryContext(ctx,query)
			defer rows.Close()
			if err!=nil{
				db.Close()
				commons.LogError(fmt.Sprintf("KEEP ALIVE | Cant ping db connection %v",err))
				os.Exit(1)
			}
			commons.LogInfo(fmt.Sprintf("KEEP ALIVE DB CONNECTION"))
		}
	}()
}

func InitService(initData *model.InitModel) (*movie_service.Service,*sql.DB) {
	postgredb:=LoadDb(initData)

	movieRepo:=movie_repository.NewOmbd(initData.ApiUrl,initData.ApiKey)
	logRepo:=log_repository.NewPostgre(postgredb)

	movieSvc:=movie_service.New(movieRepo,logRepo)

	//3. return initialized services
	return movieSvc,postgredb
}
