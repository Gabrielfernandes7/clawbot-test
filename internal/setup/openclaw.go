package setup

import (
	"os"
	"os/exec"

	"github.com/Gabrielfernandes7/crabe/internal/ui"
)

func SetupOpenClawWithOllama(model string) error {
	ui.Section("OpenClaw (via Ollama)")

	cmd := exec.Command(
		"ollama",
		"launch",
		"openclaw",
		"--model", model,
		"--yes",
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}