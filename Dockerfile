FROM golang:1.19 AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN apt-get -qq update && \
    apt-get -yqq install upx

WORKDIR /src
COPY . .

RUN go build -ldflags "-s -w -extldflags '-static'" -o /bin/tagger . && strip /bin/app && upx -q -9 /bin/tagger
RUN echo "nobody:x:65534:65534:Nobody:/:" > /etc_passwd

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc_passwd /etc/passwd
COPY --from=builder --chown=65534:0 /bin/tagger /tagger

USER nobody
ENTRYPOINT ["/tagger"]
CMD ["do", "--actions"]