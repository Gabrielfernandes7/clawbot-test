# Crabe 🦀

<p align="center">
  <img src="./docs/clawbot-icon.png" width="125" height="125" alt="Crabe icon">
</p>

**CLI moderna em Go** para rodar **OpenClaw + Ollama + Docker** 100% local.

**Objetivo principal:**  
Entre em **qualquer pasta** do seu computador e rode `crabe init`.  
Pronto. Você terá um agente inteligente trabalhando exatamente no contexto daquele projeto (lendo arquivos, entendendo o código, sugerindo melhorias, editando, criando testes, etc.).

---

## Como instalar (uma única vez)

### Opção recomendada (com Makefile)

```bash
git clone https://github.com/Gabrielfernandes7/crabe.git
cd crabe

# Compila e instala o binário Go
make install
```

Isso compila o projeto em Go e cria o comando `crabe` em `~/.local/bin/crabe`.

### Alternativa manual

```bash
cd crabe
go build -o crabe ./cmd/crabe
sudo cp crabe /usr/local/bin/   # ou cp para ~/.local/bin/
chmod +x /usr/local/bin/crabe
```

---

## Como usar (Fluxo diário)

```bash
# 1. Entre na pasta do projeto que deseja trabalhar
cd ~/Documentos/meu-projeto
# ou qualquer outra pasta do seu computador

# 2. Inicialize o agente
crabe init
```

Após o `init`, o OpenClaw + Ollama estarão configurados no contexto da pasta atual.

Você pode interagir com o agente usando linguagem natural:
- "entenda este projeto"
- "liste os arquivos desta pasta"
- "sugira melhorias no código"
- "crie um teste para a função X"
- "qual modelo você está usando?"

---

## Comandos disponíveis

| Comando              | Descrição                                      |
|----------------------|------------------------------------------------|
| `crabe init`         | Inicializa o agente no projeto atual           |
| `crabe init --force` | Força reinicialização                          |
| `crabe doctor`       | Diagnóstico do sistema (Docker, Ollama, etc.)  |
| `crabe status`       | Mostra status dos serviços e modelo atual      |
| `crabe --help`       | Lista todos os comandos e opções               |

---

## Modelos recomendados (para código)

- **`qwen2.5-coder:7b`** → **Recomendado** (melhor equilíbrio)
- **`qwen2.5-coder:14b`** → Mais inteligente (exige mais RAM)

**Como mudar o modelo padrão:**  
Em breve será possível com `crabe model set <nome>`. Por enquanto, a configuração está sendo migrada para arquivo em `~/.crabe/config.json`.

---

## Desenvolvimento (para quem contribui)

```bash
make build          # compila o binário
make install        # compila e instala
make doctor         # roda crabe doctor
make init           # roda crabe init
make clean          # limpa binários
```

Veja o `Makefile` para mais comandos.

---

## Estrutura atual do projeto (durante migração)

```
crabe/
├── cmd/crabe/          # Entry point da CLI em Go
├── internal/           # Código interno (ui, doctor, initcmd, etc.)
├── scripts/            # Scripts auxiliares (setup mínimo)
├── docker/             # docker-compose.yml
├── docs/
├── go.mod
├── Makefile
└── README.md
```

Scripts antigos em `cli/` e alguns em `scripts/` estão sendo gradualmente substituídos por código Go.

---

## Dicas importantes

- Não use `sudo` nos comandos normais.
- O binário é um **single binary** em Go → mais rápido, seguro e fácil de distribuir.
- Interface com cores e estilo moderno (usando Lipgloss).
- Após mudanças no código Go, rode `make install` novamente.

---

## Troubleshooting

- **"Comando 'crabe' não encontrado"**  
  → Rode `make install` novamente e reinicie o terminal.

- **Erro ao instalar**  
  → Rode `make remove-old` e depois `make install`.

- **Problemas com Docker**  
  → Certifique-se de que seu usuário está no grupo `docker` (`newgrp docker` ou logout/login).

- **Gateway ou Ollama não inicia**  
  → Rode `crabe doctor` para diagnosticar.

---

**Pronto!**

Agora é só entrar na pasta do seu projeto e digitar:

```bash
crabe init
```

---

## Bibliotecas utilizadas

- [spf13/cobra](https://github.com/spf13/cobra) — Framework de CLI
- [charmbracelet/lipgloss](https://github.com/charmbracelet/lipgloss) — Estilização bonita do terminal