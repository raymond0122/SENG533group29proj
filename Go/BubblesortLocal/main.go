package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()

	filename := "pi_input_25.json" // <-- Change this manually for 25/50/100 runs

	numbers, err := loadFromFile(filename)
	if err != nil {
		fmt.Println("Failed to load file:", err)
		return
	}

	linearSort(numbers)

	duration := time.Since(start)

	// Output sorted numbers and runtime
	respJSON, err := json.MarshalIndent(numbers, "", "  ")
	if err != nil {
		fmt.Println("Failed to marshal result:", err)
		return
	}

	fmt.Println(string(respJSON))
	fmt.Printf("Runtime: %d ms\n", duration.Milliseconds())
}

func loadFromFile(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data struct {
		Numbers []int `json:"numbers"`
	}
	err = json.NewDecoder(file).Decode(&data)
	return data.Numbers, err
}

func linearSort(list []int) {
	i := 0
	for i < len(list)-1 {
		if list[i] > list[i+1] {
			list[i], list[i+1] = list[i+1], list[i]
			i = 0
		} else {
			i++
		}
	}
}
