############################
# STEP 1 build executable Orchestrate binary
############################
FROM golang:1.15-buster AS builder

RUN apt-get update && \
	apt-get install --no-install-recommends -y \
	ca-certificates upx-ucl

RUN useradd appuser && mkdir /app
WORKDIR /app

# Use go mod with go 1.15
ENV GO111MODULE=on
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o /bin/main -a -tags netgo -ldflags '-w -s -extldflags "-static"' .
RUN upx /bin/main

############################
# STEP 2 build a small image
############################
FROM alpine:3.13

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd

COPY --from=builder /bin/main /go/bin/main
COPY --from=builder /app/TERMS_OF_SERVICE /

# Use an unprivileged user.
USER appuser
EXPOSE 8080
ENTRYPOINT ["/go/bin/main"]
