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

# Set working directory
WORKDIR /go/src/github.com/k-akari/payment.com

# Copy source files
COPY . .

# Install tools
RUN go install golang.org/x/tools/cmd/goimports@latest && \
  go install golang.org/x/tools/gopls@latest && \
  go install github.com/go-delve/delve/cmd/dlv@latest
