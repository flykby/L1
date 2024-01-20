package main

import (
	"fmt"
)

func main() {
	var data int64
	var n int8
	fmt.Print("Print the number int64: ")
	_, ok := fmt.Scan(&data)
	if ok != nil {
		return
	}
	fmt.Print("Print the bit number (>=0, <64): ")
	_, ok = fmt.Scan(&n)
	if ok != nil {
		return
	}
	if n < 0 || n >= 64 {
		fmt.Println("Invalid input of bit number")
		return
	}

	data = setBit(data, n)
	fmt.Println(data)

	data = clearBit(data, n)
	fmt.Println(data)

}

func setBit(data int64, n int8) int64 {
	data |= 1 << n
	return data
}

func clearBit(data int64, n int8) int64 {
	return data &^ (1 << n)
}