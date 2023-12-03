# Use the official Golang base image
FROM golang:latest

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /go/src/app

# Copia todos los archivos de tu paquete local al espacio de trabajo del contenedor
COPY . .

# Construye la aplicación Go (puedes ajustar los nombres según sea necesario)
RUN go build -o server server.go


# Expone el puerto en el que se ejecutará la aplicación (ajústalo según sea necesario)
EXPOSE 5050

# Comando para ejecutar tu aplicación, por ejemplo, el servidor
CMD ["./server"]
