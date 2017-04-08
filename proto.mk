PROTO_PATH = vendor/github.com/binhq/apis/binhq

.PHONY: proto

proto: ## Generate code from protocol buffer
	@mkdir -p apis
	protoc -I ${PROTO_PATH} ${PROTO_PATH}/githubin/v1alpha1/githubin.proto --go_out=plugins=grpc:apis

envcheck::
	$(call executable_check,protoc,protoc)
	$(call executable_check,protoc-gen-go,protoc-gen-go)
