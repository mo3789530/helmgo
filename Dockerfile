FROM golang:1.17.7-alpine

ENV ROOT=/go/src/app
ENV CGO_ENABLED 0
WORKDIR ${ROOT}

RUN apk update && apk add git
COPY go.mod go.sum ./
COPY . ./
RUN go mod download
EXPOSE 8080

CMD ["go", "run", "main.go"]