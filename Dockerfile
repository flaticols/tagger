FROM golang:1.19 AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /src
COPY . .
RUN go build -ldflags "-s -w" -o /bin/tagger .

FROM gcr.io/distroless/static-debian11
COPY --from=builder /bin/tagger /tagger
ENTRYPOINT ["/tagger"]
CMD ["create", "--actions"]
