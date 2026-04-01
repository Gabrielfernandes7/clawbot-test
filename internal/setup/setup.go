package setup

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/Gabrielfernandes7/crabe/internal/ui"
	"github.com/spf13/cobra"
)

type SystemState struct {
	OpenClawInstalled bool
	OllamaRunning     bool
	Models            []string
}

func NewSetupCmd() *cobra.Command {
	var force bool
	var startNow bool
	var model string

	cmd := &cobra.Command{
		Use:   "setup",
		Short: "Instala e configura OpenClaw + Ollama automaticamente",
		Run: func(cmd *cobra.Command, args []string) {
			RunSetup(force, startNow, model)
		},
	}

	cmd.Flags().BoolVarP(&force, "force", "f", false, "Força reinstalação")
	cmd.Flags().BoolVarP(&startNow, "start", "s", false, "Inicia serviços automaticamente")
	cmd.Flags().StringVarP(&model, "model", "m", "", "Modelo preferido")

	return cmd
}

func RunSetup(force, autoStart bool, preferredModel string) {
	ui.Title("Crabe Setup")

	state := preflight()

	if force || !state.OpenClawInstalled {
		ui.Section("Instalação do OpenClaw")
		if err := installOpenClaw(); err != nil {
			ui.Error(fmt.Sprintf("Erro ao instalar OpenClaw: %v", err))
			return
		}
	} else {
		ui.Success("OpenClaw já instalado")
	}

	if !state.OllamaRunning {
		ui.Section("Ollama")
		ui.Error("Ollama não está rodando via Docker")
		ui.Info("Execute: docker compose up -d")
		return
	}

	models := state.Models

	if len(models) == 0 {
		ui.Warning("Nenhum modelo encontrado. Baixando modelo padrão...")
		if err := pullModel(preferredModel); err != nil {
			ui.Error("Falha ao baixar modelo")
			return
		}
		models = listOllamaModels()
	}

	model := chooseModel(models, preferredModel)

	ui.Section("Configuração")
	if err := configureOpenClaw(model, force); err != nil {
		ui.Error(fmt.Sprintf("Erro ao configurar OpenClaw: %v", err))
		return
	}

	ui.Section("Gateway")
	if err := configureGateway(); err != nil {
		ui.Error(fmt.Sprintf("Erro ao configurar gateway: %v", err))
		return
	}

	ui.Section("Validação final")
	verify()

	if autoStart {
		startServices()
	}

	ui.Title("Ambiente pronto")
	ui.Success(fmt.Sprintf("Modelo ativo: %s", model))
}

func preflight() SystemState {
	ui.Section("Preflight")

	state := SystemState{}

	if _, err := exec.LookPath("openclaw"); err == nil {
		ui.Success("OpenClaw encontrado")
		state.OpenClawInstalled = true
	} else {
		ui.Warning("OpenClaw não instalado")
	}

	state.OllamaRunning = checkOllamaRunning()
	state.Models = listOllamaModels()

	return state
}

func checkOllamaRunning() bool {
	out, err := exec.Command("docker", "ps", "--format", "{{.Names}} {{.Image}}").Output()
	if err != nil {
		return false
	}
	return strings.Contains(strings.ToLower(string(out)), "ollama")
}

func installOpenClaw() error {
	ui.Info("Instalando OpenClaw via install.sh...")

	cmd := exec.Command("bash", "-c", "curl -fsSL https://install.openclaw.ai | bash")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func listOllamaModels() []string {
	cmd := exec.Command("docker", "exec", "ollama", "ollama", "list")
	out, err := cmd.Output()

	if err != nil {
		out, err = exec.Command("ollama", "list").Output()
		if err != nil {
			return nil
		}
	}

	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	var models []string

	for i := 1; i < len(lines); i++ {
		fields := strings.Fields(lines[i])
		if len(fields) > 0 {
			models = append(models, fields[0])
		}
	}

	return models
}

func pullModel(preferred string) error {
	model := preferred
	if model == "" {
		model = "qwen2.5-coder:7b"
	}

	ui.Info(fmt.Sprintf("Baixando modelo %s...", model))

	cmd := exec.Command("docker", "exec", "ollama", "ollama", "pull", model)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func chooseModel(models []string, preferred string) string {
	if preferred != "" {
		for _, m := range models {
			if m == preferred {
				return preferred
			}
		}
	}
	return models[0]
}

func configureOpenClaw(model string, force bool) error {
	ui.Info("Configurando OpenClaw (modo não interativo)...")

	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	configDir := filepath.Join(home, ".openclaw")
	configPath := filepath.Join(configDir, "openclaw.json")
	backupPath := configPath + ".bak"

	_ = os.MkdirAll(configDir, 0755)

	// Backup se já existir
	if _, err := os.Stat(configPath); err == nil {
		if !force {
			ui.Warning("Configuração já existe. Use --force para sobrescrever.")
			return nil
		}

		ui.Info("Fazendo backup da configuração atual...")
		data, _ := os.ReadFile(configPath)
		_ = os.WriteFile(backupPath, data, 0644)
	}

	config := fmt.Sprintf(`{
	"models": {
		"default": "ollama/%s"
	},
	"providers": {
		"ollama": {
		"baseUrl": "http://localhost:11434"
		}
	}
	}`, model)

	err = os.WriteFile(configPath, []byte(config), 0644)
	if err != nil {
		return err
	}

	ui.Success("Configuração aplicada com sucesso")

	return nil
}

func configureGateway() error {
	ui.Info("Instalando serviço do gateway...")

	installCmd := exec.Command("openclaw", "gateway", "install")
	installCmd.Stdout = os.Stdout
	installCmd.Stderr = os.Stderr

	if err := installCmd.Run(); err != nil {
		ui.Warning("Falha ao instalar via systemd. Tentando modo fallback...")
	}

	ui.Info("Iniciando gateway...")

	startCmd := exec.Command("openclaw", "gateway", "start")
	startCmd.Stdout = os.Stdout
	startCmd.Stderr = os.Stderr

	if err := startCmd.Run(); err != nil {
		return fmt.Errorf("falha ao iniciar gateway: %w", err)
	}

	return waitGateway()
}

func waitGateway() error {
	ui.Info("Aguardando gateway subir na porta 18789...")

	for i := 0; i < 10; i++ {
		if isPortOpen("127.0.0.1:18789") {
			ui.Success("Gateway está ativo")
			return nil
		}
		time.Sleep(1 * time.Second)
	}

	return fmt.Errorf("gateway não respondeu na porta 18789")
}

func isPortOpen(addr string) bool {
	conn, err := net.DialTimeout("tcp", addr, 1*time.Second)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

func verify() {
	run("openclaw status")
	run("openclaw models list")
	run("ollama list")
}

func run(command string) {
	parts := strings.Split(command, " ")
	cmd := exec.Command(parts[0], parts[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}

func startServices() {
	ui.Info("Subindo Docker Compose...")

	cmd := exec.Command("docker", "compose", "up", "-d")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		ui.Error("Erro ao subir serviços")
	} else {
		ui.Success("Serviços iniciados")
	}
}