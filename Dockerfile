FROM golang:1.22.3 as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o cicd-hands-on

FROM scratch
COPY --from=builder /app/cicd-hands-on /app
ENTRYPOINT ["/app"]
