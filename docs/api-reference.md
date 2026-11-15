# API Reference

> **English** | [Tiếng Việt](api-reference.vi.md)

---

Base URL: `https://api-tools.hitechcloud.vn`

## Resource Groups (15 groups, 455 endpoints)

| Group | Resource Class | Description |
|-------|---------------|-------------|
| Conversion & Formatting | `chuyen_doi_and_dinh_dang` | Unit conversion, encoding, formatting |
| Developer Tools | `cong_cu_lap_trinh` | Developer tools (regex, hash, base64, JWT) |
| Business & Tax | `doanh_nghiep_and_thue` | Business lookup, tax codes, enterprise data |
| Vietnam Data | `du_lieu_viet_nam` | Vietnamese data (provinces, districts, wards) |
| Pharma & Healthcare | `duoc_pham_and_y_te` | Pharmaceutical and medical data |
| Email & DNS | `email_and_dns` | Email validation, DNS lookup |
| Chemistry & Science | `hoa_and_khoa_hoc` | Chemistry, scientific calculations |
| Network & Infrastructure | `mang_and_ha_tang` | GeoIP, ASN, network infrastructure |
| Encryption & Validation | `ma_hoa_and_kiem_tra` | Encryption, hashing, checksums |
| QR & Payments | `qr_and_thanh_toan` | QR code generation, payment integration |
| SEO & Web | `seo_and_web` | SEO analysis, web performance |
| Finance & Exchange Rates | `tai_chinh_and_ty_gia` | Exchange rates, financial calculations |
| Time & Calendar | `thoi_gian_and_lich` | Time zones, calendar, lunar calendar |
| Domains & SSL | `ten_mien_and_ssl` | Domain WHOIS, SSL certificate checks |
| Legal & Procedures | `phap_ly_and_thu_tuc` | Legal procedures, government forms |

## Authentication

All requests require the `X-API-Key` header.

## Response Format

All responses are JSON with HTTP 200 for success.

```json
{
  "success": true,
  "data": { ... }
}
```
