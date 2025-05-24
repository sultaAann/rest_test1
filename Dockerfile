FROM golang:1.24.3

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /go_app

EXPOSE 8080

CMD [ "/go_app" ]