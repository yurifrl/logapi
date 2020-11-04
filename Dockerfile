# Build the logapi binary
FROM golang:1.15 as builder

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

COPY pkg/ pkg/
COPY cmd/ cmd/
COPY logapi.go logapi.go

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -o logapi ./cmd/**.go

# Use distroless as minimal base image to package the logapi binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/logapi .
# Copy default config file
COPY config.yaml config.yaml
USER nonroot:nonroot

ENTRYPOINT ["/logapi"]
CMD ["server"]
