package pkg

import "fmt"

func SayHi(name string) string {
	fmt.Println("hello world!")
	return fmt.Sprintf("Hi,%s", name)
}
