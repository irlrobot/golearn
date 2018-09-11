FROM golang:alpine AS builder
LABEL maintainer="josh@userdel.com"

# install ssl ca certificates
RUN apk add --update ca-certificates

# create appuser
RUN adduser -D -g '' appuser

# copy source
COPY . $GOPATH/src/github.com/irlrobot/leadership/
WORKDIR $GOPATH/src/github.com/irlrobot/leadership/

# det dependancies
RUN go get -d -v

# build the binary
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags=”-w -s” -o /go/bin/leadership
RUN go build -o /go/bin/leadership

# final image
#FROM scratch
FROM alpine:latest

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
# copy our static executable
COPY --from=builder /go/bin/leadership /go/bin/leadership

USER appuser

EXPOSE 8080
ENTRYPOINT ["/go/bin/leadership"]