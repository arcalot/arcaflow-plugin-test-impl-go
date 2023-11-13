FROM golang:1.21-alpine@sha256:110b07af87238fbdc5f1df52b00927cf58ce3de358eeeb1854f10a8b5e5e1411 AS build

ENV CGO_ENABLED=0
COPY . /src
WORKDIR /src
RUN go build -o plugin-test-impl cmd/plugin-test-impl/main.go

FROM scratch
COPY --from=build /src/plugin-test-impl /plugin-test-impl
WORKDIR /
ENTRYPOINT ["/plugin-test-impl"]
CMD []