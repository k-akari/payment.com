# 1. 開発環境のセットアップ方法
現時点では、VSCodeユーザー向けにのみ開発環境を整備しています。VSCodeユーザーはdevcontainerを利用してください。また、本アプリケーションではJWT認証を実装しています。秘密鍵と公開鍵を利用したRS256形式の署名を採用しているため、`/internal/auth/cert`ディレクトリ配下に`secret.pem`（秘密鍵）と`public.pem`（公開鍵）が配置されていることを期待しています。そのため、ローカルでアプリケーションを動作させる際は秘密鍵と公開鍵を上記のパスに配置してください。なお、devcontainerを利用した場合には自動的に秘密鍵と公開鍵がセットアップされるようになっています。

# 2. 動作確認方法
## 2-1. コンテナを起動する。
```sh
$ docker compose up -d
```

## 2-2. DBマイグレートを実行する。
```sh
$ make migrate
```

## 2-3. curlコマンドでエンドポイントを叩く。
### 2-3-1. 企業を作成する。
```sh
$ curl -X POST -H "Content-Type: application/json" -d '{"name":"company_name", "representative":"representative_name", "telephone_number":"080-1234-5678", "postal_code":"123-4567", "address":"tokyo shinjyuku-ku"}' localhost:8080/companies
```

### 2-3-2. 作成した企業の情報を確認する。
```sh
$ curl -X GET -H "Content-Type: application/json" localhost:8080/companies/1
```

### 2-3-3. 企業に紐づいた取引先を作成する。
```sh
$ curl -X POST -H "Content-Type: application/json" -d '{"name":"client_name", "representative":"representative_name", "telephone_number":"090-1234-5678", "postal_code":"765-4321", "address":"kyoto sakyo-ku"}' localhost:8080/companies/1/clients
```

### 2-3-4. 作成した取引先の情報を確認する。
```sh
$ curl -X GET -H "Content-Type: application/json" localhost:8080/companies/1/clients/1
```

### 2-3-5. 企業と取引先に紐づいた請求書データを作成する。
```sh
$ curl -X POST -H "Content-Type: application/json" -d '{"issued_date":"2023-10-10T17:44:13Z", "paid_amount":1000, "payment_due_date":"2023-10-31T17:44:13Z"}' localhost:8080/companies/1/clients/1/invoices
```

### 2-3-6. 支払い期日が指定した期間に含まれる請求書データの一覧を取得する。
```sh
$ curl -X GET -H "Content-Type: application/json" -d '{"from":"1980-10-10T17:44:13Z", "to":"2024-10-31T17:44:13Z"}' localhost:8080/companies/1/invoices
```
