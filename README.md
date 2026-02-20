# Docker Build Demo

This project demonstrates different approaches to building minimal Docker images using Go applications. It includes examples of building images with:

- `scratch` base image
- `distroless` base image

## Overview

The goal of this project is to showcase how to create minimal, secure, and efficient Docker images for Go applications. Since Go compiles to a single static binary, it's ideal for use with minimal base images like `scratch` and `distroless`.

### Why Minimal Images?

- **Smaller size**: Faster downloads and less storage.
- **Security**: Reduced attack surface due to fewer packages and tools.
- **Efficiency**: Less bloat, better performance.

## Dockerfiles

### `Dockerfile.scratch`

Uses the `scratch` base image, which is literally empty. The final image contains only the compiled Go binary.

### `Dockerfile.distroless`

Uses Google's `distroless` base image, which includes only the runtime dependencies needed to run the application (e.g., CA certificates), without shell or package managers.

## Usage

### Build the Application

```bash
GOOS=linux go build -o main .
```

### Build Scratch Image

```bash
docker build -f Dockerfile.scratch -t demo-scratch .
```

### Build Distroless Image

```bash
docker build -f Dockerfile.distroless -t demo-distroless .
```

### Run Containers

```bash
docker run -p 8080:8080 demo-scratch
docker run -p 8080:8080 demo-distroless
```

## Project Files

- `main.go`: Simple HTTP server written in Go.
- `go.mod`: Go module definition.
- `Dockerfile.scratch`: Dockerfile using scratch base image.
- `Dockerfile.distroless`: Dockerfile using distroless base image.
- `README.md`: This file.

## References

- [GoogleContainer/distroless](https://github.com/GoogleContainerTools/distroless)
- [Docker Scratch](https://hub.docker.com/_/scratch)
- [Go Docker image best practices](https://docs.docker.com/language/golang/)
