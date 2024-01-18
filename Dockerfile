FROM golang:1.23

WORKDIR /usr/src/app

COPY . .

RUN make all

EXPOSE 8080

CMD ["./bare-go", "start"]
