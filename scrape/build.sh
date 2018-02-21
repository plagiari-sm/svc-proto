#!/usr/bin/bash
protoc -I scrape/ scrape/grpc_scrape.proto --go_out=plugins=grpc:scrape
