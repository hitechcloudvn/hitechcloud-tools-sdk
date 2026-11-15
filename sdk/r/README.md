# HiTechCloud hitechcloud-tools-sdk — R SDK

> **English** | [Tiếng Việt](README.vi.md)

---

> **Status:** Planned — contributions welcome!

## Installation

```r
devtools::install_github("hitechcloud/hitechcloud-tools-sdk-r")
```

## Quick Start

```r
library(hitechcloud.tools.sdk)
client <- Client(base_url = "https://api-tools.hitechcloud.vn", api_key = "your-api-key")
result <- barcodeCheckDigit(client$chuyenDoiAndDinhDang)
print(result)
```

## Resources

| Resource | Description | Methods |
|----------|-------------|---------|
| `chuyenDoiAndDinhDang` | Conversion & Formatting | 37 |
| `congCuLapTrinh` | Developer Tools | 23 |
| `doanhNghiepAndThue` | Business & Tax | 64 |
| `duLieuVietNam` | Vietnam Data | 64 |
| `duocPhamAndYTe` | Pharma & Healthcare | 5 |
| `emailAndDns` | Email & DNS | 17 |
| `maHoaAndKiemTra` | Encryption & Validation | 28 |
| `mangAndHaTang` | Network & Infrastructure | 43 |
| `phapLyAndThuTuc` | Legal & Procedures | 26 |
| `qrAndThanhToan` | QR & Payments | 5 |
| `seoAndWeb` | SEO & Web | 12 |
| `taiChinhAndTyGia` | Finance & Exchange Rates | 15 |
| `tenMienAndSsl` | Domains & SSL | 52 |
| `thoiGianAndLich` | Time & Calendar | 17 |
| `xuatNhapKhauAndLogistics` | Import/Export & Logistics | 15 |

## License

MIT
