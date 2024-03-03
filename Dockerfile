FROM golang:1.22-alpine3.19 AS builder

WORKDIR /app

RUN go clean --modcache

COPY . .

RUN go mod download

#RUN go get -u github.com/mentalisit/rsbot@master

RUN go build -mod=readonly -o bridge .

FROM alpine:3.19 AS runner

COPY --from=builder /app/bridge /usr/local/bin/
COPY --from=builder /app/config.yml /usr/local/bin/

WORKDIR /usr/local/bin/
CMD ["bridge"]
