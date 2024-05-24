package main

import (
	"fmt"
	"strings"
)

// func IsPalindrome(str string) bool {
// 	str = strings.ToLower(str)
// 	revStr := ""

// 	for idx := len(str) - 1; idx >= 0; idx-- {
// 		revStr += string(str[idx])
// 	}

// 	return str == revStr
// }

func IsPalindrome(str string) bool {
	str = strings.ToLower(str)
	length := len(str)

	for i := 0; i < length/2; i++ {
		if str[i] != str[length-i-1] {
			return false
		}
	}

	return true
}

func multipleOfIndex(ints []int) []int {
	// good luck

	multiple := []int{}

	for idx, num := range ints {
		if idx != 0 && num%idx == 0 {
			multiple = append(multiple, num)
		}
	}

	return multiple
}

func isValid(s string) bool {
	// if (s[0]) == ')' || (s[0]) == ']' || (s[0]) == '}' {
	// 	return false
	// }
	// if (s[len(s)-1]) == ')' || (s[len(s)-1]) == ']' || (s[len(s)-1]) == '}' {
	// 	return false
	// }

	str := []rune{}
	parthesis := map[rune]rune{
		'{': '}',
		'(': ')',
		'[': ']',
	}

	for _, bracket := range s {
		fmt.Println("rune", bracket)
		if bracket == '(' || bracket == '{' || bracket == '[' {
			str = append(str, bracket)
		} else if bracket == ')' || bracket == '}' || bracket == ']' {
			if len(str) == 0 {
				return false
			}
			top := str[len(str)-1]
			if bracket == parthesis[top] {
				str = str[:len(str)-1]
			} else {
				str = append(str, bracket)
			}
		}
	}

	return len(str) == 0
}

func main() {
	fmt.Println(isValid("{}"))
	// fmt.Println(multipleOfIndex([]int{22, -6, 32, 82, 9, 25}))
	// fmt.Println(IsPalindrome("racecar"))
	// fmt.Println(IsPalindrome("Abba"))
	// fmt.Println(IsPalindrome("a"))
	// fmt.Println(IsPalindrome("mike"))

}
