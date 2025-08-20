package main

import (
    "encoding/json"
    "log"
    "net/http"
)

// healthCheckHandler verifica a saúde do serviço.
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
    // Define o header como JSON
    w.Header().Set("Content-Type", "application/json")

    // Prepara a resposta
    response := map[string]string{"status": "ok"}

    // Converte o mapa para JSON e escreve na resposta
    json.NewEncoder(w).Encode(response)
}

func main() {
    // Define a rota e o handler correspondente
    http.HandleFunc("/health", healthCheckHandler)

    port := ":8081"
    log.Printf("Servidor iniciado na porta %s", port)

    // Inicia o servidor HTTP
    if err := http.ListenAndServe(port, nil); err != nil {
        log.Fatalf("Erro ao iniciar o servidor: %s\n", err)
    }
}