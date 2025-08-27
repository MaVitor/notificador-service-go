# Build da aplicação
# Usa a imagem do Go pra compilar o binário (versão 1.24, compatível com o go.mod)
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copia os arquivos de dependência primeiro pra aproveitar o cache do Docker
COPY go.mod go.sum ./
RUN go mod download

# Copia o resto do código fonte
COPY . .

# Compila o app. CGO_ENABLED=0 pra garantir um binário estático
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/notificador-service .


# Imagem final de produção
# Parte de uma imagem Alpine limpa bem pequena
FROM alpine:latest

WORKDIR /app

# Copia apenas o binário compilado do estagio anterior
COPY --from=builder /app/notificador-service .

# Copia o .env pra dentro do container
COPY .env .

# Exponhe a porta que a aplicação vai usar
EXPOSE 8080

# Comando para iniciar o serviço quando o container subir
CMD ["./notificador-service"]
