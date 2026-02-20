# Clawbot - Test

1. Executar o comando do docker-compose
    ```shell
    sudo docker compose up --build
    ```
2. Escolher o modelo de IA generativa:
    
    ```shell
    sudo docker exec -it ollama ollama pull llama3.2:1b
    ```

    ```shell
    # Opcionalmente teste
    sudo docker exec -it ollama ollama run llama3.2:1b
    ```