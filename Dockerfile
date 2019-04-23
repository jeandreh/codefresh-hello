# Go Builder Image
FROM golang:1.11.2-alpine AS builder

RUN mkdir /app
WORKDIR /app

# copy sources
COPY ./app .

# Run tests, skip 'vendor'
RUN apk add --update git build-base
RUN go mod tidy
RUN go test -v $(go list ./...)

# Build application
RUN CGO_ENABLED=0 go build -v -o "myapp"

# Application Runtime Image
FROM alpine:3.9

# copy file from builder image
COPY --from=builder /app/myapp /usr/bin/myapp

CMD ["myapp", "--help"]