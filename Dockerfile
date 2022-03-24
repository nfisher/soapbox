FROM golang:1.17-bullseye as build

WORKDIR /go/src/app

COPY go.* ./
RUN go mod download

COPY . ./
RUN --mount=type=cache,target=/root/.cache/go-build GIT_SHA=`git rev-parse --short HEAD` go build -mod=readonly -trimpath -v -ldflags "-X main.Version=$(GIT_SHA)"

FROM registry.access.redhat.com/ubi8/ubi-micro
COPY --from=build /go/src/app/soapbox /
CMD ["/soapbox"]