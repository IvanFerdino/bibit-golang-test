package log_repository

import (
	"IvanFerdino/bibit-golang-test/commons"
	"IvanFerdino/bibit-golang-test/internal/model"
	sqlc "IvanFerdino/bibit-golang-test/internal/repository/sqlc_generated"
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"
)

type LogRepo struct {
	db *sql.DB
}

func NewPostgre(db *sql.DB) *LogRepo{
	return &LogRepo{
		db: db,
	}
}


func (repo *LogRepo) LogApiCall(ctx context.Context, wg *sync.WaitGroup,data *model.ApiCall) {
	ctx,cancel:=context.WithTimeout(ctx,3*time.Second)
	defer cancel()

	query:=sqlc.New(repo.db)
	_,err:=query.InsertLog(ctx,sqlc.InsertLogParams{
		HttpRequest: sql.NullString{
			String: data.Request,
			Valid:  true,
		},
		Type:        sql.NullString{
			String: data.Type,
			Valid:  true,
		},
	})

	if err!=nil{
		commons.LogError(fmt.Sprintf("Error insert log: %v\n",err))
	}

	wg.Done()
}
