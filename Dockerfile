FROM golang:1.25-alpine
WORKDIR /usr/src/
COPY src .
RUN go mod download
RUN go build -o /usr/bin/app .
CMD ["app"]
