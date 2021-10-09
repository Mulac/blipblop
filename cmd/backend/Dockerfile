################################
# STEP 1 build executable binary
################################
FROM golang:1.17 AS builder

WORKDIR /build

COPY . .

RUN go env

RUN go build -o /main cmd/backend/main.go

############################
# STEP 2 build a small image
############################
FROM golang:1.17
EXPOSE 8080
COPY --from=builder /main /main
COPY --from=builder /build/cmd/prototype.db /prototype.db
ENTRYPOINT ["/main"]