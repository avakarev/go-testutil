tidy:
	@echo ">> Tidying..."
	@go mod tidy

fmt:
	@echo ">> Formatting..."
	@go fmt ./...

vet:
	@echo ">> Vetting..."
	@go vet ./...

lint:
	@echo ">> Running revive..."
	@revive -config .revive.toml -formatter friendly ./...
	@echo ">> Running staticcheck..."
	@staticcheck ./...

sec:
	@echo ">> Auditing..."
	@gosec -quiet -tests ./...

test:
	@echo ">> Running tests..."
	@go test -v -race ./...
.PHONY: test

setup-ci:
	@GO111MODULE=off go get -u github.com/myitcv/gobin
	@gobin github.com/mgechev/revive
	@gobin github.com/securego/gosec/v2/cmd/gosec
	@gobin honnef.co/go/tools/cmd/staticcheck@latest

ci: lint vet sec test
