FROM golang:1.20-alpine

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN cd ./cmd && go build -o /server-meet-app

EXPOSE 8000

ENTRYPOINT [ "/server-meet-app" ]