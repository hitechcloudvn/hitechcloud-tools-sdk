# Bắt đầu nhanh

> [English](getting-started.md) | **Tiếng Việt**

---

## Cài đặt

### Python
```bash
pip install hitechcloud-tools
```

### Node.js
```bash
npm install @hitechcloud/tools-sdk
```

### PHP
```bash
composer require hitechcloud/tools-sdk
```

### Go
```bash
go get github.com/hitechcloudvn/hitechcloud-tools-sdk/sdk/go
```

### Ruby
```bash
gem install hitechcloud_tools
```

### Java
```xml
<dependency>
    <groupId>vn.hitechcloud</groupId>
    <artifactId>tools-sdk</artifactId>
    <version>1.0.0</version>
</dependency>
```

### C#
```bash
dotnet add package HiTechCloud.Tools.SDK
```

## Xác thực

Tools API sử dụng xác thực bằng API key thông qua header `X-API-Key`.

```python
from hitechcloud_tools import HiTechCloudTools

client = HiTechCloudTools(api_key="your-api-key")
```
