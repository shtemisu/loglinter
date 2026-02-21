.PHONY: build plugin test_rules test_rules_extend

build:
	go build -o loglinter cmd/main.go 

plugin:
	@golangci-lint custom -v

test_rules:
	@go test ./internal/rules

test_rules_extend:
	@go test -v ./internal/rules

test_plugin:
	@./custom-gcl run -c .golangci.yml ./analyzer/testdata/russian.go
	@./custom-gcl run -c .golangci.yml ./analyzer/testdata/only_english.go
	@./custom-gcl run -c .golangci.yml ./analyzer/testdata/sensetive_data.go
	@./custom-gcl run -c .golangci.yml ./analyzer/testdata/specchars.go
	@./custom-gcl run -c .golangci.yml ./analyzer/testdata/testZap.go

rebuild_plugin: clean	

clean_plugin:
	@rm custom-gcl

clean_exe:
	@rm loglinter

clean all: clean_plugin clean_exe
