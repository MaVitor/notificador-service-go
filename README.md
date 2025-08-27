# Notificador Service (Go)

Este é um microsserviço de notificações de alta performance escrito em Go, responsável por enviar mensagens através de diversas plataformas (atualmente, Telegram).

Este serviço faz parte de uma arquitetura maior de monitoramento de preços.

## Tecnologias Utilizadas

- **Linguagem:** [Go](https://go.dev/)
- **Roteamento HTTP:** Pacote `net/http` da biblioteca padrão.
- **Variáveis de Ambiente:** [godotenv](https://github.com/joho/godotenv)

## Configuração do Ambiente

Siga os passos abaixo para rodar este projeto localmente.

### Pré-requisitos

- **Go:** Versão 1.18 ou superior.
- **Credenciais do Telegram:**
    1. Um **Token de Bot** (obtido através do `@BotFather`).
    2. O **Chat ID** do destinatário da mensagem.

### Passos

1.  **Clone o repositório:**
    ```bash
    git clone https://github.com/MaVitor/notificador-service-go.git
    cd notificador-service-go
    git checkout develop # Se certifique de estar na branch develop
    ```

2.  **Crie o arquivo de ambiente:**
    Crie um arquivo chamado `.env` na raiz do projeto e adicione seu token do Telegram:
    ```
    TELEGRAM_BOT_TOKEN=seu_token_aqui
    ```

3.  **Instale as dependências:**
    O Go Modules cuidará disso automaticamente ao rodar o projeto.

4.  **Execute o projeto:**
    ```bash
    go run main.go
    ```
    O servidor estará disponível em `http://127.0.0.1:8081`.

## API Endpoints

A seguir estão os endpoints disponíveis na API.

### Health Check

Verifica a saúde e a disponibilidade do serviço.

-   **Método:** `GET`
-   **Path:** `/health`
-   **Resposta de Sucesso (200 OK):**
    ```json
    {
      "status": "ok"
    }
    ```

### Enviar Notificação

Envia uma mensagem de texto para um chat específico do Telegram.

-   **Método:** `POST`
-   **Path:** `/notificacao`
-   **Corpo da Requisição (JSON):**
    ```json
    {
      "chat_id": "ID_DO_CHAT_DE_DESTINO",
      "mensagem": "Sua *mensagem* aqui. O `ParseMode` Markdown está ativado."
    }
    ```
-   **Resposta de Sucesso (200 OK):**
    ```json
    {
      "status": "notificacao enviada"
    }
    ```