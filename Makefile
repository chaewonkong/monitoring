# makefile
install:
	@go mod tidy && go mod vendor
.PHONY: install