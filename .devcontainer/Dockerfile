FROM golang:1.21.3-bookworm

# Install required packages
RUN apt update && \
  apt install -y --no-install-recommends \
  default-mysql-client \
  gcc \
  git \
  jq \
  libc-dev \
  make \
  vim \
  wget \
  zip && \
  apt clean && \
  rm -rf /var/lib/apt/lists/*

# Install sqldef
RUN ARCH=$(uname -m) && \
  case "$ARCH" in \
    x86_64) FILE="mysqldef_linux_amd64.tar.gz" ;; \
    aarch64) FILE="mysqldef_linux_arm64.tar.gz" ;; \
    *) echo "Unsupported architecture" && exit 1 ;; \
  esac && \
  curl -OL "https://github.com/k0kubun/sqldef/releases/download/v0.16.9/$FILE" && \
  tar -zxvf "$FILE" && \
  mv mysqldef /usr/local/bin && \
  rm "$FILE"

# Set working directory
WORKDIR /go/src/github.com/k-akari/payment.com

# Copy source files
COPY . .

# Install tools
RUN go install go.uber.org/mock/mockgen@latest && \
  go install golang.org/x/tools/cmd/goimports@latest && \
  go install golang.org/x/tools/gopls@latest && \
  go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest && \
  go install github.com/go-delve/delve/cmd/dlv@latest
