# Go Server

HTTP server prototype written in Go.

## Sample ./server.yaml Config

```
server:
  port: :5000
  protocol: HTTP
  https:
    cert: cert.pem
    key: key.pem

db:
  connection: postgres://<username>:<passowrd>@<host>:<port>/<database>?sslmode=disable

gin:
  release: true

gorm:
  level: 3

logger:
  cli:
    level: 9
    colorful: true
    title: false
    timestamp: "[15:04:05]"
  file:
    level: 9
    title: false
    timestamp: "[15:04:05]"
    path: logs/2006-01-02.log
```
