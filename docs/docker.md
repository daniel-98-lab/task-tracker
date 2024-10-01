# Dockerfile:

Se utliza como archivo de configuración, su función es construir la imagen de Docker para tu aplicación.

## FROM golang:1.20 AS builder

Imagen base que usaremos para compilar la aplicación

## WORKDIR /app

Establece el directorio de trabajo dentro del conteendor en /app

## COPY . .

Copia todo el contenido del directorio actual (donde está el Dockerfile) en el directorio /app dentro del contenedor.

## RUN go mod init yask-tracker && go mod tidy

- RUN go mod init yask-tracker: Inicializa un nuevo módulo Go con el nombre yask-tracker, creando un archivo go.mod.

- go mod tidy: Limpia las dependencias, eliminando las que no se utilizan y descargando las que faltan, organizando el archivo go.mod.

## RUN CGO_ENABLED=0 GOOS=linux go build -o yask-tracker cmd/tracker/main.go

Este comando compila el código fuente de Go en un binario para Linux.

- CGO_ENABLED=0: Desactiva CGO, lo que genera un binario completamente estático (independiente de bibliotecas C).

- GOOS=linux: Indica que el binario debe estar construido para el sistema operativo Linux, independientemente del SO donde se ejecute la construcción.

- go build -o yask-tracker: Compila el archivo main.go de la ruta cmd/tracker/ y genera el binario llamado yask-tracker.

## FROM alpine:latest

Cambia la imagen base a Alpine Linux, que es una distribución de Linux muy ligera, lo que hace que la imagen final sea mucho más pequeña en tamaño.

## COPY --from=builder /app/yask-tracker /usr/local/bin/yask-tracker

Copia el binario yask-tracker que se generó en la primera etapa de construcción desde el directorio /app del contenedor builder a la carpeta /usr/local/bin/ en la nueva imagen de Alpine. De esta manera, el binario estará listo para ser ejecutado en la imagen final.

## ENTRYPOINT ["yask-tracker"]

Configura el binario yask-tracker como el punto de entrada de la imagen. Esto significa que cada vez que ejecutas el contenedor, se llamará automáticamente a yask-tracker con los argumentos que pases al contenedor.

# Uso de programa

## docker build -t yask-tracker .

- docker build: Este es el comando que le dice a Docker que cree (build) una nueva imagen.

- -t yask-tracker .: La opción -t asigna una etiqueta (tag) a la imagen. En este caso, estamos etiquetando la imagen con el nombre yask-tracker. Esto facilita identificarla cuando la necesites ejecutar más tarde.

## docker run -it --rm -v $(pwd)/data:/app/data yask-tracker

- docker run: Ejecuta un contenedor de Docker.

- -it: Permite que el contenedor se ejecute en modo interactivo.

- --rm: Este flag le indica a Docker que elimine el contenedor una vez que termine su ejecución. Esto es útil para evitar dejar contenedores inactivos en tu sistema.

-v $(pwd)/data:/app/data: Monta un volumen (volume mount). Necesario para actualizar el fichero tasks.json

- yask-tracker: Este es el nombre de la imagen que se ejecutará en el contenedor. Al estar definida como el punto de entrada (ENTRYPOINT), el contenedor ejecuta directamente tu aplicación Go.
