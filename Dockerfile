FROM golang:1.17.2-alpine3.14 as builder
ENV USERNAME n3dr
RUN adduser -D -g '' $USERNAME
COPY . /go/${USERNAME}/
WORKDIR /go/${USERNAME}/cmd/${USERNAME}
RUN apk add git=~2 && \
    CGO_ENABLED=0 go build && \
    cp n3dr /n3dr

FROM alpine:3.14.2
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /n3dr /usr/local/bin/n3dr
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
USER n3dr
ENTRYPOINT ["/usr/local/bin/n3dr"]
