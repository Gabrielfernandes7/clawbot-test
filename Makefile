# Makefile - Crabe CLI (Migração para Go)

.PHONY: build install clean test doctor init help remove-old

BINARY_NAME = crabe
INSTALL_DIR = $(HOME)/.local/bin

# Build o binário
build:
	@echo "🔨 Compilando binário Go..."
	go build -o $(BINARY_NAME) ./cmd/crabe
	@echo "✅ Binário criado: ./$(BINARY_NAME)"

# Remove possíveis links ou binários antigos problemáticos
remove-old:
	@echo "🧹 Removendo versões antigas..."
	rm -f $(INSTALL_DIR)/$(BINARY_NAME)
	rm -f ./$(BINARY_NAME)

# Instala de forma segura
install: remove-old build
	@echo "📦 Instalando crabe em $(INSTALL_DIR)..."
	mkdir -p $(INSTALL_DIR)
	cp -f $(BINARY_NAME) $(INSTALL_DIR)/
	chmod +x $(INSTALL_DIR)/$(BINARY_NAME)
	@echo "✅ Crabe instalado com sucesso!"
	@echo "   Teste com: crabe doctor"

# Limpeza
clean:
	@echo "🧹 Limpando..."
	rm -f $(BINARY_NAME)
	go clean

# Comandos de teste rápidos
doctor: build
	./$(BINARY_NAME) doctor

init: build
	./$(BINARY_NAME) init

init-force: build
	./$(BINARY_NAME) init --force

# Ajuda
help:
	@echo "Comandos disponíveis:"
	@echo "  make build          → Compila o binário"
	@echo "  make install        → Compila + instala (recomendado)"
	@echo "  make clean          → Remove binários"
	@echo "  make doctor         → Executa crabe doctor"
	@echo "  make init           → Executa crabe init"
	@echo "  make init-force     → Executa crabe init --force"
	@echo "  make remove-old     → Remove link antigo quebrado"
	@echo "  make help           → Mostra esta ajuda"