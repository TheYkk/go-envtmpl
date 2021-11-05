# builder image
FROM golang:1.17-alpine as builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . .
RUN make build

FROM gcr.io/distroless/static:nonroot
WORKDIR /app
COPY --from=builder /app/bin/envtmpl /app/envtmpl
ENTRYPOINT ["/app/envtmpl"]