# greet-go-grpc


 protoc -I. --go_out=. --go-grpc_out=. greet/proto/dummy.proto


protoc -I. --go_out=. --go_opt=module=github.com/greet-go-grpc --go-grpc_out=. --go-grpc_opt=module=github.com/greet-go-grpc greet/proto/dummy.proto