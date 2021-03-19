package log_repository

import (
	"IvanFerdino/bibit-golang-test/commons"
	"IvanFerdino/bibit-golang-test/internal/model"
	sqlc "IvanFerdino/bibit-golang-test/internal/repository/sqlc_generated"
	"context"
	"database/sql"
	"fmt"
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


func (repo *LogRepo) LogApiCall(ctx context.Context,data *model.ApiCall) {
	ctx,cancel:=context.WithTimeout(ctx,3*time.Second)
	defer cancel()

	errMsgValid:=false
	if data.ErrorMsg == nil {
		errMsgValid=true
	}

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
		Code:        sql.NullInt32{
			Int32: int32(data.Code),
			Valid: true,
		},
		Error: sql.NullString{
			String: *data.ErrorMsg,
			Valid:  errMsgValid,
		},
	})

	if err!=nil{
		commons.LogError(fmt.Sprintf("Error insert log: %v\n",err))
	}
}
