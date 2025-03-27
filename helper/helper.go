package helper

import (
	"fmt"
)

func Greet(name string) string {
	return fmt.Sprintf("Hello %s", name)
}
func TestRecursion(n int) int {
	if n == 0 {
		return 0
	}
	return n + TestRecursion(n-1)
}


func PrintName () {
	fmt.Println("Hello Nhan")
}

func TestReturnMultipleValue() (int, int){
	return 1, 2
}

