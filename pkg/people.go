package pkg

import "fmt"

func SayHi(name string) string {
	fmt.Println("v2")
	fmt.Println("hello world!")
	return fmt.Sprintf("Hi,%s", name)
}
