# Crabe 🦀

<p align="center">
  <img src="./docs/clawbot-icon.png" width="125" height="125">
</p>

Repositório de testes para rodar **OpenClaw** 100% local usando Ollama + Docker.

**Objetivo principal:**  
Entrar em **qualquer pasta** do seu computador e digitar `crabe init` para ter um agente inteligente trabalhando exatamente naquela pasta (lendo arquivos, entendendo o projeto, sugerindo melhorias, editando código, etc.).

---

## Como usar (Fluxo Recomendado)

### 1. Primeira vez (configuração inicial)

```bash
chmod +x /scripts/setup.sh
```

### 2. Uso diário (o mais importante)

```bash
# 1. Entre na pasta do projeto que você quer trabalhar
cd ~/Documentos/meu-projeto
# ou qualquer outra pasta:
# cd /caminho/para/qualquer/projeto

# 2. Inicie o agente
crabe init
```

**Alternativa rápida** (se preferir um único comando):

```bash
~/Documentos/clawbot-test/start-crabe.sh
```

---

## Comandos principais do Crabe

- `crabe init` → **Inicia o agente no contexto da pasta atual** (comando recomendado)
- `crabe` → Inicia o agente (sem inicialização explícita)
- `crabe status` → Mostra status do agente e modelo atual

### Comandos úteis dentro do Crabe

- `status`
- `qual modelo você está usando?`
- `liste os arquivos desta pasta`
- `entenda este projeto e me diga o que ele faz`
- `sugira melhorias no código`
- `crie um teste para a função X`
- `analise o README.md`

---

## Modelos recomendados (SLM para código)

- `qwen2.5-coder:7b` → **Recomendado** (melhor equilíbrio qualidade/velocidade/RAM)
- `qwen2.5-coder:14b` → Mais inteligente (usa mais RAM)
- `glm-4.7-flash` → Já está baixado (uso atual)

**Como trocar o modelo padrão:**  
Edite o arquivo `~/.local/bin/crabe` e altere a linha `--model "ollama/glm-4.7-flash"`.

---

## Estrutura final

```md
crabe/
├── cli/
│   └── crabe.sh
├── scripts/
│   ├── setup.sh
│   ├── start.sh
│   ├── doctor.sh
│   └── stop.sh
├── docker/
│   └── docker-compose.yml
├── config/
│   └── crabe.config.json
├── core/
│   └── context-resolver.sh
├── docs/
└── README.md
```

---

## Dicas importantes

- **Nunca rode os scripts de setup com `sudo`** (exceto `./fix-crabe.sh` uma única vez).
- Após rodar `./fix-crabe.sh`, feche e abra o terminal se o comando `crabe` ainda não for reconhecido.
- O Gateway roda em background. Para parar: `pkill -f "openclaw gateway"`
- Logs do gateway: `tail -f ~/.openclaw/gateway.log`

---

## Troubleshooting

- **"Comando 'crabe' não encontrado"**  
  → Rode `./fix-crabe.sh` e depois feche/abra o terminal.

- **Erro de permissão no Docker**  
  → Rode `newgrp docker` ou adicione seu usuário ao grupo docker.

- **Gateway não inicia**  
  → Rode `pkill -f "openclaw gateway"` e depois `crabe init` novamente.

- **Quer trocar de modelo**  
  → Edite `~/.local/bin/crabe` e mude o nome do modelo (ex: `qwen2.5-coder:7b`).

---

**Pronto!**  
Agora basta entrar na pasta do seu projeto e digitar:

```bash
crabe init
```

Quer que eu adicione alguma outra seção (ex: como usar com outros modelos, atalhos no shell, etc.)?
```

Essa versão está limpa, prática e reflete exatamente o que você quer: **crabe init** como comando principal.

Se quiser, posso deixar ainda mais curta ou adicionar badges no topo.  
O que acha? Quer alguma mudança?