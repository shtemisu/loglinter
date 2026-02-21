.PHONY: build plugin test_rules test_rules_extend

build:
	go build -o loglinter cmd/main.go 

plugin:
	@golangci-lint custom -v

test_rules:
	@go test ./internal/rules

test_rules_extend:
	@go test -v ./internal/rules

rebuild_plugin: clean	

clean_plugin:
	@rm custom-gcl

clean_exe:
	@rm loglinter

clean all: clean_plugin clean_exe
