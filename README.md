# 1. Setting Up the Development Environment
Currently, the development environment is only set up for VSCode users. If you are a VSCode user, please use the devcontainer. This application implements JWT authentication, using RS256 signatures with private and public keys. Therefore, it is expected that secret.pem (private key) and public.pem (public key) are placed under the /internal/auth/cert directory. When running the application locally, please place the private and public keys at the above path. Note that if you use the devcontainer, the private and public keys will be set up automatically.

# 2. How to Test the Application
## 2-1. Start the container.
```sh
$ docker compose up -d
```

## 2-2. Execute database migration.
```sh
$ make migrate
```

## 2-3. Use the curl command to hit the endpoints.
### 2-3-1. Create a company.
```sh
$ curl -X POST -H "Content-Type: application/json" -d '{"name":"company_name", "representative":"representative_name", "telephone_number":"080-1234-5678", "postal_code":"123-4567", "address":"tokyo shinjyuku-ku"}' localhost:8080/companies
```

### 2-3-2. Check the information of the created company.
```sh
$ curl -X GET -H "Content-Type: application/json" localhost:8080/companies/1
```

### 2-3-3. Create a client associated with a company.
```sh
$ curl -X POST -H "Content-Type: application/json" -d '{"name":"client_name", "representative":"representative_name", "telephone_number":"090-1234-5678", "postal_code":"765-4321", "address":"kyoto sakyo-ku"}' localhost:8080/companies/1/clients
```

### 2-3-4. Check the information of the created client.
```sh
$ curl -X GET -H "Content-Type: application/json" localhost:8080/companies/1/clients/1
```

### 2-3-5. Create invoice data associated with a company and its client.
```sh
$ curl -X POST -H "Content-Type: application/json" -d '{"issued_date":"2023-10-10T17:44:13Z", "paid_amount":1000, "payment_due_date":"2023-10-31T17:44:13Z"}' localhost:8080/companies/1/clients/1/invoices
```

### 2-3-6. Retrieve a list of invoice data where the payment due date falls within a specified period.
```sh
$ curl -X GET -H "Content-Type: application/json" -d '{"from":"1980-10-10T17:44:13Z", "to":"2024-10-31T17:44:13Z"}' localhost:8080/companies/1/invoices
```
