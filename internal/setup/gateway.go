package setup

import (
	"net"
	"os"
	"os/exec"
	"time"

	"github.com/Gabrielfernandes7/crabe/internal/ui"
)

func SetupGateway() error {
	ui.Section("Gateway")

	exec.Command("openclaw", "gateway", "install").Run()

	cmd := exec.Command("openclaw", "gateway", "start")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	return waitGateway()
}

func waitGateway() error {
	ui.Info("Aguardando gateway...")

	for i := 0; i < 10; i++ {
		conn, err := net.DialTimeout("tcp", "127.0.0.1:18789", time.Second)
		if err == nil {
			conn.Close()
			ui.Success("Gateway ativo")
			return nil
		}
		time.Sleep(time.Second)
	}

	return nil
}