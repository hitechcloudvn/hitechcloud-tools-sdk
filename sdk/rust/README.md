# HiTechCloud hitechcloud-tools-sdk — Rust SDK

> **English** | [Tiếng Việt](README.vi.md)

---

> **Status:** Planned — contributions welcome!

## Installation

```toml
[dependencies]
hitechcloud_tools_sdk = "0.1.0"
```

## Quick Start

```rust
use hitechcloud_tools_sdk::Client;

#[tokio::main]
async fn main() {
    let client = Client::new(Some("https://api-tools.hitechcloud.vn"), "your-api-key", Some(30), Some(3));
    let result = client.chuyen_doi_and_dinh_dang.barcode_check_digit(None).await.unwrap();
    println!("{:?}", result);
}
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
