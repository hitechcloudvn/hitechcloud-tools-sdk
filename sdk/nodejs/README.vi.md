# HiTechCloud Tools SDK - Node.js

> [English](README.md) | **Tiếng Việt**

---

SDK Node.js chính thức cho HiTechCloud Tools API (455 endpoint).

## Cài đặt

```bash
npm install hitechcloud-tools-sdk
```

## Bắt đầu nhanh

```javascript
const { HiTechCloudTools } = require("hitechcloud-tools-sdk");

const client = new HiTechCloudTools({ apiKey: "your-api-key" });

// GeoIP lookup
const result = await client.network.geoip_lookup({ ip: "8.8.8.8" });
```
