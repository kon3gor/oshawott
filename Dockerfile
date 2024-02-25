FROM registry.semaphoreci.com/golang:1.21 as builder

ENV APP_HOME /app

WORKDIR "$APP_HOME"
COPY . .

RUN go mod download
RUN go mod verify
RUN go run cmd/fetch/main.go
RUN go build -o oshawott cmd/main/main.go

FROM registry.semaphoreci.com/golang:1.21

ENV APP_HOME /app
ENV SA_FILE sa.json

RUN mkdir -p "$APP_HOME"
WORKDIR "$APP_HOME"

COPY --from=builder "$APP_HOME"/oshawott $APP_HOME
COPY --from=builder "$APP_HOME"/keys.txt $APP_HOME
COPY --from=builder "$APP_HOME"/sa.json $APP_HOME

CMD ["./oshawott"]
