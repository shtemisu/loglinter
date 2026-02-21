PLUGIN_PATH=./bin/loglinter.so

build_plugin:
	CGO_ENABLED=1 go build -buildmode=plugin -o $(PLUGIN_PATH) ./cmd/plugin/main.go
make build:
	CGO_ENABLED=1 go build -buildmode=plugin -o $(PLUGIN_PATH) ./cmd/main.go
test_rules:
	go test ./internal/rules

test_rules_expand:
	go test -v ./internal/rules

lint: build_plugin
	golangci-lint run