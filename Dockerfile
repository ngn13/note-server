FROM golang:1.21.4

WORKDIR /app

COPY *.go ./
COPY *.mod ./
COPY *.sum ./
COPY lib ./lib
COPY static ./static
COPY views ./views
COPY routes ./routes

RUN go build . 

ENTRYPOINT ["/app/note-server"]
