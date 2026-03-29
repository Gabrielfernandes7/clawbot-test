package main

import (
	"github.com/Gabrielfernandes7/crabe/internal/ui"
)

func main() {
	ui.Init()

	ui.Title("Crabe CLI - Inicializando agente")

	ui.Section("Status do sistema")
	ui.Success("Docker está rodando")
	ui.Info("Modelo atual: qwen2.5-coder:7b")
	ui.Highlight("Contexto do projeto: /home/gabriel/meu-projeto")

	ui.Section("Ações concluídas")
	ui.Success("OpenClaw configurado com sucesso")
	ui.Warning("RAM disponível: 3.2 GB (pode ficar apertado com 14b)")
}