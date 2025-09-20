# Api-Getway

High-performance API gateway to Golang

```bash
go version go1.24.6
```

## How to get started?

```bash
git clone https://gitlab.com/fmkv/fmkv-backend.git
cd fmkv-backend.git

go mod tidy

touch .env

echo "
PORT=<port>
SERVICES=<name>=<url>:<port>,<name>=<url>:<port>?
READ_TIMEOUT_SEC=<number>
WRITE_TIMEOUT_SEC=<number>
PROXY_TIMEOUT_SEC=<number>
" > .env

go run ./cmd/gateway/main.go
```
