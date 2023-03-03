FROM golang:alpine

WORKDIR /app

COPY go.mod ./
COPY *.go ./
COPY ./bin/ ./bin/
COPY ./assets/index.html ./assets/index.html

RUN go build -o /ebs

EXPOSE 8081

CMD [ "/ebs" ]
