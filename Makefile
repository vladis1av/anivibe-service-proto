GEN_FOLDER = ./gen
GEN_GO_FOLDER  = ./gen/go
PROTO_FILES = ./proto/**/**/**/*.proto
REPOSITORY_GEN_GO = github.com/vladis1av/anivibe-service-proto/gen/go
REPOSITORY_GEN_FOLDER = github.com/vladis1av/anivibe-service-proto/gen

.DEFAULT_GOAL := init_go
.PHONY: generate_grpc_go go_mod_init create_gen_folder init_go

create_gen_folder:
	mkdir $(GEN_FOLDER) || true

go_mod_init:
	cd $(GEN_GO_FOLDER); \
	go mod init $(REPOSITORY_GEN_GO) && go mod tidy

generate_grpc_go:
	protoc \
 	--go_out=$(GEN_FOLDER) \
 	--go_opt=module=$(REPOSITORY_GEN_FOLDER) \
	--go-grpc_out=$(GEN_FOLDER) \
 	--go-grpc_opt=module=$(REPOSITORY_GEN_FOLDER) \
 	$(PROTO_FILES) \
	--experimental_allow_proto3_optional

init_go: create_gen_folder generate_grpc_go go_mod_init
