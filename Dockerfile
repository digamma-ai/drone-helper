FROM docker:latest

RUN apk add go

WORKDIR /app

COPY *.go ./

RUN go mod init codeminders.com/drone-helper
RUN go mod tidy

ARG GOBIN=/usr/local/bin
RUN go install

ENTRYPOINT ["drone-helper"]
