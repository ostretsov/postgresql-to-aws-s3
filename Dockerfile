FROM golang:1.13.5-alpine3.10
RUN apk update && apk add --no-cache git ca-certificates tzdata curl postgresql-client && update-ca-certificates

# Create appuser
RUN adduser -D -g '' appuser

WORKDIR $GOPATH/src/postgres_to_aws_s3
COPY . .

# Fetch dependencies.
RUN go mod download

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/postgres_to_aws_s3 .

COPY entrypoint.sh /bin

ENTRYPOINT ["/bin/entrypoint.sh"]
CMD ["/go/bin/postgres_to_aws_s3"]