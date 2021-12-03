FROM golang:1.17-buster AS build

WORKDIR /go/src/app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build -o /app/api-server

FROM gcr.io/distroless/base-debian10

WORKDIR /app

COPY --from=build /app/api-server /app/api-server
EXPOSE 8080
USER nonroot:nonroot

ENTRYPOINT ["/app/api-server"]