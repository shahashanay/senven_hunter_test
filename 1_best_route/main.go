package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
)

func readJSONFile(filename string) ([][]int, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var data [][]int
	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func maxPathSum(triangle [][]int) int {
	for row := len(triangle) - 2; row >= 0; row-- {
		for col := 0; col <= row; col++ {
			triangle[row][col] += int(math.Max(float64(triangle[row+1][col]), float64(triangle[row+1][col+1])))
		}
	}
	return triangle[0][0]
}

func main() {
	// Test case 1 (Hardcoded data)
	testData := [][]int{
		{59},
		{73, 41},
		{52, 40, 53},
		{26, 53, 6, 34},
	}
	result := maxPathSum(testData)
	fmt.Println("Maximum path sum for test case 1:", result)

	// Test case 2 (Reading input from JSON file)
	filename := "hard.json" 
	triangle, err := readJSONFile(filename)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	result = maxPathSum(triangle)
	fmt.Println("Maximum path sum for test case 2:", result)
}
