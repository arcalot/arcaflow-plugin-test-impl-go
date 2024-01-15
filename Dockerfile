FROM golang:1.21-alpine@sha256:fd78f2fb1e49bcf343079bbbb851c936a18fc694df993cbddaa24ace0cc724c5 AS build

ENV CGO_ENABLED=0
COPY . /src
WORKDIR /src
RUN go build -o plugin-test-impl cmd/plugin-test-impl/main.go

FROM scratch
COPY --from=build /src/plugin-test-impl /plugin-test-impl
WORKDIR /
ENTRYPOINT ["/plugin-test-impl"]
CMD []