FROM golang:1.21-alpine@sha256:4bc6541af94a67d9dcabba9826c36e4e9497dacf4e8755ac503000b6ff75318f AS build

ENV CGO_ENABLED=0
COPY . /src
WORKDIR /src
RUN go build -o plugin-test-impl cmd/plugin-test-impl/main.go

FROM scratch
COPY --from=build /src/plugin-test-impl /plugin-test-impl
WORKDIR /
ENTRYPOINT ["/plugin-test-impl"]
CMD []