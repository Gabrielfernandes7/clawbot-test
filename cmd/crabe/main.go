// cmd/crabe/main.go
package main

import (
	"os"

	"github.com/Gabrielfernandes7/crabe/internal/doctor"
	"github.com/Gabrielfernandes7/crabe/internal/ui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "crabe",
	Short: "🦀 Crabe CLI - Agente de IA local com OpenClaw",
	Long:  `Ferramenta para facilitar o uso de OpenClaw + Ollama 100% local no contexto do seu projeto.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		ui.Init()
	},
}

func init() {
	rootCmd.AddCommand(doctor.NewDoctorCmd()) // doctor
	// Aqui vamos adicionando outros comandos no futuro: init, status, start, etc.
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		ui.Error("Erro ao executar comando: %v", err)
		os.Exit(1)
	}
}