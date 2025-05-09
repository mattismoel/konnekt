FROM golang:1.24-alpine AS build

WORKDIR /app

RUN apk add --no-cache gcc musl-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o /konnekt-backend ./cmd/http/main.go

FROM alpine

ENV TZ=Europe/Copenhagen

WORKDIR /app

COPY --from=build /konnekt-backend .

EXPOSE 8080

CMD ["./konnekt-backend", "-frontendURL", "http://frontend:3000"]
