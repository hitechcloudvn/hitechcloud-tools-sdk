# Hướng dẫn SDK Ruby

> [English](ruby.md) | **Tiếng Việt**

---

## Cài đặt

```bash
gem install hitechcloud_tools
```

## Khởi tạo

```ruby
require 'hitechcloud_tools'

client = HiTechCloudTools::Client.new(api_key: 'your-api-key')
```

## Sử dụng cơ bản

```ruby
# GeoIP lookup
result = client.mang_and_ha_tang.geoip_lookup(ip: "8.8.8.8")
puts result

# Domain WHOIS
whois = client.ten_mien_and_ssl.domain_whois(domain: "example.com")
puts whois

# Exchange rates
rates = client.tai_chinh_and_ty_gia.exchange_rates(base: "USD")
puts rates
```

## Biến môi trường

| Biến | Mô tả |
|----------|-------------|
| `HITECHCLOUD_TOOLS_API_KEY` | API key Tools của bạn |
| `HITECHCLOUD_TOOLS_BASE_URL` | Base URL tùy chỉnh |

## Thêm ví dụ

Xem [`examples/ruby/`](../../examples/ruby/) để có ví dụ đầy đủ.
