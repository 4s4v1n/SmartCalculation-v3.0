FROM gcc:latest as core_build


FROM golang:latest

WORKDIR /app

COPY ./ /app

RUN go mod download
RUN make lib
RUN cp /app/pkg/model/core/libcalculator.so /usr/lib/
RUN make build

EXPOSE 8080

CMD ["./calculator"]
