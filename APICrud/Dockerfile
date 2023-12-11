# Usar la imagen oficial de Golang para construir la aplicación
FROM golang:latest AS builder

WORKDIR /app

# Copiar archivos necesarios para descargar las dependencias
COPY go.mod .
COPY go.sum .

# Descargar las dependencias
RUN go mod download

# Copiar el resto de los archivos de la aplicación
COPY . .

# Construir la aplicación
RUN go build -o main .

# Usar una imagen mínima de Alpine para ejecutar la aplicación
FROM alpine:latest

WORKDIR /app

# Copiar solo los archivos necesarios de la imagen de construcción
COPY go.mod .
COPY go.sum .

# Instalar las dependencias de PostgreSQL
RUN apk --no-cache add ca-certificates postgresql-client

# Establecer las variables de entorno para PostgreSQL
ENV DB_USER=postgres
ENV DB_PASSWORD=mi_contraseña
ENV DB_NAME=todolist
ENV DB_HOST=localhost
ENV DB_PORT=5432

# Exponer el puerto en el que escucha tu aplicación
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./main"]
