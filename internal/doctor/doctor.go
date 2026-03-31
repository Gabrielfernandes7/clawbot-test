package doctor

import (
	"fmt"
	"os/exec"
	"strings"
	"github.com/Gabrielfernandes7/crabe/internal/ui"
)

func Run() {
	ui.Init()
	ui.Title("Crabe Doctor - Diagnóstico do ambiente")

	ui.Section("Docker")

	if isCommandAvailable("docker") {
		ui.Success("Docker está instalado")
	} else {
		ui.Error("Docker não foi encontrado ou a engine não está ligada")
		ui.Warning("Instale o Docker e certifique-se de que a engine está rodando")
	}

	ui.Section("Docker Compose")
	if checkDockerCompose() {
		ui.Success("Docker compose instalado")
	} else {
		ui.Error("Docker Compose não foi encontrado")
		ui.Warning("Instale o Docker Compose")
	}

	ui.Section("Ollama Container")
	if isOllamaContainerRunning() {
		ui.Success("Container Ollama está rodando")
	} else {
		ui.Info("Container Ollama não está rodando")
	}

	ui.Section("Porta 11434 (Ollama)")
	if isPortInUse(11434) {
		ui.Success("Porta 11434 está ativa (Ollama acessível)")
	} else {
		ui.Info("Porta 11434 está inativa")
	}
	
}

// Funções auxiliares
func isCommandAvailable(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

func checkDockerCompose() bool {
	out, err := exec.Command("docker", "compose", "version").CombinedOutput()
	if err != nil {
		return false
	}

	return strings.Contains(string(out), "Docker Compose")
}

func isOllamaContainerRunning() bool {
	// Tenta primeiro sem sudo (quando o usuário está no grupo docker)
	out, err := exec.Command("docker", "ps", "--format", "{{.Names}}|{{.Image}}").CombinedOutput()
	if err == nil {
		output := strings.ToLower(string(out))
		if strings.Contains(output, "ollama") || strings.Contains(output, "alpine/ollama") {
			return true
		}
	}

	// Se falhar, tenta com sudo (fallback)
	out, err = exec.Command("sudo", "docker", "ps", "--format", "{{.Names}}|{{.Image}}").CombinedOutput()
	if err != nil {
		return false
	}

	output := strings.ToLower(string(out))
	return strings.Contains(output, "ollama") || strings.Contains(output, "alpine/ollama")
}

func isPortInUse(port int) bool {
	// lsof não existe em todos os sistemas (ex: Windows), então usamos uma abordagem mais compatível
	cmd := exec.Command("sh", "-c", fmt.Sprintf("lsof -i:%d 2>/dev/null || ss -tuln 2>/dev/null | grep :%d || netstat -tuln 2>/dev/null | grep :%d", port, port, port))
	err := cmd.Run()
	return err == nil
}
