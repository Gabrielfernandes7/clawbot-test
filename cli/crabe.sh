#!/bin/bash

set -e

# Resolve caminho real mesmo via symlink
SCRIPT_PATH="$(readlink -f "$0")"
BASE_DIR="$(dirname "$SCRIPT_PATH")/.."

CRABE_DIR="$HOME/.crabe"
PROJECT_DIR="$(pwd)"

CORE="$BASE_DIR/core/context-resolver.sh"
START="$BASE_DIR/scripts/start.sh"
DOCTOR="$BASE_DIR/scripts/doctor.sh"

CONFIG_FILE="$CRABE_DIR/config.json"

# =========================
# Validações básicas
# =========================
function check_dependencies() {
  if ! command -v jq &> /dev/null; then
    echo "❌ jq não está instalado"
    echo "👉 Instale com: sudo apt install jq"
    exit 1
  fi

  if [ ! -f "$CORE" ]; then
    echo "❌ core não encontrado em: $CORE"
    exit 1
  fi
}

# =========================
# Configuração
# =========================
function load_config() {
  LOCAL_CONFIG="$PROJECT_DIR/model.crabe.json"
  GLOBAL_CONFIG="$CRABE_DIR/config.json"

  MODEL=""
  SOURCE=""

  # 1. Config local
  if [ -f "$LOCAL_CONFIG" ]; then
    MODEL=$(jq -r '.model // empty' "$LOCAL_CONFIG")
    SOURCE="local"
  fi

  # 2. Config global
  if [ -z "$MODEL" ] && [ -f "$GLOBAL_CONFIG" ]; then
    MODEL=$(jq -r '.model // empty' "$GLOBAL_CONFIG")
    SOURCE="global"
  fi

  # 3. Fallback
  if [ -z "$MODEL" ]; then
    MODEL="llama3.2:3b"
    SOURCE="default"
  fi
}

# =========================
# Importar módulos
# =========================
check_dependencies

source "$CORE"
source "$START"
source "$DOCTOR"

# =========================
# Entrada
# =========================
COMMAND=$1

if [ -z "$COMMAND" ]; then
  echo "Uso: crabe {init|status|doctor|model|version}"
  exit 1
fi

load_config

case $COMMAND in
  init)
    echo "🦞 Crabe iniciando..."

    crabe_start
    crabe_set_context "$PROJECT_DIR"

    echo ""
    echo "🧠 Modelo: $MODEL ($SOURCE)"
    echo "📂 Projeto: $PROJECT_DIR"
    echo "🔌 Gateway: ativo"
    echo ""
    echo "✅ Crabe pronto"
    ;;

  status)
    crabe_status
    ;;

  doctor)
    crabe_doctor
    ;;

  model)
    NEW_MODEL=$2

    if [ -z "$NEW_MODEL" ]; then
      echo "Modelo atual: $MODEL ($SOURCE)"
    else
      mkdir -p "$CRABE_DIR"
      echo "{ \"model\": \"$NEW_MODEL\" }" > "$CONFIG_FILE"
      echo "✅ Modelo alterado para: $NEW_MODEL"
    fi
    ;;

  version)
    echo "Crabe v0.1.0"
    echo "Base dir: $BASE_DIR"
    ;;

  *)
    echo "Uso: crabe {init|status|doctor|model|version}"
    ;;
esac