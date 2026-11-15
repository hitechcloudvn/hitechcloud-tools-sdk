# HiTechCloud hitechcloud-tools-sdk — Swift SDK

> **English** | [Tiếng Việt](README.vi.md)

---

> **Status:** Planned — contributions welcome!

## Installation

```swift
dependencies: [
    .package(url: "https://github.com/hitechcloud/hitechcloud-tools-sdk-swift.git", from: "1.0.0")
]
```

## Quick Start

```swift
import HiTechCloudSDK

let client = Client(baseUrl: "https://api-tools.hitechcloud.vn", apiKey: "your-api-key")
let result = try client.chuyenDoiAndDinhDang.barcodeCheckDigit()
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
