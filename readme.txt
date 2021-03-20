to generate pb.go:
1. install protoc compiler
2. go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
3. go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
4. go get -u google.golang.org/grpc
   go get -u google.golang.org/grpc/codes
   go get -u google.golang.org/grpc/status
5. ./generate_pb_go.sh


1. localhost:8320/movie/search?keyword=Batman&page=1
   localhost:8320/movie/detail?movie_id=tt0372784