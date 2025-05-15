# Build frontend
FROM node:lts-alpine AS frontend-builder
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable
WORKDIR /app
COPY frontend/ ./
# cache pnpm installs
RUN --mount=type=cache,id=pnpm,target=/pnpm/store pnpm install --frozen-lockfile
RUN pnpm run build

# Build backend
FROM golang:1.24-alpine AS backend-builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server

# Copy builds for final image
FROM alpine:latest
WORKDIR /app
COPY --from=frontend-builder /app/dist /app/frontend/dist
COPY --from=backend-builder /app/server ./
EXPOSE 8080
CMD [ "./server" ]