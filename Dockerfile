FROM golang:1.17 AS build

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build

FROM busybox:latest

COPY --from=build /app/lach /app/lach

ENTRYPOINT [ "/app/lach" ]
