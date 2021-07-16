package main

import "fmt"

func main() {
	/*for i := 0; i < 5; i ++{
		fmt.Println(i)
		fmt.Println("Hello World")
	}*/
		i := 0
		k := 0
		for i < 110 {
			i += 10
			k++
		}
		fmt.Println(k)
	/*a := 10
	b := a / 1
	switch b {
	case 1:
		fmt.Println("Простая")
		break
	default:
		fmt.Println("Не простая")
	}*/
}
