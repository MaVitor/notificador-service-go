# Notificador Service (Go)

Este é um microsserviço de notificações de alta performance escrito em Go, responsável por enviar mensagens através do Telegram.

Este serviço faz parte de uma arquitetura maior de monitoramento de preços.

## Tecnologias Utilizadas

-   **Linguagem:** Go
-   **Roteamento HTTP:** Pacote `net/http` da biblioteca padrão.
-   **Variáveis de Ambiente:** `godotenv`

## Como Executar (Com Docker Compose)

Este serviço não foi projetado para ser executado de forma isolada. Ele é orquestrado pelo Docker Compose a partir do repositório principal da plataforma.

1.  **Clone o repositório principal:**
    ```bash
    git clone [https://github.com/MaVitor/price-alert-platform.git](https://github.com/MaVitor/price-alert-platform.git)
    cd price-alert-platform
    ```

2.  **Suba o ambiente:**
    O `docker-compose.yml` no diretório raiz cuidará de construir a imagem deste serviço e iniciá-lo junto com os outros.
    ```bash
    docker-compose up
    ```
    A API estará disponível na porta `8080` do seu host.

## API Endpoints

### Health Check

-   **Método:** `GET`
-   **Path:** `/health`
-   **Resposta de Sucesso (200 OK):**
    ```json
    {
      "status": "ok"
    }
    ```

### Enviar Notificação

-   **Método:** `POST`
-   **Path:** `/notificacao`
-   **Corpo da Requisição (JSON):**
    ```json
    {
      "chat_id": "ID_DO_CHAT_DE_DESTINO_NO_TELEGRAM",
      "mensagem": "Sua mensagem aqui."
    }
    ```
-   **Resposta de Sucesso (200 OK):**
    ```json
    {
      "status": "notificacao enviada"
    }
    ```