FROM golang:1.22-alpine AS build
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go mod tidy
RUN go build -o /out/frogshort ./cmd/main.go

FROM alpine:3.20
RUN adduser -D -g "" appuser
USER appuser
WORKDIR /app
COPY --from=build  /out/frogshort /app/frogshort
EXPOSE 8080
ENV PORT=8080
ENTRYPOINT ["/app/frogshort"]
