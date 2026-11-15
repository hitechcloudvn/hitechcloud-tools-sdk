# Hướng dẫn SDK C#

> [English](csharp.md) | **Tiếng Việt**

---

## Cài đặt

```csharp
dotnet add package HiTechCloud.Tools.SDK
```

## Khởi tạo

```csharp
using HiTechCloud.Tools;

var client = new HiTechCloudToolsClient("your-api-key");
```

## Sử dụng cơ bản

```csharp
// GeoIP lookup
var result = await client.MangAndHaTang.GeoipLookupAsync("8.8.8.8");
Console.WriteLine(result);

// Domain WHOIS
var whois = await client.TenMienAndSsl.DomainWhoisAsync("example.com");
Console.WriteLine(whois);
```

## Biến môi trường

| Biến | Mô tả |
|----------|-------------|
| `HITECHCLOUD_TOOLS_API_KEY` | API key Tools của bạn |
| `HITECHCLOUD_TOOLS_BASE_URL` | Base URL tùy chỉnh |

## Thêm ví dụ

Xem [`examples/csharp/`](../../examples/csharp/) để có ví dụ đầy đủ.
