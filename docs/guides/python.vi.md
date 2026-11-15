# Hướng dẫn SDK Python

> [English](python.md) | **Tiếng Việt**

---

## Cài đặt

```bash
pip install hitechcloud-tools
```

## Khởi tạo

```python
from hitechcloud_tools import HiTechCloudTools

client = HiTechCloudTools(api_key="your-api-key")
```

## Sử dụng cơ bản

```python
# GeoIP lookup
result = client.mang_and_ha_tang.geoip_lookup(ip="8.8.8.8")
print(result)

# Domain WHOIS
whois = client.ten_mien_and_ssl.domain_whois(domain="example.com")
print(whois)

# Exchange rates
rates = client.tai_chinh_and_ty_gia.exchange_rates(base="USD")
print(rates)

# Vietnamese provinces
provinces = client.du_lieu_viet_nam.provinces()
print(provinces)
```

## Biến môi trường

| Biến | Mô tả |
|----------|-------------|
| `HITECHCLOUD_TOOLS_API_KEY` | API key Tools của bạn |
| `HITECHCLOUD_TOOLS_BASE_URL` | Base URL tùy chỉnh |

## Thêm ví dụ

Xem [`examples/python/`](../../examples/python/) để có ví dụ đầy đủ.
