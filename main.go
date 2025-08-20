package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// NotificacaoRequest é a estrutura do JSON que esperamos receber no corpo da requisição.
type NotificacaoRequest struct {
	ChatID   string `json:"chat_id"`
	Mensagem string `json:"mensagem"`
}

// TelegramRequestBody é a estrutura do JSON que enviaremos para a API do Telegram.
type TelegramRequestBody struct {
	ChatID    string `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode"`
}

// healthCheckHandler verifica a saúde do serviço.
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

// notificacaoHandler trata as requisições para enviar notificações.
func notificacaoHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Apenas permitir o método POST
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	// 2. Decodificar o JSON do corpo da requisição
	var req NotificacaoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Requisição JSON inválida", http.StatusBadRequest)
		return
	}

	// 3. Validar se os campos não estão vazios
	if req.ChatID == "" || req.Mensagem == "" {
		http.Error(w, "Campos 'chat_id' e 'mensagem' são obrigatórios", http.StatusBadRequest)
		return
	}

	// 4. Enviar a mensagem para o Telegram
	if err := sendTelegramNotification(req.ChatID, req.Mensagem); err != nil {
		log.Printf("Erro ao enviar notificação para o Telegram: %v", err)
		http.Error(w, "Erro interno ao enviar notificação", http.StatusInternalServerError)
		return
	}

	// 5. Retornar sucesso
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "notificacao enviada"})
}

// sendTelegramNotification envia a mensagem usando a API do Telegram.
func sendTelegramNotification(chatID, message string) error {
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		return fmt.Errorf("variável de ambiente TELEGRAM_BOT_TOKEN não definida")
	}

	// Monta a URL da API
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)

	// Monta o corpo da requisição para o Telegram
	reqBody := TelegramRequestBody{
		ChatID:    chatID,
		Text:      message,
		ParseMode: "Markdown", // Permite usar negrito, itálico, etc.
	}

	// Converte o corpo para JSON
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("erro ao converter requisição para JSON: %w", err)
	}

	// Faz a chamada POST para a API do Telegram
	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return fmt.Errorf("erro ao fazer requisição para o Telegram: %w", err)
	}
	defer resp.Body.Close()

	// Verifica se a resposta não foi um sucesso (status code diferente de 2xx)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("telegram retornou status inesperado: %s", resp.Status)
	}

	return nil
}

func main() {
	// Carrega as variáveis de ambiente do arquivo .env
	if err := godotenv.Load(); err != nil {
		log.Println("Aviso: arquivo .env não encontrado.")
	}

	http.HandleFunc("/health", healthCheckHandler)
	http.HandleFunc("/notificacao", notificacaoHandler)

	port := ":8081"
	log.Printf("Servidor iniciado na porta %s", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %s\n", err)
	}
}