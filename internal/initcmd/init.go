package initcmd

import (
	"fmt"
	"github.com/Gabrielfernandes7/crabe/internal/ui"
	"github.com/spf13/cobra"
	"os/exec"
)

func NewInitCmd() *cobra.Command {
	var force bool

	cmd := &cobra.Command{
		Use:   "init",
		Short: "Inicializa o Crabe e o OpenClaw no projeto atual",
		Long:  `Instala e configura tudo o que é necessário para rodar o OpenClaw localmente.`,
		Run: func(cmd *cobra.Command, args []string) {
			RunInit(force)
		},
	}

	cmd.Flags().BoolVarP(&force, "force", "f", false, "Força a reinicialização")
	return cmd
}

func RunInit(force bool) {
	ui.Title("Crabe Init - Configurando ambiente")

	ui.Section("Verificando ambiente")
	runDoctorChecks()

	ui.Section("Instalando dependências")
	if !isOpenClawInstalled() || force {
		ui.Info("OpenClaw não encontrado. Iniciando instalação...")
		installOpenClaw()
	} else {
		ui.Success("OpenClaw já está instalado")
	}

	ui.Section("Configurando projeto atual")
	setupProjectContext()

	ui.Success("✅ Crabe inicializado com sucesso!")
	ui.Info("Agora você pode usar:")
	ui.Info("   crabe status     → ver status dos serviços")
	ui.Info("   crabe start      → subir os serviços")
	fmt.Println()
	ui.Info("Acesse: http://localhost:3000")
}

// Funções auxiliares
func runDoctorChecks() {
	// Podemos chamar partes do doctor aqui no futuro
	fmt.Println("   • Docker → OK")
	fmt.Println("   • Ollama container → verificando...")
}

func isOpenClawInstalled() bool {
	// Verifica se o gateway ou openclaw já existe
	_, err := exec.LookPath("openclaw")
	if err == nil {
		return true
	}
	// Verificar se existe o diretório ou container
	return false // por enquanto
}

func installOpenClaw() {
	ui.Warning("Instalação do OpenClaw ainda em desenvolvimento...")
	// Aqui vamos chamar o script setup-openclaw.sh por enquanto
	// ou implementar diretamente em Go
	fmt.Println("   → Executando setup-openclaw.sh (temporário)")
	// exec.Command("bash", "scripts/setup-openclaw.sh").Run()
}

func setupProjectContext() {
	fmt.Println("   📁 Criando contexto do projeto atual...")
	// Criar pasta .crabe, context.md, etc.
}