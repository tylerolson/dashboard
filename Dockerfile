# Build frontend
FROM node:22.15.0-alpine3.21 AS frontend-builder
# setup pnpm
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable
WORKDIR /app
# copy contents of frontend dir into build root
COPY frontend ./
# cache pnpm installs
RUN --mount=type=cache,id=pnpm,target=/pnpm/store pnpm install --frozen-lockfile
RUN pnpm run build

# Build backend
FROM golang:1.24.3-alpine3.21 AS backend-builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
#  copy backend into builder
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o server

# Copy builds for final image
FROM alpine:3.21
WORKDIR /app
COPY --from=frontend-builder /app/dist frontend/dist/
COPY --from=backend-builder /app/server ./
EXPOSE 8080
CMD [ "./server" ]