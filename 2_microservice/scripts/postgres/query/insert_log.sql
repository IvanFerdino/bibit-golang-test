-- name: InsertLog :one
insert into tb_search_log(http_request, type, code, error)
VALUES ($1,$2,$3,$4) returning id;
