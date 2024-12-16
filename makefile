# Nome do binário
BINARY_NAME=myapp

# Diretório de saída
BUILD_DIR=build

# Variáveis de compilação
GOFLAGS=-ldflags="-s -w"  # Remove símbolos de depuração e reduz o tamanho do binário

# Comandos principais
all: build-windows build-ubuntu

build-windows:
	@echo "Building for Windows..."
	GOOS=windows GOARCH=amd64 go build $(GOFLAGS) -o $(BUILD_DIR)/windows/$(BINARY_NAME).exe .
	@echo "Binary for Windows created: $(BUILD_DIR)/windows/$(BINARY_NAME).exe"

build-ubuntu:
	@echo "Building for Ubuntu..."
	GOOS=linux GOARCH=amd64 go build $(GOFLAGS) -o $(BUILD_DIR)/ubuntu/$(BINARY_NAME) .
	@echo "Binary for Ubuntu created: $(BUILD_DIR)/ubuntu/$(BINARY_NAME)"

clean:
	@echo "Cleaning up..."
	rm -rf $(BUILD_DIR)
	@echo "Build directory cleaned."

# Limpeza e execução
run-ubuntu: build-ubuntu
	@echo "Running on Ubuntu..."
	./$(BUILD_DIR)/ubuntu/$(BINARY_NAME)

run-windows: build-windows
	@echo "Running on Windows is not directly supported from here. Binary available."

.PHONY: all build-windows build-ubuntu clean run-ubuntu run-windows
