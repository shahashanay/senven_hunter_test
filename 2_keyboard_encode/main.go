package main

import (
	"fmt"
	"strings"
)

func decodeToMinSum(encoded string) string {
	n := len(encoded) + 1 
	nums := make([]int, n)

	nums[0] = 0

	for i := 0; i < len(encoded); i++ {
		switch encoded[i] {
		case 'L':
			nums[i+1] = nums[i] - 1
			if nums[i+1] < 0 {
				nums[i]++ 
				nums[i+1] = 0 
				for j := i - 1; j >= 0 && (encoded[j] == 'L' || encoded[j] == '='); j-- {
					nums[j]++
				}
			}else if nums[i+1] > 0{
				nums[i+1] = 0
			}
		case 'R':
			nums[i+1] = nums[i] + 1
		case '=':
			nums[i+1] = nums[i]
		}
	}

	var result strings.Builder
	for _, num := range nums {
		result.WriteString(fmt.Sprintf("%d", num))
	}

	return result.String()
}

func main() {
	var encoded string
	fmt.Print("Enter encoded string: ")
	fmt.Scan(&encoded)

	decoded := decodeToMinSum(encoded)
	fmt.Println("Decoded number sequence with minimum sum:", decoded)
}