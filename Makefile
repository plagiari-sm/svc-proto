get:
	go get -u github.com/golang/protobuf/protoc-gen-go

protocols:
		protoc -I/usr/local/include -I. \
 -I${GOPATH}/src \
 --go_out=plugins=grpc:. \
 plagiarism/grpc_plagiarism.proto
		protoc -I/usr/local/include -I. \
 -I${GOPATH}/src \
 --go_out=plugins=grpc:. \
 scrape/grpc_scrape.proto

 build: get protocols