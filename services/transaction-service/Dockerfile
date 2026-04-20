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

# Create non-root user
ARG USERNAME=devel
ARG UID=1000
ARG GID=1000

ENV GOPATH=/home/$USERNAME/go
ENV GOBIN=/home/$USERNAME/go/bin
ENV PATH=$PATH:$GOBIN

RUN addgroup -g $GID $USERNAME \
    && adduser -D -u $UID -G $USERNAME $USERNAME

RUN mkdir -p /home/$USERNAME/go \
    && chown -R $USERNAME:$USERNAME /home/$USERNAME/go

USER devel

# Dev util
RUN go install github.com/air-verse/air@latest

# Cache dependecies
COPY --chown=${USERNAME}:${GID} go.mod  ./
RUN go mod download

COPY --chown=${USERNAME}:${GID} . .

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