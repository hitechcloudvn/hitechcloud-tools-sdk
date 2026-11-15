# Hướng dẫn SDK Node.js

> [English](nodejs.md) | **Tiếng Việt**

---

## Cài đặt

```bash
npm install @hitechcloud/tools-sdk
```

## Khởi tạo

```nodejs
const HiTechCloudTools = require('@hitechcloud/tools-sdk');

const client = new HiTechCloudTools({ apiKey: 'your-api-key' });
```

## Sử dụng cơ bản

```nodejs
// GeoIP lookup
const result = await client.mangAndHaTang.geoipLookup({ ip: '8.8.8.8' });
console.log(result);

// Domain WHOIS
const whois = await client.tenMienAndSsl.domainWhois({ domain: 'example.com' });
console.log(whois);

// Exchange rates
const rates = await client.taiChinhAndTyGia.exchangeRates({ base: 'USD' });
console.log(rates);
```

## Biến môi trường

| Biến | Mô tả |
|----------|-------------|
| `HITECHCLOUD_TOOLS_API_KEY` | API key Tools của bạn |
| `HITECHCLOUD_TOOLS_BASE_URL` | Base URL tùy chỉnh |

## Thêm ví dụ

Xem [`examples/nodejs/`](../../examples/nodejs/) để có ví dụ đầy đủ.
