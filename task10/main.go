package main

import "fmt"

func main() {
	temperatures := []float32 {-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	mapTemperatures := make(map[int][]float32)
	
	for _, v := range temperatures {
		mapTemperatures[int(v) / 10 * 10] = append(mapTemperatures[int(v) / 10 * 10], v)
	}

	for k, v := range mapTemperatures {
		fmt.Print(k,":", v, " ")
	}

}