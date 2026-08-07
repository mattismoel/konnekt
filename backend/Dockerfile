# ==============================================================================
# Build stage.
# 
# CGO dependencies are installed, Golang project dependencies are installed,
# and the project is built with CGO.
# ==============================================================================
FROM golang:1.25-alpine AS build

WORKDIR /app

RUN apk add --no-cache gcc musl-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -v -o /konnekt-backend ./cmd/http/...

# ==============================================================================
# Production stage.
# ==============================================================================
FROM alpine
WORKDIR /app

ENV DB_CONN_STRING=${DB_CONN_STRING:-/app/data/data.db}
ENV ORIGIN=${ORIGIN:-http://localhost:4000}

# System upgrade before SQLite install.
RUN apk update && apk upgrade && apk add --no-cache sqlite

# Copy build executable from 'build' stage.
COPY --from=build /konnekt-backend .

ENV TZ=Europe/Copenhagen

EXPOSE 8080

CMD ./konnekt-backend -origin=${ORIGIN} -host=0.0.0.0 -port=8080 -dbConnStr=${DB_CONN_STRING}
