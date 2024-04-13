FROM golang:1.22.2 as build

WORKDIR /app

COPY *.go ./
COPY *.mod ./
COPY *.sum ./
COPY lib ./lib
COPY static ./static
COPY views ./views
COPY routes ./routes

EXPOSE 8080
RUN CGO_ENABLED=0 go build .

FROM alpine as main
COPY --from=build /app /app

WORKDIR /app
ENTRYPOINT ["/app/note-server"]
