# Use multi-stage build
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY . .
RUN go build -o bangladesh_geocode

# Final minimal images
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/bangladesh_geocode .
COPY .env .env

# Set environment variables if needed
ENV GIN_MODE=release

EXPOSE 1552

CMD ["./bangladesh_geocode"]
