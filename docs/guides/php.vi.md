# Hướng dẫn SDK PHP

> [English](php.md) | **Tiếng Việt**

---

## Cài đặt

```bash
composer require hitechcloud/tools-sdk
```

## Khởi tạo

```php
<?php
use HiTechCloudTools\HiTechCloudTools;

$client = new HiTechCloudTools('your-api-key');
```

## Sử dụng cơ bản

```php
// GeoIP lookup
$result = client->mangAndHaTang()->geoipLookup(['ip' => '8.8.8.8']);
print_r($result);

// Domain WHOIS
$whois = client->tenMienAndSsl()->domainWhois(['domain' => 'example.com']);
print_r($whois);

// Exchange rates
$rates = client->taiChinhAndTyGia()->exchangeRates(['base' => 'USD']);
print_r($rates);
```

## Biến môi trường

| Biến | Mô tả |
|----------|-------------|
| `HITECHCLOUD_TOOLS_API_KEY` | API key Tools của bạn |
| `HITECHCLOUD_TOOLS_BASE_URL` | Base URL tùy chỉnh |

## Thêm ví dụ

Xem [`examples/php/`](../../examples/php/) để có ví dụ đầy đủ.
