package main

import (
	"fmt"
	"os"
	"strconv"
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
			continue
		case "(hex)":
			previousword := finaldata[i-1]
			hex, err := strconv.ParseInt(previousword, 16, 64)

			if err != nil {
				fmt.Println(err)
				return
			}
			result = append(result[:len(result)-1], strconv.FormatInt(hex, 10))
		case "(up,":
			if i+1 < len(finaldata) {

				nextWord := finaldata[i+1]

				numpart, rest, found := strings.Cut(nextWord, ")")

				var intnum int

				var err error

				if found {

					intnum, err = strconv.Atoi(numpart)

					if err != nil {

						intnum = 1
					}

				} else {
					intnum = 1
				}

				start := len(result) - intnum

				if start < 0 {
					start = 0

				}
				for j := start; j < len(result); j++ {
					result[j] = strings.ToUpper(result[j])
				}
				if rest != "" {
					result = append(result, rest)
				}
				i++

				continue
			}
		case "(low,":
			if i+1 < len(finaldata) {
				nextWord := finaldata[i+1]
				numpart, rest, found := strings.Cut(nextWord, ")")

				var intnum int
				var err error
				if found {
					intnum, err = strconv.Atoi(numpart)
					if err != nil {
						intnum = 1
					}
				} else {
					intnum = 1
				}
				start := len(result) - intnum
				if start < 0 {
					start = 0
				}
				for j := start; j < len(result); j++ {
					result[j] = strings.ToLower(result[j])
				}
				if rest != "" {
					result = append(result, rest)
				}
				i++
				continue
			}

		default:
			result = append(result, finaldata[i])
		}

	}

	fmt.Println(strings.Join(result, " "))
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
}
