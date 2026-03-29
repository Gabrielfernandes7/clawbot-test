// internal/ui/style.go
package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	// Cores inspiradas no Claude Code (clean, moderno, suave)
	primary   = lipgloss.NewStyle().Foreground(lipgloss.Color("#00D1FF")) // azul ciano
	success   = lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF9D")) // verde neon suave
	warning   = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFCC00")) // amarelo
	error     = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF4D4D")) // vermelho
	info      = lipgloss.NewStyle().Foreground(lipgloss.Color("#888888")) // cinza

	// Estilos de caixas (igual Claude)
	titleBox = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#00D1FF")).
			Padding(0, 2).
			MarginBottom(1)

	header = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FFFFFF")).
		Background(lipgloss.Color("#00D1FF")).
		Padding(0, 1)

	// Emoji + texto
	crab = "🦀"
)

// Init deve ser chamado no início de todo comando
func Init() {
	// Força cores mesmo se o terminal não detectar
	lipgloss.SetColorProfile(lipgloss.ColorProfile())
}

// Title — usado no topo de cada comando (igual Claude)
func Title(text string) {
	fmt.Println(titleBox.Render(fmt.Sprintf("%s  %s", crab, text)))
}

// Success — mensagem de sucesso com check
func Success(msg string, args ...any) {
	fmt.Println(success.Render("✅ " + fmt.Sprintf(msg, args...)))
}

// Error — mensagem de erro
func Error(msg string, args ...any) {
	fmt.Println(error.Render("❌ " + fmt.Sprintf(msg, args...)))
}

// Info — informação neutra
func Info(msg string, args ...any) {
	fmt.Println(info.Render("ℹ️  " + fmt.Sprintf(msg, args...)))
}

// Highlight — destaque (usado em status, modelo atual, etc.)
func Highlight(msg string, args ...any) {
	fmt.Println(primary.Render("➜  " + fmt.Sprintf(msg, args...)))
}

// Warning
func Warning(msg string, args ...any) {
	fmt.Println(warning.Render("⚠️  " + fmt.Sprintf(msg, args...)))
}

// Section — separa seções (muito usado no Claude Code)
func Section(title string) {
	fmt.Println()
	fmt.Println(header.Render(" "+title+" "))
	fmt.Println(strings.Repeat("─", 50))
}