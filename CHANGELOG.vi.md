# Nhật ký thay đổi

> [English](CHANGELOG.md) | **Tiếng Việt**

---

Mọi thay đổi đáng chú ý của HiTechCloud Tools SDK đều được ghi lại trong tệp này.

## [1.1.0] - 2026-11-16

### Đã đổi
- Đồng bộ mã nguồn SDK với các định nghĩa mới nhất của HiTechCloud API Tools
- Ngừng theo dõi các tệp build trong kho mã
- Sửa các quy trình CI/CD (release.yml, ci.yml), các lỗi của client Java/PHP/C# và npm scope cho GitHub Packages

### Thêm
- Tài liệu song ngữ Việt-Anh trên toàn bộ repo (mỗi tài liệu đều có bản
  mirror `.vi.md`).
- Pipeline GitLab CI mirror các test của GitHub Actions, kèm GitLab
  Releases khi có tag `v*`.

## [1.0.0] - 2024-12-01

### Thêm
- Phát hành đầu tiên với 7 SDK ngôn ngữ (Python, Node.js, PHP, Go, Ruby, Java, C#)
- 455 endpoint trên 15 nhóm tài nguyên
- Đặc tả OpenAPI 3.0
- Hỗ trợ Docker
- Các pipeline CI/CD
- Tài liệu toàn diện
