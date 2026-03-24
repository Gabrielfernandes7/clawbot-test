# Crabe

Repositório para testar e usar **OpenClaw** 100% local com Ollama.  

O objetivo é ter um agente inteligente que você chama facilmente de **qualquer pasta** do seu computador.

## Como usar (fluxo recomendado)

1. **Suba o Ollama** (uma vez ou automaticamente):
    ```bash
    ./start-ollama.sh
    ```

<p align="center">
  <img src="./docs/clawbot-icon.png" width="100" height="100">
</p>

1. Executar o comando do docker-compose
    ```shell
    sudo docker compose up --build
    ```
2. Escolher o modelo de IA generativa:
    
    ```shell
    # Para puxar a imagem do projeto
    sudo docker exec -it ollama ollama pull llama3.2:1b
    ```

    ```shell
    # Executar uma imagem do projeto
    sudo docker exec -it ollama ollama run llama3.2:1b
    ```