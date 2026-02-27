# Clawbot - Test

<center>
    <img 
        src="./docs/clawbot-icon.png"
        height="100"
        width="100"
        style="border-radius: 20px"
    />
</center>

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