#!/bin/bash

set -e

CRABE_DIR="$HOME/.crabe"
INSTALL_DIR="$HOME/.local/bin"

case "$1" in
  basic)
    rm -f "$INSTALL_DIR/crabe"
    echo "Crabe removido (CLI)"
    ;;
  
  full)
    rm -f "$INSTALL_DIR/crabe"
    rm -rf "$CRABE_DIR"
    echo "Crabe removido (CLI + config)"
    ;;
  
  purge)
    rm -f "$INSTALL_DIR/crabe"
    rm -rf "$CRABE_DIR"

    sed -i '' '/.local\/bin/d' ~/.zshrc 2>/dev/null
    sed -i '' '/.local\/bin/d' ~/.zprofile 2>/dev/null
    sed -i '' '/.local\/bin/d' ~/.bashrc 2>/dev/null

    echo "Crabe removido completamente (incluindo PATH)"
    ;;
  
  *)
    echo "Uso:"
    echo "crabe uninstall basic   # remove CLI"
    echo "crabe uninstall full    # remove CLI + config"
    echo "crabe uninstall purge   # remove tudo + PATH"
    ;;
esac