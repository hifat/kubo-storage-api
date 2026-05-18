# Kubo Storage API

## Overview

**Kubo Storage API** is a gRPC-based file storage service built with Go. It provides a secure interface for uploading, managing, and accessing files stored in Cloudflare R2 (S3-compatible object storage).

### Key Features
- **File Upload**: Upload files with automatic path normalization
- **Presigned URLs**: Generate temporary signed URLs for secure file access
- **File Deletion**: Remove files from storage
- **gRPC API**: High-performance RPC interface
- **Cloudflare R2**: S3-compatible cloud storage backend

---

## How to Run

### Prerequisites
- Go 1.25.5 or higher
- Docker & Docker Compose (optional)

### Quick Start

1. **Clone and setup**
   ```bash
   git clone https://github.com/hifat/kubo-storage-api.git
   cd kubo-storage-api
   go mod download
   ```

2. **Configure environment**
   ```bash
   cp env/.env.example env/.env
   # Edit env/.env with your Cloudflare R2 credentials
   ```

3. **Run the application**
   ```bash
   make run-grpc
   ```
   Or directly:
   ```bash
   go run ./cmd/r2 -envPath=./env/.env
   ```

### Docker

```bash
# Build image
make dck-build

# Run container
make dck-run
```

---

## Tech Stack

| Component | Technology |
|-----------|-----------|
| Language | Go |
| App Framework | GoFr |
| Dependency Injection | Google Wire |
| Cloud Storage | AWS SDK v2 |
| S3 Service | AWS SDK S3 |
| Configuration | Viper |

---

## Note
[AWS S3 SDK](https://docs.aws.amazon.com/sdk-for-go/v2/developer-guide/go_s3_code_examples.html)

