# protoc is bash function
#
# function protoc () {
#  docker run --rm -it -v $PWD:$PWD -w $PWD znly/protoc $@
# }
build:
	@#protoc -I src src/echo.proto --go_out=plugins=grpc:build
	docker run --rm -it -v ${PWD}:${PWD} -w ${PWD} znly/protoc -I src src/echo.proto --go_out=plugins=grpc:build

.PHONY: build
