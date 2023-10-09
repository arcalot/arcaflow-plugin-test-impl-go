FROM golang:1.21-alpine@sha256:a76f153cff6a59112777c071b0cde1b6e4691ddc7f172be424228da1bfb7bbda AS build

ENV CGO_ENABLED=0
COPY . /src
WORKDIR /src
RUN go build -o plugin-test-impl cmd/plugin-test-impl/main.go

FROM scratch
COPY --from=build /src/plugin-test-impl /plugin-test-impl
WORKDIR /
ENTRYPOINT ["/plugin-test-impl"]
CMD []