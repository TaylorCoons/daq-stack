# Build
FROM golang:1.18-bullseye as build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

ADD src/ /app/src

WORKDIR /app/src

RUN go build -o /daq-stack-api

CMD /app/src/daq-stack-api

# Deploy
FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=build /daq-stack-api /daq-stack-api

USER nonroot:nonroot

ENTRYPOINT [ "/daq-stack-api" ]
