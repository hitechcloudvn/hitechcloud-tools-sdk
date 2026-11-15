# Xác thực

> [English](authentication.md) | **Tiếng Việt**

---

HiTechCloud Tools API sử dụng xác thực bằng API key.

## API key

Chuyển API key của bạn khi khởi tạo client:

```python
client = HiTechCloudTools(api_key="your-api-key")
```

Key được gửi trong header HTTP `X-API-Key`.

## Lấy API key của bạn

1. Đăng nhập vào [my.hitechcloud.vn](https://my.hitechcloud.vn)
2. Vào **Cài đặt** → **API Keys**
3. Tạo một API key Tools mới

## Bảo mật

- Không bao giờ commit API key vào kho lưu trữ mã nguồn
- Sử dụng biến môi trường: `HITECHCLOUD_TOOLS_API_KEY`
- Thay key định kỳ
