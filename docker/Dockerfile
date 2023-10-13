FROM golang:1.21.3-bookworm as builder

ENV CGO_ENABLED 0

WORKDIR /go/src/github.com/k-akari/payment.com

COPY . .

RUN apt-get update && apt-get install --no-install-recommends -y make

RUN make build

FROM gcr.io/distroless/static-debian12:latest as api

COPY --from=builder --chown=nonroot:nonroot /go/src/github.com/k-akari/payment.com/dist/api /api

USER nonroot

CMD ["/api"]
