-- name: InsertLog :one
insert into tb_search_log(http_request)
VALUES ($1) returning id;
