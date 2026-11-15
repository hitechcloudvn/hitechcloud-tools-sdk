# Máy chủ MCP HiTechCloud Tools

> [English](README.md) | **Tiếng Việt**

---

Máy chủ Model Context Protocol cho HiTechCloud Tools API.

## Thiết lập

```bash
cd mcp && npm install
export HITECHCLOUD_TOOLS_API_KEY=your-api-key
node server.js
```

## Công cụ khả dụng

| Công cụ | Mô tả |
|------|-------------|
| `geoip_lookup` | Tra cứu vị trí địa lý IP |
| `geoip_batch` | Tra cứu IP hàng loạt |
| `domain_whois` | Tra cứu WHOIS tên miền |
| `ssl_check` | Kiểm tra chứng chỉ SSL |
| `exchange_rates` | Lấy tỷ giá |
| `vietnam_provinces` | Liệt kê tỉnh thành Việt Nam |
| `hash_generate` | Tạo hash |
| `base64_encode` | Mã hóa base64 |
| `regex_test` | Kiểm tra biểu thức regex |

## Cấu hình Claude Desktop

```json
{
  "mcpServers": {
    "hitechcloud-tools": {
      "command": "node",
      "args": ["/path/to/hitechcloud-tools-sdk/mcp/server.js"],
      "env": {
        "HITECHCLOUD_TOOLS_API_KEY": "your-api-key"
      }
    }
  }
}
```
