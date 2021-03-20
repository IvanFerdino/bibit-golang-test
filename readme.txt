NOTES: menggunakan postgresql, golang v1.16, sqlc, dan protoc-gen-go protoc-gen-go-grpc

to generate pb.go if needed:
1. install protoc compiler
2. go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
3. go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
4. go get -u google.golang.org/grpc
   go get -u google.golang.org/grpc/codes
   go get -u google.golang.org/grpc/status
5. ./generate_pb_go.sh

to generate .go from sqlc:
1. ./generate_sqlc.sh

to run:
edit consts in cmd/run_server/main.go //not using env variable
then run or build cmd/run_server/main.go
server started at port: 8320 (rest), 8321 (grpc)

to try grpc client:
run or build cmd/run_grpc_client/main.go

unit test:
unit test available at service layer service_test.go

1. localhost:8320/movie/search?keyword=Batman&page=1
   localhost:8320/movie/detail?movie_id=tt0372784