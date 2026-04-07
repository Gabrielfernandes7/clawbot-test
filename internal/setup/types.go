package setup

type SystemState struct {
	OpenClawInstalled bool
	DockerAvailable   bool
	DockerRunning     bool
	OllamaRunning     bool
	Models            []string
}