# HiTechCloud Tools SDK - Python

> [English](README.md) | **Tiếng Việt**

---

SDK Python chính thức cho HiTechCloud Tools API (455 endpoint, 15 nhóm).

## Cài đặt

```bash
pip install hitechcloud-tools-sdk
```

## Bắt đầu nhanh

```python
from hitechcloud_tools import HiTechCloudTools

client = HiTechCloudTools(api_key="your-api-key")

# GeoIP lookup
result = client.network.geoip_lookup(ip="8.8.8.8")

# Domain WHOIS
result = client.domains.domain_whois(domain="example.com")

# SSL check
result = client.domains.ssl(domain="example.com")
```

## Nhóm API

- **Mạng & Hạ tầng** (46 endpoint)
- **Tên miền & SSL** (53 endpoint)
- **Email & DNS** (17 endpoint)
- **Công cụ lập trình** (25 endpoint)
- **SEO & Web** (13 endpoint)
- **Dữ liệu Việt Nam** (66 endpoint)
- **Xuất nhập khẩu & Logistics** (15 endpoint)
- **Dược phẩm & Y tế** (5 endpoint)
- **Mã hóa & Kiểm tra** (39 endpoint)
- **Chuyển đổi & Định dạng** (45 endpoint)
- **Thời gian & Lịch** (18 endpoint)
- **Tài chính & Tỷ giá** (16 endpoint)
- **Pháp lý & Thủ tục** (28 endpoint)
- **QR & Thanh toán** (5 endpoint)
- **Doanh nghiệp & Thuế** (64 endpoint)

## Giấy phép

MIT
