package uninstall

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/Gabrielfernandes7/crabe/internal/ui"
	"github.com/spf13/cobra"
)

func NewUninstallCmd() *cobra.Command {
	var all bool
	var yes bool

	cmd := &cobra.Command{
		Use:   "uninstall",
		Short: "Remove completamente OpenClaw, Gateway, Ollama e configurações",
		Run: func(cmd *cobra.Command, args []string) {
			runUninstall(all, yes)
		},
	}

	cmd.Flags().BoolVarP(&all, "all", "a", false, "Remove também Ollama, volumes Docker e todos os dados")
	cmd.Flags().BoolVarP(&yes, "yes", "y", false, "Não pede confirmação (modo não-interativo)")

	return cmd
}

func runUninstall(all, yes bool) {
	ui.Title("🗑️  Crabe Uninstall - Removendo OpenClaw")

	if !yes {
		ui.Warning("⚠️  Esta ação é irreversível!")
		ui.Info("Vai remover OpenClaw, Gateway, ~/.openclaw e (se --all) Ollama.")
		ui.Info("Execute com --yes para confirmar.")
		return
	}

	// 1. Tentar uninstall oficial do OpenClaw primeiro
	ui.Section("Executando uninstall oficial do OpenClaw")
	if err := runOfficialUninstall(); err != nil {
		ui.Warning("OpenClaw CLI não respondeu ao uninstall oficial. Prosseguindo com limpeza manual...")
	}

	// 2. Limpeza manual agressiva
	ui.Section("Limpando pastas do OpenClaw")
	cleanOpenClawDirectories()

	// 3. Ollama (se --all)
	if all {
		ui.Section("Removendo Ollama + Docker")
		removeOllama()
	} else {
		ui.Success("Ollama mantido (use --all para remover também)")
	}

	// 4. Limpeza final do Crabe
	ui.Section("Limpando configurações do Crabe")
	cleanCrabeDirectories()

	ui.Title("✅ Desinstalação concluída!")
	ui.Success("OpenClaw foi removido o máximo possível.")
	ui.Info("Se ainda sobrar algo, rode manualmente:")
	ui.Info("  rm -rf ~/.openclaw")
	ui.Info("  npm uninstall -g openclaw   # se foi instalado via npm")
}

func runOfficialUninstall() error {
	cmd := exec.Command("openclaw", "uninstall", "--all", "--yes", "--non-interactive")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func cleanOpenClawDirectories() {
	home, _ := os.UserHomeDir()

	paths := []string{
		filepath.Join(home, ".openclaw"),
		filepath.Join(home, ".openclaw-dev"),
		"/usr/local/bin/openclaw",
	}

	for _, p := range paths {
		if err := os.RemoveAll(p); err == nil {
			ui.Success(fmt.Sprintf("Removido: %s", p))
		} else if !os.IsNotExist(err) {
			ui.Warning(fmt.Sprintf("Não conseguiu remover: %s", p))
		}
	}
}

func removeOllama() {
	ui.Info("Parando e removendo Ollama...")
	_ = exec.Command("docker", "compose", "down", "--volumes", "--remove-orphans").Run()
	_ = exec.Command("docker", "rm", "-f", "ollama").Run()
	_ = exec.Command("docker", "volume", "prune", "-f").Run()
	ui.Success("Ollama removido")
}

func cleanCrabeDirectories() {
	home, _ := os.UserHomeDir()
	_ = os.RemoveAll(filepath.Join(home, ".crabe"))
}