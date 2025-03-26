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
```golang
    var conferenceName string = "GopherCon"
	const conferenceTickets = 50
	var remainingTickets = 50
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("Welcome to", conferenceName)
	fmt.Printf("Number of tickets left: %v and number of tickets remaining%v\n", conferenceTickets, remainingTickets)
	fmt.Println("Number of tickets remaining:", remainingTickets)

    //%v: value, %T : type
```
#### Data types
- phải khai báo với go compiler, có thể không khai báo bằng cách cung cấp giá trị
- int8, int16, int32, int64, uint8, 16, 32, 64
```go
    //os.Stdin là một đối tượng kiểu *os.File đại diện cho đầu vào tiêu chuẩn (standard input - stdin) của hệ thống.
	//Nó được sử dụng để nhận dữ liệu nhập từ bàn phím hoặc từ một file/kênh nhập khác. 
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(os.Stdin)

	var lastName string = "Nhân"
	const middleName string = "Ngọc"
	var numberTicket int = 1
	firstName := "Trần" // một cách gán giá trị kiểu khác
	
	var fullName string
	// fmt.Scanln(&fullName)

	fmt.Println("Nhập nhiều dòng (ấn Enter trên dòng trống để dừng):")
	var description string
	for scanner.Scan(){
		line := scanner.Text();
		if line == ""{
			break
		}
		description += line + "\n"
	}
	//nhập một dòng
	fullName, _ = reader.ReadString('\r') // \n
	fmt.Println(fullName)
	fmt.Printf("%T\n", numberTicket)
	fmt.Printf("%v %v %v",firstName, middleName, lastName)
```
#### logic
```go
    var a = 5
	a %= 5
	a--
	a -= 51
```
#### array
```go
    // var variable_name [length]type
	var numbers = [50]int{1,23,4} // var numbers [50]int
	numbers[2] = 3
	length := len(numbers)
	fmt.Println(numbers)
	fmt.Println(length)
	//get type of array
	fmt.Printf("%T\n", numbers)
```

#### slices
- không cần chỉ định size
- index, resize khi cần thiết
```go
    //var numbers []int
	//var numbers = []int{}
	numbers := []int{}

	numbers = append(numbers, 1)

	fmt.Println(numbers)
	fmt.Printf("%v\n", numbers)
```