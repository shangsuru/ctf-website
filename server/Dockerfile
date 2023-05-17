# Build the application from source
FROM golang:1.19 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
COPY controllers/*.go ./controllers/
COPY db/*.go ./db/
COPY handlers/*.go ./handlers/
COPY models/*.go ./models/
COPY routes/*.go ./routes/
COPY .env ./.env

RUN CGO_ENABLED=0 GOOS=linux go build -o /main

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /main /main
COPY --from=build-stage /app/.env /.env

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/main"]