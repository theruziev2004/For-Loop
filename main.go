package main

import "fmt"

func main() {
	/*for i := 0; i < 5; i ++{
		fmt.Println(i)
		fmt.Println("Hello World")
	}*/

	/*a := 10
	b := a / 1
	switch b {
	case 1:
		fmt.Println("Простая")
		break
	default:
		fmt.Println("Не простая")
	}*/

	i := 10_000
    k := 0
    for i < 20_000 {
      i += 152
      k++
	  fmt.Println(k)
	  fmt.Println(i)
    }
    

	
}
