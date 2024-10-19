# Build application from source
FROM golang:1.22.0 AS build-stage
    WORKDIR /app

    COPY go.mod go.sum ./
    RUN go mod download

    COPY . .

    RUN CGO_ENABLED=0 GOOS=linux go build -o /api ./cmd/main.go

# Deploy application binary into a lean image
FROM scratch AS build-realease-stage
    WORKDIR /

    COPY --from=build-stage /api /api

    EXPOSE 8080

    ENTRYPOINT [ "/api" ]