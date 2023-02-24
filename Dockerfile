FROM golang:1.19-alpine as builder
WORKDIR /app
RUN mkdir -p /build
COPY . .
RUN go install; \
    go build -o /build/gurl
RUN chmod +x /build/gurl
COPY ./web/templates /build/web/templates/
COPY ./web/static /build/web/static/
ADD https://raw.githubusercontent.com/litvinav/vscript/main/vscript.js /build/web/static/vscript.js


FROM alpine
WORKDIR /app
COPY --from=builder /build ./
ENTRYPOINT ["./gurl"]