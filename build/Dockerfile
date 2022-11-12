FROM golang:1.19-alpine3.16

WORKDIR /app

COPY  ./go.mod .
COPY  ./go.sum .
COPY  ./routeFile.yaml .
RUN go mod download

COPY  ./src/ ./src/

RUN go build -o api ./src/main.go 

RUN addgroup -S gouser \
  && adduser -S gouser -G gouser \
  && chown -R gouser:gouser /app
USER gouser

EXPOSE 4001

ENTRYPOINT [ "./api" ]
