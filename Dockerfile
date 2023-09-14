FROM docker:latest

RUN apk add go

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
ARG GOBIN=/usr/local/bin
RUN go install

ENTRYPOINT ["drone-helper"]
