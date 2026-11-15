# Tra cứu dữ liệu Việt Nam

> [English](vietnam-data.md) | **Tiếng Việt**

---

```yaml
---
name: hitechcloudvn-data
description: Query Vietnamese administrative data (provinces, districts, wards)
---
```

Sử dụng kỹ năng này khi người dùng cần dữ liệu hành chính Việt Nam.

## Mẫu code

```python
# List all provinces
provinces = client.du_lieu_viet_nam.provinces()

# Get districts of a province
districts = client.du_lieu_viet_nam.districts(province_code="01")

# Get wards of a district
wards = client.du_lieu_viet_nam.wards(district_code="001")
```
