# Hướng dẫn SDK Java

> [English](java.md) | **Tiếng Việt**

---

## Cài đặt

```java
<dependency>
    <groupId>vn.hitechcloud</groupId>
    <artifactId>tools-sdk</artifactId>
    <version>1.0.0</version>
</dependency>
```

## Khởi tạo

```java
import vn.hitechcloud.tools.HiTechCloudTools;

HiTechCloudTools client = new HiTechCloudTools("your-api-key");
```

## Sử dụng cơ bản

```java
// GeoIP lookup
var result = client.mangAndHaTang().geoipLookup("8.8.8.8");
System.out.println(result);

// Domain WHOIS
var whois = client.tenMienAndSsl().domainWhois("example.com");
System.out.println(whois);
```

## Biến môi trường

| Biến | Mô tả |
|----------|-------------|
| `HITECHCLOUD_TOOLS_API_KEY` | API key Tools của bạn |
| `HITECHCLOUD_TOOLS_BASE_URL` | Base URL tùy chỉnh |

## Thêm ví dụ

Xem [`examples/java/`](../../examples/java/) để có ví dụ đầy đủ.
