FROM golang:latest
RUN apt-get update && apt-get install -y mingw-w64
WORKDIR /app
COPY . .
RUN GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build -o myapp.exe main.go
#-- test

