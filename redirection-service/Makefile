PROTOC = protoc
PROTO_FILES = proto/redirection.proto

generate-grpc:
	$(PROTOC) --go_out=. --go_opt=paths=source_relative --go-grpc_out=.  --go-grpc_opt=paths=source_relative $(PROTO_FILES)

.PHONY: generate-grpc