FROM golang:1.18

WORKDIR /app

COPY go.mod ./

RUN go mod download && go mod verify

COPY . .

RUN go build

CMD [ "go", "run",  "main.go" ]

