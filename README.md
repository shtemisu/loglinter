# loglinter
Implementation of log linter using by Go

[![Go Version](https://img.shields.io/badge/Go-1.26+-00ADD8?logo=go)](https://golang.org)
[![golangci-lint](https://img.shields.io/badge/golangci--lint-2.10.1+-blue)](https://golangci-lint.run)

## RULES

- **Checks register of first char** — log messages must start with a lowercase letter
- **Checks English language** — messages must contain only English (latin) characters
- **Checks special characters** — no emojis, multiple punctuation marks (!!!, ...), or forbidden symbols
- **Checks sensitive data** — detects potential passwords, tokens, API keys in logs

## Supported Loggers

- **log/slog**
- **go.uber.org/zap**

## Test rules
```
make test_rules
```
or extended output
```
make test_rules_extend
```

## Quick start
for a separate linter:
```
make build
```
and
```
./loglinter ./path/to/your/go_file.go
```
for module plugin system:

```
golangci-lint custom -v
./custom-gcl run -c .golangci.yml ./path/to/your/go_file.go
``` 

### Plugin test
```
make test_plugin
```