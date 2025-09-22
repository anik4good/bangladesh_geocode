# Use multi-stage build
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY . .
RUN go build -o bangladesh_geocode

# Final minimal image
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/bangladesh_geocode .

RUN mkdir logs
# Set environment variables if needed
ENV GIN_MODE=release

EXPOSE 1552

CMD ["./bangladesh_geocode"]
