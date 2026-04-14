package start

import (
	"github.com/Gabrielfernandes7/crabe/internal/setup"
	"github.com/Gabrielfernandes7/crabe/internal/ui"
	"github.com/spf13/cobra"
)

func NewStartCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Inicia o Ollama via docker-compose",
		Run: func(cmd *cobra.Command, args []string) {
			RunStart()
		},
	}
	return cmd
}

func RunStart() {
	ui.Title("Crabe Start - Iniciando Ollama")

	if err := setup.EnsureDockerUp(); err != nil {
		ui.Error("Falha ao iniciar o Ollama")
		return
	}

	ui.Success("Ollama está pronto para uso!")
	ui.Info("Agora execute: crabe install --model qwen2.5:3b")
}