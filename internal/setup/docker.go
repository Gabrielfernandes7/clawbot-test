package setup

import (
	"os"
	"os/exec"

	"github.com/Gabrielfernandes7/crabe/internal/ui"
)

const composePath = "docker/docker-compose.yml"

func EnsureDockerUp() error {
	ui.Section("Docker")

	if _, err := os.Stat(composePath); err != nil {
		ui.Error("docker-compose.yml não encontrado")
		return err
	}

	cmd := exec.Command("docker", "compose", "-f", composePath, "up", "-d")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	ui.Info("Subindo containers...")
	if err := cmd.Run(); err != nil {
		return err
	}

	ui.Success("Containers ativos")
	return nil
}