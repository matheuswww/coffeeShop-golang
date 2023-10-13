FROM golang:1.20 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify && CGO_ENABLED=0 go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
COPY . .
RUN CGO_ENABLED=0 go build -o /app/server

FROM alpine
WORKDIR /app
EXPOSE 8080
COPY --from=build /go/bin/migrate /go/bin/
ENV PATH=$PATH:/go/bin
COPY --from=build /app/ .
CMD ["./server"]