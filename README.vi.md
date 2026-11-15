# HiTechCloud Tools SDK

> [English](README.md) | **Tiếng Việt**

---

[![CI](https://github.com/hitechcloudvn/hitechcloud-tools-sdk/actions/workflows/ci.yml/badge.svg)](https://github.com/hitechcloudvn/hitechcloud-tools-sdk/actions/workflows/ci.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

SDK đa ngôn ngữ chính thức cho HiTechCloud Tools API — **455 endpoint** trên **15 nhóm**.

## Ngôn ngữ hỗ trợ

| Ngôn ngữ | Trạng thái | Gói |
|----------|--------|---------|
| Python | ✅ Đầy đủ | `pip install hitechcloud-tools-sdk` |
| Node.js | ✅ Đầy đủ | `npm install hitechcloud-tools-sdk` |
| PHP | ✅ Đầy đủ | `composer require hitechcloud/tools-sdk` |
| Go | ✅ Đầy đủ | `go get github.com/hitechcloudvn/hitechcloud-tools-sdk` |
| Ruby | ✅ Đầy đủ | `gem install hitechcloud_tools` |
| Java | ✅ Đầy đủ | Maven: `vn.hitechcloud:hitechcloud-tools-sdk` |
| C# | ✅ Đầy đủ | NuGet: `HiTechCloud.Tools` |
| Rust | 📋 Bản mẫu | Hoan nghênh đóng góp |
| Kotlin | 📋 Bản mẫu | Hoan nghênh đóng góp |
| Swift | 📋 Bản mẫu | Hoan nghênh đóng góp |
| Dart | 📋 Bản mẫu | Hoan nghênh đóng góp |
| C | 📋 Bản mẫu | Hoan nghênh đóng góp |
| Clojure | 📋 Bản mẫu | Hoan nghênh đóng góp |
| F# | 📋 Bản mẫu | Hoan nghênh đóng góp |
| Julia | 📋 Bản mẫu | Hoan nghênh đóng góp |
| Objective-C | 📋 Bản mẫu | Hoan nghênh đóng góp |
| OCaml | 📋 Bản mẫu | Hoan nghênh đóng góp |
| PowerShell | 📋 Bản mẫu | Hoan nghênh đóng góp |
| R | 📋 Bản mẫu | Hoan nghênh đóng góp |
| Shell | 📋 Bản mẫu | Hoan nghênh đóng góp |

## Nhóm API (455 endpoint)

| # | Nhóm | Endpoint |
|---|-------|-----------|
| 1 | Mạng & Hạ tầng | 46 |
| 2 | Tên miền & SSL | 53 |
| 3 | Email & DNS | 17 |
| 4 | Công cụ lập trình | 25 |
| 5 | SEO & Web | 13 |
| 6 | Dữ liệu Việt Nam | 66 |
| 7 | Xuất nhập khẩu & Logistics | 15 |
| 8 | Dược phẩm & Y tế | 5 |
| 9 | Mã hóa & Kiểm tra | 39 |
| 10 | Chuyển đổi & Định dạng | 45 |
| 11 | Thời gian & Lịch | 18 |
| 12 | Tài chính & Tỷ giá | 16 |
| 13 | Pháp lý & Thủ tục | 28 |
| 14 | QR & Thanh toán | 5 |
| 15 | Doanh nghiệp & Thuế | 64 |

## Bắt đầu nhanh

### Python
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

### Node.js
```javascript
const { HiTechCloudTools } = require("hitechcloud-tools-sdk");

const client = new HiTechCloudTools({ apiKey: "your-api-key" });

// GeoIP lookup
const result = await client.network.geoip_lookup({ ip: "8.8.8.8" });
```

### PHP
```php
use HiTechCloud\Tools\HiTechCloudTools;

$client = new HiTechCloudTools('https://api-tools.hitechcloud.vn', 'your-api-key');

$result = $client->network->geoip_lookup(['ip' => '8.8.8.8']);
```

### Go
```go
client := hitechcloudtools.NewClient("your-api-key")

result, err := client.Network.GeoipLookup(map[string]string{"ip": "8.8.8.8"})
```

## Xác thực

Mọi yêu cầu đều cần API key qua header `X-API-Key`:

```bash
curl -H "X-API-Key: your-api-key" https://api-tools.hitechcloud.vn/api/geoip/lookup?ip=8.8.8.8
```

Cách khác: `Authorization: Bearer your-api-key`

## Giới hạn tần suất

- **authenticated**: 100 request/phút
- **authenticated**: 1000 request/giờ

## Tài liệu

- Tài liệu API: https://doc-api-tools.hitechcloud.vn
- Tài liệu SDK: https://doc-api-tools.hitechcloud.vn/sdk

## Giấy phép

MIT © HiTechCloud Vietnam
