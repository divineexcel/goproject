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
		default:
			result = append(result, finaldata[i])
		}
	}

	fmt.Println(result)

}
