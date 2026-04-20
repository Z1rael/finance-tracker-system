#
# Development stage
#
FROM golang:1.26-alpine AS development

WORKDIR /app

# Install dev tools
RUN apk add --no-cache git \
    bash \
    curl \
    build-base

# Install Go tools
RUN go install github.com/cosmtrek/air@latest

# Cache dependecies
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

CMD ["air"]

#
# Builder stage
#
FROM golang:1.26-alpine AS builder

WORKDIR /app

# Install git (needed for modules sometimes)
# RUN apk add --no-cache git

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o app ./cmd

#
# Produciton stage
#
FROM gcr.io/distroless/base-debian12 AS production

WORKDIR /app

COPY --from=builder /app/app .

USER nonroot:nonroot

EXPOSE 8080

ENTRYPOINT [ "/app/app" ]