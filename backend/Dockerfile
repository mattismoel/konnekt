# ==============================================================================
# Build stage.
# 
# CGO dependencies are installed, Golang project dependencies are installed,
# and the project is built with CGO.
# ==============================================================================
FROM golang:1.24-alpine AS build

WORKDIR /app

RUN apk add --no-cache gcc musl-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o /konnekt-backend ./cmd/http/main.go


# ==============================================================================
# Production stage.
# ==============================================================================
FROM alpine
WORKDIR /app

ENV DB_DIR=${DB_DIR:-/app/data}
ENV ORIGIN=${ORIGIN:-http://localhost:4000}
ENV DB_FILE_NAME=${DB_FILE_NAME:-data.db}

COPY tables.sql seed.sql ./

# System upgrade before SQLite install.
RUN apk update && apk upgrade && apk add --no-cache sqlite

# Copy build executable from 'build' stage.
COPY --from=build /konnekt-backend .

ENV TZ=Europe/Copenhagen

# Make $DB_DIR directory for the database file.
RUN mkdir -p ${DB_DIR}

EXPOSE 8080

CMD ./konnekt-backend -origin=${ORIGIN} -host=0.0.0.0 -port=8080 -dbConnStr=${DB_DIR}/${DB_FILE_NAME}
