
FROM golang:1.23 AS build

ARG VERSION
ARG GIT_COMMIT_SHA
ARG GIT_REF
ARG GIT_BUILD

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN mkdir -p /data/sigils
RUN CGO_ENABLED=1 GOOS=linux go build \
 -ldflags " \
 -s -w \
 -X 'github.com/dusto/sigils/internal/version.Version=${VERSION}' \
 -X 'github.com/dusto/sigils/internal/version.GIT_COMMIT=${GIT_COMMIT_SHA}' \
 -X 'github.com/dusto/sigils/internal/version.GIT_BRANCH=${GIT_REF}' \
 -X 'github.com/dusto/sigils/internal/version.GIT_BUILD=${GIT_BUILD}' \
 " \
 -o /sigils

FROM gcr.io/distroless/base-debian12 AS release

WORKDIR /
USER nonroot:nonroot

COPY --from=build --chown=nonroot:nonroot /sigils /sigils
COPY --from=build --chown=nonroot:nonroot /data /data

# Default port if no params passed
EXPOSE 8888
# Default port for metrics
EXPOSE 9001

ENTRYPOINT ["/sigils"]
