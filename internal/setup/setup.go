package setup

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/Gabrielfernandes7/crabe/internal/ui"
	"github.com/spf13/cobra"
)

func NewSetupCmd() *cobra.Command {
	var force bool
	var startNow bool

	cmd := &cobra.Command{
		Use:   "setup",
		Short: "Instala e configura o OpenClaw",
		Long:  `Clona o repositório do OpenClaw, configura diretórios e variáveis de ambiente.`,
		Run: func(cmd *cobra.Command, args []string) {
			RunSetup(force, startNow)
		},
	}

	cmd.Flags().BoolVarP(&force, "force", "f", false, "Força a reinstalação")
	cmd.Flags().BoolVarP(&startNow, "start", "s", false, "Inicia automaticamente após a instalação")
	return cmd
}

func RunSetup(force, autoStart bool) {
	ui.Title("🦀 Crabe Setup - Instalando OpenClaw")

	ui.Section("Verificando dependências")
	if err := checkDependencies(); err != nil {
		return
	}

	installDir := filepath.Join(os.Getenv("HOME"), ".openclaw")

	ui.Section("OpenClaw Repository")
	if err := cloneOrUpdateOpenClaw(installDir, force); err != nil {
		return
	}

	ui.Section("Configurando ambiente")
	if err := setupDirectoriesAndEnv(installDir); err != nil {
		return
	}

	ui.Success("✅ OpenClaw instalado e configurado com sucesso!")

	if autoStart || askToStart() {
		ui.Section("Iniciando serviços")
		startOpenClaw(installDir)
	} else {
		ui.Info("Você pode iniciar depois com: crabe start")
	}

	ui.Info(fmt.Sprintf("📂 OpenClaw instalado em: %s", installDir))
}

func checkDependencies() error {
	deps := []string{"git", "docker", "curl"}
	for _, dep := range deps {
		if _, err := exec.LookPath(dep); err != nil {
			ui.Error(fmt.Sprintf("Dependência não encontrada: %s", dep))
			ui.Info("Instale as dependências e tente novamente.")
			return fmt.Errorf("missing dependency")
		}
		ui.Success(fmt.Sprintf("%s OK", dep))
	}
	return nil
}

func cloneOrUpdateOpenClaw(installDir string, force bool) error {
	repoURL := "https://github.com/openclaw/openclaw.git"

	if _, err := os.Stat(installDir); err == nil {
		if !force {
			ui.Warning("OpenClaw já está instalado. Use --force para reinstalar.")
			return nil
		}
		ui.Info("Atualizando repositório existente...")
		return exec.Command("git", "-C", installDir, "pull").Run()
	}

	ui.Info("Clonando OpenClaw...")
	return exec.Command("git", "clone", repoURL, installDir).Run()
}

func setupDirectoriesAndEnv(installDir string) error {
	configDir := filepath.Join(installDir, "config") // ou $HOME/.openclaw/config
	workspaceDir := filepath.Join(installDir, "workspace")

	os.MkdirAll(configDir, 0755)
	os.MkdirAll(workspaceDir, 0755)

	ui.Success("Diretórios criados")

	envFile := filepath.Join(installDir, ".env")
	envExample := filepath.Join(installDir, ".env.example")

	// Copia .env.example se não existir
	if _, err := os.Stat(envFile); os.IsNotExist(err) && fileExists(envExample) {
		copyFile(envExample, envFile)
		ui.Success(".env criado a partir do .env.example")
	}

	// Adiciona / atualiza variáveis
	appendEnvVariables(envFile, configDir, workspaceDir)

	return nil
}

func appendEnvVariables(envFile, configDir, workspaceDir string) {
	content := fmt.Sprintf(`
# Configurado automaticamente pelo Crabe
OPENCLAW_CONFIG_DIR=%s
OPENCLAW_WORKSPACE_DIR=%s
`, configDir, workspaceDir)

	f, _ := os.OpenFile(envFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	f.WriteString(content)
	f.Close()

	ui.Success("Variáveis de ambiente configuradas")
}

func startOpenClaw(installDir string) {
	ui.Info("Subindo containers...")

	cmd := exec.Command("docker", "compose", "up", "-d", "--build")
	cmd.Dir = installDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		ui.Error("Falha ao subir os containers")
	} else {
		ui.Success("OpenClaw iniciado com sucesso!")
	}
}

func askToStart() bool {
	// Podemos melhorar isso depois com prompt interativo
	return false // por enquanto desativado, usamos a flag --start
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func copyFile(src, dst string) error {
	data, _ := os.ReadFile(src)
	return os.WriteFile(dst, data, 0644)
}