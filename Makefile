protocols:
	protoc -I/usr/local/include -I. -I${GOPATH}/src --go_out=:. --go-grpc_out=:. plagiarism/grpc_plagiarism.proto
	protoc -I/usr/local/include -I. -I${GOPATH}/src --go_out=:. --go-grpc_out=:. scrape/grpc_scrape.proto
	protoc -I/usr/local/include -I. -I${GOPATH}/src --go_out=:. --go-grpc_out=:. compare/grpc_compare.proto
	protoc -I/usr/local/include -I. -I${GOPATH}/src --go_out=:. --go-grpc_out=:. enrich/grpc_enrich.proto
