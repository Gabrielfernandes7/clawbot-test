#!/bin/bash

set -e

echo "Configurando Crabe..."

CRABE_DIR="$HOME/.crabe"
INSTALL_DIR="$HOME/.local/bin"
PROJECT_DIR="$(pwd)"

mkdir -p "$CRABE_DIR"

# Config padrão
if [ ! -f "$CRABE_DIR/config.json" ]; then
  echo '{ "model": "llama3.2:3b" }' > "$CRABE_DIR/config.json"
fi

# Garantir permissão
chmod +x "$PROJECT_DIR/cli/crabe.sh"

# Criar pasta bin
mkdir -p "$INSTALL_DIR"

# 🔥 USAR SYMLINK (ESSENCIAL)
ln -sf "$PROJECT_DIR/cli/crabe.sh" "$INSTALL_DIR/crabe"

echo "✅ Crabe instalado corretamente (symlink)"
echo "👉 Teste com: crabe version"