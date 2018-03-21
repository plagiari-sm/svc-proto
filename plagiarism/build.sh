#!/usr/bin/bash
protoc -I plagiarism/ plagiarism/grpc_plagiarism.proto --go_out=plugins=grpc:plagiarism
