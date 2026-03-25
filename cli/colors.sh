#!/bin/bash

# Reset
RESET="\033[0m"

# Cores principais
RED="\033[0;31m"
GREEN="\033[0;32m"
YELLOW="\033[1;33m"
ORANGE="\033[38;5;208m"  # ANSI extended (fica bonito no terminal moderno)

# Estilos
BOLD="\033[1m"

# Funções utilitárias (melhor DX)
log_info() {
  echo -e "${GREEN}$1${RESET}"
}

log_warn() {
  echo -e "${YELLOW}$1${RESET}"
}

log_error() {
  echo -e "${RED}$1${RESET}"
}

log_highlight() {
  echo -e "${ORANGE}${BOLD}$1${RESET}"
}