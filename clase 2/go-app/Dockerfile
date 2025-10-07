
# Stage 1: Build

FROM golang:1.22-alpine AS builder

# Instalar dependencias necesarias para compilar
RUN apk add --no-cache git

# Crear directorio de trabajo
WORKDIR /app

# Copiar los archivos del módulo Go
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código fuente
COPY . .

# Compilar el binario de manera estática
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Etapa 2: Production

FROM alpine:3.20

# Crear directorio para el binario
WORKDIR /root/

# Copiar el binario desde la etapa anterior
COPY --from=builder /app/main .

# Exponer el puerto donde corre la app
EXPOSE 8089

# Comando de ejecución
CMD ["./main"]
