build:
	protoc -I. --go_out=./ --micro_out=./ proto/*.proto
	protoc-go-inject-tag -input=./proto/user.pb.go 
	protoc-go-inject-tag -input=./proto/book.pb.go