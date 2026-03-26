#!/bin/bash

set -e

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
BASE_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"

CRABE_DIR="$HOME/.crabe"
INSTALL_DIR="$HOME/.local/bin"
PROJECT_DIR="$BASE_DIR"

COLORS="$BASE_DIR/cli/colors.sh"

# Validação defensiva
if [ ! -f "$COLORS" ]; then
  echo "Erro: colors.sh não encontrado em $COLORS"
  exit 1
fi

source "$COLORS"

log_highlight "Configurando Crabe..."

mkdir -p "$CRABE_DIR"
log_info "Diretório criado/verificado: $CRABE_DIR"

if [ ! -f "$CRABE_DIR/config.json" ]; then
  echo '{ "model": "llama3.2:1b" }' > "$CRABE_DIR/config.json"
  log_info "Config padrão criada"
else
  log_warn "Config já existe, mantendo atual"
fi

if [ ! -f "$PROJECT_DIR/cli/crabe.sh" ]; then
  log_error "crabe.sh não encontrado em: $PROJECT_DIR/cli/"
  exit 1
fi

chmod +x "$PROJECT_DIR/cli/crabe.sh"
log_info "Permissão aplicada ao CLI"

mkdir -p "$INSTALL_DIR"
log_info "Diretório bin verificado: $INSTALL_DIR"

ln -sf "$PROJECT_DIR/cli/crabe.sh" "$INSTALL_DIR/crabe"

log_info "Symlink criado: $INSTALL_DIR/crabe"

SHELL_CONFIG=""

# Detectar shell
if [[ "$SHELL" == *"zsh"* ]]; then
  SHELL_CONFIG="$HOME/.zshrc"
elif [[ "$SHELL" == *"bash"* ]]; then
  SHELL_CONFIG="$HOME/.bashrc"
fi

if [[ ":$PATH:" != *":$INSTALL_DIR:"* ]]; then
  log_warn "$INSTALL_DIR não está no PATH"

  if [ -n "$SHELL_CONFIG" ]; then
    if ! grep -q "$INSTALL_DIR" "$SHELL_CONFIG"; then
      echo "export PATH=\"$INSTALL_DIR:\$PATH\"" >> "$SHELL_CONFIG"
      log_info "PATH adicionado em $SHELL_CONFIG"
    else
      log_info "PATH já existe em $SHELL_CONFIG"
    fi

    log_warn "Reinicie o terminal ou execute: source $SHELL_CONFIG"
  else
    log_warn "Shell não identificado. Configure manualmente o PATH"
  fi
else
  log_info "PATH já configurado"
fi

echo
log_highlight "Crabe instalado com sucesso"
log_info "Teste com: crabe version"

