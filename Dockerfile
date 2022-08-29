FROM golang:1.19-alpine as dev

WORKDIR /work

FROM golang:1.19-alpine as build

WORKDIR /validate-csv
COPY ./ /validate-csv/
RUN go mod init github.com/robinlieb/validate-csv
RUN go build -o validate-csv ./cmd/

FROM alpine as runtime
COPY --from=build /validate-csv/validate-csv /usr/local/bin/validate-csv
COPY run.sh /
RUN chmod +x /run.sh
ENTRYPOINT [ "./run.sh" ]