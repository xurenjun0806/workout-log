# ====== Config ======
APP        := workout-log
CMD_DIR    := ./cmd/api
BIN_DIR    := bin
BIN        := $(BIN_DIR)/$(APP)
PKG        := ./...
GO         := go

# ====== Recipes ======
.PHONY: help run build test cover fmt vet tidy deps clean generate

verify: fmt vet tidy

run: verify ## 実行（go run）
	$(GO) run $(CMD_DIR)

build: verify ## ビルド（バイナリ: $(BIN)）
	@mkdir -p $(BIN_DIR)
	$(GO) build -trimpath -o $(BIN) $(CMD_DIR)
	@echo "→ Built $(BIN)"

test: ## テスト（raceあり）
	$(GO) test -race -count=1 $(PKG)

cover: ## カバレッジ測定
	$(GO) test -race -covermode=atomic -coverprofile=coverage.out $(PKG)
	@$(GO) tool cover -func=coverage.out | tail -n 1

fmt: ## go fmt
	$(GO) fmt $(PKG)

vet: ## go vet
	$(GO) vet $(PKG)

tidy: ## go mod tidy
	$(GO) mod tidy

clean: ## 生成物を削除
	rm -rf $(BIN_DIR) coverage.out
