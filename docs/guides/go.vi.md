# Hướng dẫn SDK Go

> [English](go.md) | **Tiếng Việt**

---

## Cài đặt

```bash
go get github.com/hitechcloudvn/hitechcloud-tools-sdk/sdk/go
```

## Khởi tạo

```go
package main

import (
    "github.com/hitechcloudvn/hitechcloud-tools-sdk/sdk/go"
)

func main() {
    client := hitechcloudtools.NewClient("your-api-key")
}
```

## Sử dụng cơ bản

```go
// GeoIP lookup
result, err := client.MangAndHaTang.GeoipLookup("8.8.8.8")
if err != nil {
    log.Fatal(err)
}
fmt.Println(result)

// Domain WHOIS
whois, err := client.TenMienAndSsl.DomainWhois("example.com")
if err != nil {
    log.Fatal(err)
}
fmt.Println(whois)
```

## Biến môi trường

| Biến | Mô tả |
|----------|-------------|
| `HITECHCLOUD_TOOLS_API_KEY` | API key Tools của bạn |
| `HITECHCLOUD_TOOLS_BASE_URL` | Base URL tùy chỉnh |

## Thêm ví dụ

Xem [`examples/go/`](../../examples/go/) để có ví dụ đầy đủ.
