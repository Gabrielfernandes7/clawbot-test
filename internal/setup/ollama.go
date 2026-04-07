package setup

import (
	"os"
	"os/exec"
	"strings"

	"github.com/Gabrielfernandes7/crabe/internal/ui"
)

const defaultModel = "qwen2.5:3b"

func listModels() []string {
	cmd := exec.Command("docker", "exec", "ollama", "ollama", "list")
	out, err := cmd.Output()
	if err != nil {
		return nil
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

func EnsureModel(models []string) (string, error) {
	if len(models) > 0 {
		return models[0], nil
	}

	ui.Section("Modelo")
	ui.Info("Baixando modelo padrão: %s", defaultModel)

	cmd := exec.Command("docker", "exec", "ollama", "ollama", "pull", defaultModel)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return "", err
	}

	return defaultModel, nil
}