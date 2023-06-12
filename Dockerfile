FROM golang:1.18-alpine AS build

ENV CGO_ENABLED=0
COPY . /src
WORKDIR /src
RUN go build -o test-impl-plugin cmd/test-impl-plugin/main.go

FROM scratch
COPY --from=build /src/test-impl-plugin /test-impl-plugin
WORKDIR /
ENTRYPOINT ["/test-impl-plugin"]
CMD []