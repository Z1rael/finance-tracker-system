ARG NODE_VERSION=25-alpine
ARG NGINX_VERSION=alpine3.23
#
# DevContainer stage
#
FROM node:24-slim AS devel

RUN apt update && apt install -y git

WORKDIR /app

#COPY package.json package-lock.json* ./

RUN npm install -g @angular/cli

RUN npm install @types/geojson

COPY . .

#
# Build stage
#
FROM node:${NODE_VERSION} AS builder

WORKDIR /app

# copy package files first
COPY package.json *package-lock.json* ./

# install dependencies
RUN --mount=type=cache,target=/root/.npm npm ci

# copy the rest of the app code
COPY . .

# build angular application
RUN npm run build

#
# Release stage
#
FROM nginxinc/nginx-unprivileged:${NGINX_VERSION} AS live

# copy output from the build to default serving dir
COPY --chown=nginx:nginx --from=builder /app/dist/*/browser /usr/share/nginx/html

# use non root user
USER nginx

# default nginx container listens at 8080
EXPOSE 8080
