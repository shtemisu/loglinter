BINARY_PATH = ./bin/loglinter
PLUGIN_PATH = ./bin/loglinter.so

.PHONY: build build_bin build_plugin test_rules test_rules_expand lint

build: build_bin build_plugin

build_bin:
	mkdir -p ./bin
	CGO_ENABLED=0 go build -o $(BINARY_PATH) ./cmd/main.go
	chmod +x $(BINARY_PATH)

build_plugin:
	mkdir -p ./bin
	CGO_ENABLED=1 go build -buildmode=plugin -o $(PLUGIN_PATH) ./cmd/plugin/plugin.go

test_rules:
	go test ./internal/rules

test_rules_expand:
	go test -v ./internal/rules

lint: build_plugin
	golangci-lint run --config .golangci.yml ./analyzer/testdata/...

lint-bin: build_bin
	$(BINARY_PATH) ./analyzer/testdata/...

clean:
	rm -f $(BINARY_PATH) $(PLUGIN_PATH)
	go clean -testcache

help:
	@echo "Available targets:"
	@echo "  build          - build binary and plugin"
	@echo "  build_bin      - build binary only"
	@echo "  build_plugin   - build plugin only"
	@echo "  test_rules     - run rules tests"
	@echo "  lint           - run linter with golangci-lint"
	@echo "  lint-bin       - run binary directly"
	@echo "  clean          - remove built files"