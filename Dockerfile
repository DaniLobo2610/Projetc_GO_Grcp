# Etapa de Construcción
FROM golang:1.23.3 AS builder

WORKDIR /app

# Descargar dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código fuente y compilar
COPY . .
RUN go build -o server .

# Etapa de Ejecución
FROM debian:bookworm-slim

WORKDIR /app

# Instalar certificados necesarios
RUN apt-get update && apt-get install -y ca-certificates && update-ca-certificates


# Copiar el binario y otros archivos necesarios
COPY --from=builder /app/server .
COPY .env /app/.env

# Exponer puertos
EXPOSE 50051 8081

# Ejecutar el servidor
CMD ["./server"]
