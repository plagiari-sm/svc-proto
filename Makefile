protocols:
	protoc -I/usr/local/include -I. -I${GOPATH}/src --go_out=plugins=grpc:. plagiarism/grpc_plagiarism.proto
	protoc -I/usr/local/include -I. -I${GOPATH}/src --go_out=plugins=grpc:. scrape/grpc_scrape.proto
	protoc -I/usr/local/include -I. -I${GOPATH}/src --go_out=plugins=grpc:. compare/grpc_compare.proto
	protoc -I/usr/local/include -I. -I${GOPATH}/src --go_out=plugins=grpc:. enrich/grpc_enrich.proto
