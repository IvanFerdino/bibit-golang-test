NOTES: menggunakan postgresql, golang v1.16, sqlc, dan protoc-gen-go protoc-gen-go-grpc

==================================================================================
to run:
1. execute this in db
create table tb_search_log
(
    id serial not null
        constraint tb_search_log_pk
            primary key,
    http_request varchar,
    timestamp timestamp with time zone default timezone('Asia/Jakarta'::text, now())
);

create unique index tb_search_log_id_uindex
	on tb_search_log (id);


2. edit consts in cmd/run_server/main.go //not using env variable
   then run or build cmd/run_server/main.go
   server started at port: 8320 (rest), 8321 (grpc)

==================================================================================
to try grpc client:
run or build cmd/run_grpc_client/main.go

unit test:
unit test available at service layer service_test.go

==================================================================================
1. localhost:8320/movie/search?keyword=Batman&page=1
   localhost:8320/movie/detail?movie_id=tt0372784