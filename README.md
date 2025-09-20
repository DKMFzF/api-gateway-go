# Api-Getway

High-performance API gateway to Golang

The api is a getway for a social network application, this repository is a duplicate of my repository in gitlab. The project is still in a closed stage, but there are already functional modules, such as api-getway.

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
