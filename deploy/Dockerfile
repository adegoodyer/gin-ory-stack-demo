FROM golang:1.22-alpine as build

RUN apk add --update ca-certificates tzdata

WORKDIR /app
COPY . .
RUN go mod download
RUN GOOS=linux CGO_ENABLED=0 go build -o /app/bin/gin-ory-stack-demo ./cmd/server-side-browser/

FROM scratch

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /app/bin/gin-ory-stack-demo /
COPY --from=build /app/config /config
COPY --from=build /app/web /web

WORKDIR /

CMD ["/gin-ory-stack-demo"]