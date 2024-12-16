FROM golang:1.23.4 as build

WORKDIR /app

COPY Makefile ./
COPY *.go     ./
COPY *.mod    ./
COPY *.sum    ./

COPY lib      ./lib
COPY static   ./static
COPY views    ./views
COPY routes   ./routes

RUN make CGO_ENABLED=0

FROM alpine as main
COPY --from=build /app /app

RUN apk update && apk add make

WORKDIR /app
RUN make install

EXPOSE 8080
ENTRYPOINT ["/app/note-server", "-interface", "0.0.0.0:8080", "-notes", "/notes"]
