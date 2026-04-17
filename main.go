package main

import (
	"fmt"
	"os"
	"strings"
)

// uppercase,lowercase,bin,hex

func Readfile(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error in reading file")
	}
	redata := string(data)
	return redata
}
func Captializefirst(word string) string {
	return strings.ToUpper(string(word[0])) + word[1:]
}

func main() {
	Value := Readfile("sample.txt")
	finaldata := strings.Fields(Value)
	var result []string

	for i := 0; i < len(finaldata); i++ {
		switch finaldata[i] {
		case "(up)":
			previousword := finaldata[i-1]
			upperword := strings.ToUpper(previousword)
			result = append(result[:len(result)-1], upperword)
			continue
		case "(low)":
			previousword := finaldata[i-1]
			lowerword := strings.ToLower(previousword)
			result = append(result[:len(result)-1], lowerword)
			continue
		case "(cap)":
			previousword := finaldata[i-1]
			capword := Captializefirst(previousword)
			result = append(result[:len(result)-1], capword)
		default:
			result = append(result, finaldata[i])
		}

	}
	// for i := 0; i < len(finaldata); i++ {
	// 	switch finaldata[i] {
	// 	case "(low)":
	// 		oldword := finaldata[i-1]
	// 		lowerword := strings.ToLower(oldword)
	// 		result = append(result[:len(result)-1], lowerword)
	// 	default:
	// 		result = append(result, finaldata[i])
	// 	}
	// }

	fmt.Println(result)
}
