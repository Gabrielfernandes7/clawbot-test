#!/bin/bash

set -e

# Resolve caminho real mesmo via symlink
SCRIPT_PATH="$(realpath "$0" 2>/dev/null || greadlink -f "$0")"
BASE_DIR="$(dirname "$SCRIPT_PATH")/.."

CRABE_DIR="$HOME/.crabe"
PROJECT_DIR="$(pwd)"

# Paths dos módulos de comandos
CORE="$BASE_DIR/core/context-resolver.sh"
START="$BASE_DIR/scripts/start.sh"
DOCTOR="$BASE_DIR/scripts/doctor.sh"
SETUP_OPENCLAW="$BASE_DIR/scripts/setup-openclaw.sh"
START_OLLAMA="$BASE_DIR/scripts/start-ollama.sh"
COLORS="$BASE_DIR/cli/colors.sh"

CONFIG_FILE="$CRABE_DIR/config.json"

# Importar cores
source "$COLORS"

# Validações básicas
function check_dependencies() {
  if ! command -v jq &> /dev/null; then
    log_error "jq não está instalado"
    log_warn "Instale com: sudo apt install jq"
    exit 1
  fi

  if [ ! -f "$CORE" ]; then
    log_error "core não encontrado em: $CORE"
    exit 1
  fi
}

# Configuração
function load_config() {
  LOCAL_CONFIG="$PROJECT_DIR/model.crabe.json"
  GLOBAL_CONFIG="$CRABE_DIR/config.json"

  MODEL=""
  SOURCE=""

  if [ -f "$LOCAL_CONFIG" ]; then
    MODEL=$(jq -r '.model // empty' "$LOCAL_CONFIG")
    SOURCE="local"
  fi

  if [ -z "$MODEL" ] && [ -f "$GLOBAL_CONFIG" ]; then
    MODEL=$(jq -r '.model // empty' "$GLOBAL_CONFIG")
    SOURCE="global"
  fi

  if [ -z "$MODEL" ]; then
    MODEL="llama3.2:3b"
    SOURCE="default"
  fi
}

# Importar módulos
check_dependencies

source "$CORE"
source "$START"
source "$DOCTOR"

# Entrada
COMMAND=$1

if [ -z "$COMMAND" ]; then
  log_warn "Uso: crabe {init|status|doctor|model|version|install}"
  exit 1
fi

load_config

case $COMMAND in
  init)
    log_highlight "Crabe iniciando..."

    crabe_start
    crabe_set_context "$PROJECT_DIR"

    echo ""
    log_info "Modelo: $MODEL ($SOURCE)"
    log_info "Projeto: $PROJECT_DIR"
    log_info "Gateway: ativo"
    echo ""
    log_info "Crabe pronto"
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
      log_info "Modelo atual: $MODEL ($SOURCE)"
    else
      mkdir -p "$CRABE_DIR"
      echo "{ \"model\": \"$NEW_MODEL\" }" > "$CONFIG_FILE"
      log_info "Modelo alterado para: $NEW_MODEL"
    fi
    ;;

  install)
    TARGET=$2

    case $TARGET in
      openclaw)
        log_highlight "Instalando OpenClaw..."
        bash "$SETUP_OPENCLAW"
        ;;

      ollama)
        shift 2

        MODEL_ARG=""

        while [[ $# -gt 0 ]]; do
          case $1 in
            --model)
              MODEL_ARG="$2"
              shift 2
              ;;
            *)
              log_error "Parâmetro inválido: $1"
              exit 1
              ;;
          esac
        done

        log_highlight "Configurando Ollama..."

        if [ -n "$MODEL_ARG" ]; then
          bash "$START_OLLAMA" --model "$MODEL_ARG"
        else
          bash "$START_OLLAMA"
        fi
        ;;

      *)
        log_warn "Uso: crabe install {openclaw|ollama}"
        exit 1
        ;;
    esac
    ;;

  version)
    log_highlight "Crabe v0.1.0"
    echo "Base dir: $BASE_DIR"
    ;;

  uninstall)
    bash "$BASE_DIR/scripts/uninstall.sh" "$2"
    ;;

  *)
    log_warn "Uso: crabe {init|status|doctor|model|version|install}"
    ;;
esac
