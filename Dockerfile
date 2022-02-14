FROM golang:latest AS build
COPY . /api
WORKDIR /api
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/main.go


FROM alpine:latest
EXPOSE 8000
RUN mkdir wasm-task
COPY --from=build /api/server /api/server
ENTRYPOINT ["./api/server"]