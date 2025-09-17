FROM golang:1.25.1 as builder

WORKDIR /app
COPY . .
RUN go build -o hello-slsa

FROM gcr.io/distroless/base-debian12
COPY --from=builder /app/hello-slsa /hello-slsa
CMD ["/hello-slsa"]