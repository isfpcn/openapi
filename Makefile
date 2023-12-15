.PHONY:model
.PHONY:douyin
.PHONY:build

VERSION=$(shell git describe --tags --always)


.PHONY: init
# init env
init:
	go install github.com/gogo/protobuf/protoc-gen-gogo@latest
	go install github.com/doeasycode/protoc-gen-fiber@latest
	go install github.com/doeasycode/protoc-gen-swagger@latest


build:
	go build -ldflags "-X main.Version=$(VERSION)"

model:
	go run . --gen_model=true


# build pb
douyin:
	protoc --proto_path=./third_party --proto_path=./ --gogo_out=plugins=grpc:. --fiber_out=. --swagger_out=.  ./api/douyin/*.proto


