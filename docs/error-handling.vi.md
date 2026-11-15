# Xử lý lỗi

> [English](error-handling.md) | **Tiếng Việt**

---

## Mã trạng thái HTTP

| Mã | Ý nghĩa |
|------|---------|
| 200 | Thành công |
| 401 | API key không hợp lệ hoặc bị thiếu |
| 403 | Bị cấm truy cập |
| 429 | Vượt giới hạn tần suất |
| 500 | Lỗi máy chủ nội bộ |

## Định dạng phản hồi lỗi

```json
{
  "error": "Invalid API key",
  "code": 401
}
```

## Xử lý lỗi trong SDK

```python
from hitechcloud_tools import HiTechCloudTools, HiTechCloudToolsError

client = HiTechCloudTools(api_key="your-key")

try:
    result = client.mang_and_ha_tang.geoip_lookup(ip="8.8.8.8")
except HiTechCloudToolsError as e:
    print(f"Error: {e}")
```
