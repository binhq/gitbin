PROTO_PATH = vendor/github.com/binhq/apis/binhq

.PHONY: proto

proto: ## Generate code from protocol buffer
	@mkdir -p apis
	protowrap -I ${PROTO_PATH} ${PROTO_PATH}/githubin/v1alpha1/githubin.proto --go_out=.:apis

envcheck::
	$(call executable_check,protoc,protoc)
	$(call executable_check,protowrap,protowrap)
