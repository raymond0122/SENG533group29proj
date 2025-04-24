package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()

	// Hardcoded file
	filename := "string_input_500.json"

	str, err := loadStringFromFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	letterCounts := countLetters(str)

	duration := time.Since(start)

	// Print output and duration
	respJSON, err := json.MarshalIndent(letterCounts, "", "  ")
	if err != nil {
		fmt.Println("Failed to marshal result:", err)
		return
	}

	fmt.Println(string(respJSON))
	fmt.Printf("Runtime: %d ms\n", duration.Milliseconds())
}

func loadStringFromFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var data struct {
		Data string `json:"data"`
	}
	err = json.NewDecoder(file).Decode(&data)
	return data.Data, err
}

func countLetters(s string) [26]int {
	var counts [26]int
	for i := 0; i < len(s); i++ {
		index := s[i] - 'a'
		if index >= 0 && index < 26 {
			counts[index]++
		}
	}
	return counts
}
