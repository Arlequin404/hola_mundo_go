# Usa una imagen base de Go
FROM golang:1.16

# Crea un directorio de trabajo
WORKDIR /app

# Copia los archivos de la aplicación
COPY . .

# Compila la aplicación
RUN go build -o main main.go

# Ejecuta la aplicación
CMD ["./main"]
