sync:
	GOLANG_PATH=$$(echo $$GOPATH); \
	rm -rf -- "$${GOLANG_PATH}/src/github.com/cixtor/slackapi"; \
	cp -r -- "/Users/yorman/Projects/slackapi" "$${GOLANG_PATH}/src/github.com/cixtor/slackapi"

update:
	GOLANG_PATH=$$(echo $$GOPATH); \
	rm -rf -- "$${GOLANG_PATH}/bin/slackapi"; \
	rm -rf -- "$${GOLANG_PATH}/src/github.com/cixtor/slackapi"; \
	rm -rf -- "$${GOLANG_PATH}/pkg/darwin_amd64/github.com/cixtor/slackapi.a"; \
	go get -u github.com/cixtor/slackapi/slackapi
