FROM golang:1.23 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=1 GOOS=linux go build -o /sigils

FROM gcr.io/distroless/base-debian12 AS release

WORKDIR /

COPY --from=build /sigils /sigils

# Default port if no params passed
EXPOSE 8888
# Default port for metrics
EXPOSE 9001

USER nonroot:nonroot

ENTRYPOINT ["/sigils"]
