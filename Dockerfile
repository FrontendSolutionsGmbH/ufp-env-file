FROM golang:1.11.1   AS build-env

WORKDIR /go/src/ufp-env-file
COPY ./go/src/ufp-env-file .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build
RUN ls -ll

# final stage
FROM alpine
WORKDIR /
COPY --from=build-env /go/src/ufp-env-file /
ENTRYPOINT ./ufp-env-file
