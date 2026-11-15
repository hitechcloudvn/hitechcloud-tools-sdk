# Tham chiếu API

> [English](api-reference.md) | **Tiếng Việt**

---

Base URL: `https://api-tools.hitechcloud.vn`

## Nhóm tài nguyên (15 nhóm, 455 endpoint)

| Nhóm | Lớp tài nguyên | Mô tả |
|-------|---------------|-------------|
| Chuyển đổi & Định dạng | `chuyen_doi_and_dinh_dang` | Chuyển đổi đơn vị, mã hóa, định dạng |
| Công cụ lập trình | `cong_cu_lap_trinh` | Công cụ lập trình (regex, hash, base64, JWT) |
| Doanh nghiệp & Thuế | `doanh_nghiep_and_thue` | Tra cứu doanh nghiệp, mã số thuế, dữ liệu doanh nghiệp |
| Dữ liệu Việt Nam | `du_lieu_viet_nam` | Dữ liệu Việt Nam (tỉnh, quận/huyện, phường/xã) |
| Dược phẩm & Y tế | `duoc_pham_and_y_te` | Dữ liệu dược phẩm và y tế |
| Email & DNS | `email_and_dns` | Kiểm tra email, tra cứu DNS |
| Hóa & Khoa học | `hoa_and_khoa_hoc` | Hóa học, tính toán khoa học |
| Mạng & Hạ tầng | `mang_and_ha_tang` | GeoIP, ASN, hạ tầng mạng |
| Mã hóa & Kiểm tra | `ma_hoa_and_kiem_tra` | Mã hóa, băm, checksum |
| QR & Thanh toán | `qr_and_thanh_toan` | Tạo mã QR, tích hợp thanh toán |
| SEO & Web | `seo_and_web` | Phân tích SEO, hiệu năng web |
| Tài chính & Tỷ giá | `tai_chinh_and_ty_gia` | Tỷ giá, tính toán tài chính |
| Thời gian & Lịch | `thoi_gian_and_lich` | Múi giờ, lịch, lịch âm |
| Tên miền & SSL | `ten_mien_and_ssl` | WHOIS tên miền, kiểm tra chứng chỉ SSL |
| Pháp lý & Thủ tục | `phap_ly_and_thu_tuc` | Thủ tục pháp lý, biểu mẫu nhà nước |

## Xác thực

Mọi yêu cầu đều cần header `X-API-Key`.

## Định dạng phản hồi

Mọi phản hồi là JSON với HTTP 200 khi thành công.

```json
{
  "success": true,
  "data": { ... }
}
```
