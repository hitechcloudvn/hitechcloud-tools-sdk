# HiTechCloud hitechcloud-tools-sdk — Clojure SDK

> **English** | [Tiếng Việt](README.vi.md)

---

> **Status:** Planned — contributions welcome!

## Quick Start

```clojure
(require '[hitechcloud.sdk.client :as client])
(def c (client/make-client :base-url "https://api-tools.hitechcloud.vn" :api-key "your-api-key"))
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
