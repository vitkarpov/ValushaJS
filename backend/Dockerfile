FROM golang:1.14.6-alpine3.12 as builder
COPY go.mod go.sum .env /go/src/github.com/vitkarpov/valushajs/backend/
WORKDIR /go/src/github.com/vitkarpov/valushajs/backend
RUN go mod download
COPY . /go/src/github.com/vitkarpov/valushajs/backend
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/api github.com/vitkarpov/valushajs/backend

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/vitkarpov/valushajs/backend/build/api /usr/bin/api
EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/api"]