# GOLANG

### Đặc điểm
- `Go` được thiết kế để chạy trên **multiple cores** và xây dựng để hỗ trợ **concurrency**.
- `Concurrency của Go` có chi phí rẻ và dễ dàng.
- `Go` được sử dụng viết các `performant applications`
- Các ứng dụng chạy trên `scaled, distributed systems` như là có hàng ngàn server, cloud platform.
- `Go` kết hợp giữa systax đơn giản và dễ đọc như Python với effiency và safety của ngôn ngữ bậc trung như C++.
- Tiết kiệm tài nguyên hệ thống.
- Code được compiled thành machine code nên tốt hơn ác ngôn ngữ thông dịch như python phải chạy từng dòng, khi chương trình chạy thì máy tính sẽ thực thi trực tiếp mã máy, không cần xử lí trung gian còn thông dịch thì phải chạy từng dòng, gây tốn tài nguyên và thời gian.
- Consistent giữa các hệ điều hành.
- IDE : Visual Studio Code
- Go compiler
- Extensions: go,

- Tạo file go.mod: mô tả về module path, version, module path có thể là một import path
```sh
go mod init booking-app
```
- Tất cả code thuộc về package
- First statement in Go file là package ...
- Entry point: chỉ có một main func
- `fmt` : Format package
- Go được tổ chức thành những package (tập hợp nhiều source files)
```go
func main()
```

- Chạy code
```sh
go run main.go
```

### Variables & Constant
25 keywords
- variables, packages phải được sử dụng, không thì compiler sẽ báo lỗi.
- constant: giá trị không được thay đổi

```sh
var conferenceName string = "GopherCon" #variable
const conferenceTickets = 50 #constant
```

#### Formatting Output