# HiTechCloud Tools SDK - Python

> **English** | [Tiếng Việt](README.vi.md)

---

Official Python SDK for HiTechCloud Tools API (455 endpoints, 15 groups).

## Installation

```bash
pip install hitechcloud-tools-sdk
```

## Quick Start

```python
from hitechcloud_tools import HiTechCloudTools

client = HiTechCloudTools(api_key="your-api-key")

# GeoIP lookup
result = client.network.geoip_lookup(ip="8.8.8.8")

# Domain WHOIS
result = client.domains.domain_whois(domain="example.com")

# SSL check
result = client.domains.ssl(domain="example.com")
```

## API Groups

- **Network & Infrastructure** (46 endpoints)
- **Domains & SSL** (53 endpoints)
- **Email & DNS** (17 endpoints)
- **Developer Tools** (25 endpoints)
- **SEO & Web** (13 endpoints)
- **Vietnam Data** (66 endpoints)
- **Import/Export & Logistics** (15 endpoints)
- **Pharma & Healthcare** (5 endpoints)
- **Encryption & Validation** (39 endpoints)
- **Conversion & Formatting** (45 endpoints)
- **Time & Calendar** (18 endpoints)
- **Finance & Exchange Rates** (16 endpoints)
- **Legal & Procedures** (28 endpoints)
- **QR & Payments** (5 endpoints)
- **Business & Tax** (64 endpoints)

## License

MIT
