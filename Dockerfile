# Etapa 1: Construção
FROM golang:1.23.2 AS builder

# Defina o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copie o módulo Go e arquivos go.mod e go.sum
COPY go.mod go.sum ./

# Baixe as dependências do módulo Go
RUN go mod download

# Copie o restante do código fonte
COPY . .

# Compile o binário da aplicação
RUN go build -o api-go main.go

# Verifique se o binário foi criado com sucesso
RUN ls -l api-go

# Etapa 2: Execução
FROM alpine:latest

# Instale as dependências necessárias
RUN apk --no-cache add ca-certificates libc6-compat

# Defina o diretório de trabalho dentro do contêiner
WORKDIR /root/

# Copie o binário da aplicação do estágio de construção
COPY --from=builder /app/api-go .

# Verifique se o binário foi copiado corretamente
RUN ls -l api-go

# Ajuste as permissões do binário
RUN chmod +x api-go

# Exponha a porta que a aplicação vai usar
EXPOSE 8080

# Comando para executar o binário
CMD ["./api-go"]
