.PHONY: all build clean test run run-gui

# 変数定義
BINARY_NAME=now
BINARY_DIR=bin

# デフォルトターゲット
all: build

# ビルドディレクトリの作成
$(BINARY_DIR):
	mkdir -p $(BINARY_DIR)

# ビルド
build: $(BINARY_DIR)
	go build -o $(BINARY_DIR)/$(BINARY_NAME)
	@if [ "$(shell uname)" = "Darwin" ]; then \
		codesign --force --deep --sign - $(BINARY_DIR)/$(BINARY_NAME); \
	fi

# 実行 (CLIモード)
run:
	@if [ -f $(BINARY_DIR)/$(BINARY_NAME) ]; then \
		$(BINARY_DIR)/$(BINARY_NAME); \
	else \
		go run main.go; \
	fi

# 実行 (GUIモード)
run-gui:
	@if [ -f $(BINARY_DIR)/$(BINARY_NAME) ]; then \
		$(BINARY_DIR)/$(BINARY_NAME) -gui; \
	else \
		go run main.go -gui; \
	fi

# テスト実行
test:
	go test ./...

# 依存関係の更新
deps:
	go mod tidy

# クリーンアップ
clean:
	go clean
	rm -rf $(BINARY_DIR)
	rm -rf *.app

# macOS用アプリケーションバンドルの作成
bundle:
	@if [ "$(shell uname)" = "Darwin" ]; then \
		go install fyne.io/fyne/v2/cmd/fyne@latest && \
		fyne package -os darwin; \
	else \
		echo "This target is only supported on macOS"; \
		exit 1; \
	fi

# ヘルプ
help:
	@echo "Available targets:"
	@echo "  make          - Build the application with code signing (on macOS)"
	@echo "  make build    - Same as above"
	@echo "  make run      - Run the application in CLI mode"
	@echo "  make run-gui  - Run the application in GUI mode"
	@echo "  make test     - Run tests"
	@echo "  make deps     - Update dependencies"
	@echo "  make clean    - Clean build artifacts"
	@echo "  make bundle   - Create macOS application bundle (macOS only)"
	@echo "  make help     - Show this help message"
