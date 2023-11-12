#!/bin/sh

go mod download

# Generate private and public key for JWT if not exist
openssl version
if [ ! -f './internal/auth/cert/secret.pem' ]; then
  openssl genrsa 4096 > ./internal/auth/cert/secret.pem
fi
if [ ! -f './internal/auth/cert/public.pem' ]; then
  openssl rsa -pubout < ./internal/auth/cert/secret.pem > ./internal/auth/cert/public.pem
fi
